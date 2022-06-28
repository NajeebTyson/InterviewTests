package onefootball

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"github.com/NajeebTyson/onefootball/internal/config"
	"github.com/NajeebTyson/onefootball/pkg/utils"
)

var (
	errTeamNotFound = errors.New("team id not found")
	errTeamNotExist = errors.New("team not exist")
)

// GetPlayers takes a list of team names as an argument and return their players data in a combine list ordered by player id
// if any team name not exist in the API, it will return error
func GetPlayers(teamNames []string) ([]Player, error) {
	playersMap := make(map[PlayerId]Player)

	teams, err := getTeamsFaster(teamNames)
	if err != nil {
		return nil, err
	}

	for _, team := range teams {
		for _, teamPlayer := range team.Players {
			player, ok := playersMap[teamPlayer.ID]
			if !ok {
				id, _ := strconv.Atoi(string(teamPlayer.ID))
				playersMap[teamPlayer.ID] = Player{
					ID:       uint32(id),
					FullName: teamPlayer.FullName(),
					Age:      teamPlayer.Age,
					Teams:    []string{team.Name},
				}
				continue
			}
			player.Teams = append(player.Teams, team.Name)
			playersMap[teamPlayer.ID] = player
		}
	}

	players := utils.GetMapValues(playersMap)
	sort.Slice(players, func(i, j int) bool {
		return players[i].ID < players[j].ID
	})

	return players, nil
}

// getTeams takes a list of team names as an argument and return their data
// if any name not exist in the API, it will return errTeamNotExist error
func getTeams(teamNames []string) ([]Team, error) {
	isTeamInTheList := func(name string) bool {
		return utils.Contains(teamNames, name)
	}

	teams := make([]Team, 0, len(teamNames))
	teamId := 1
	for {
		team, err := getTeam(TeamId(teamId))
		teamId++

		if err != nil {
			if err == errTeamNotFound { // no teams left to get
				break
			} else {
				return teams, err // something went wrong
			}
		}
		if isTeamInTheList(team.Name) {
			teams = append(teams, *team)
		}
		if len(teams) == len(teamNames) {
			break
		}
	}
	if len(teams) != len(teamNames) {
		return nil, errTeamNotExist
	}
	return teams, nil
}

// getTeamsFaster takes a list of team names as an argument and return their data
// if any name not exist in the API, it will return errTeamNotExist error
// this function calls the API in parallel to boost performance
func getTeamsFaster(teamNames []string) ([]Team, error) {
	isTeamInTheList := func(name string) bool {
		return utils.Contains(teamNames, name)
	}

	var (
		totalTeams     = len(teamNames)
		noOfGoRoutines = 40
		teamsCh        = make(chan *Team, totalTeams)
		teamIdCh       = make(chan TeamId, noOfGoRoutines)
		teams          = make([]Team, 0, totalTeams)
		teamId         = 1
		apiScanned     = false
	)
	var errApi error

	routine := func() {
		for id := range teamIdCh {
			if len(teamsCh) == totalTeams {
				break
			}
			team, err := getTeam(id)
			if err != nil {
				if err == errTeamNotFound { // no teams left to get
					apiScanned = true
					break
				}
				errApi = err
				break
			}
			if isTeamInTheList(team.Name) {
				teamsCh <- team
			}
		}
	}

	for i := 0; i < noOfGoRoutines; i++ {
		go routine()
	}

	defer func(ch chan<- TeamId) {
		close(ch)

	}(teamIdCh)

	for {
		teamIdCh <- TeamId(teamId)
		teamId++

		if errApi != nil {
			return nil, errApi
		}
		if apiScanned && (len(teamsCh) != totalTeams) {
			return nil, errTeamNotExist
		}
		if len(teamsCh) == totalTeams {
			for i := 0; i < totalTeams; i++ {
				team := <-teamsCh
				teams = append(teams, *team)
			}
			return teams, nil
		}
	}
}

// getTeam returns the team data given by team id
// if team does not exist against the team id, it will return errTeamNotFound error
func getTeam(teamId TeamId) (*Team, error) {
	res, err := http.Get(fmt.Sprintf(config.CONFIG.TeamApi, teamId))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var teamResponse teamApiResponse
	if err := json.Unmarshal(body, &teamResponse); err != nil {
		return nil, err
	}

	if teamResponse.Code != 0 {
		return nil, errTeamNotFound
	}

	team := teamResponse.Data.TeamData
	return &team, nil
}

// teamApiResponse is struct to hold the data received from the API
type teamApiResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   struct {
		TeamData Team `json:"team"`
	} `json:"data"`
}
