package onefootball

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/NajeebTyson/onefootball/internal/config"
	"github.com/NajeebTyson/onefootball/pkg/utils"
)

func setup() {
	config.CONFIG.TeamApi = "https://api-origin.onefootball.com/score-one-proxy/api/teams/en/%d.json"
}

func TestGetPlayers(t *testing.T) {
	setup()
	testCases := []struct {
		teams []string
		err   error
	}{
		{teams: []string{"Apoel FC"}, err: nil},
		{teams: []string{"Germany"}, err: nil},
		{teams: []string{"Spain", "Real Madrid"}, err: nil},
		{teams: []string{"Brazil"}, err: nil},
		{teams: []string{"abcd"}, err: errTeamNotExist},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(strings.Join(tc.teams, "_"), func(t *testing.T) {
			t.Parallel()
			players, err := GetPlayers(tc.teams)
			if tc.err != err {
				t.Fatalf("expected error: %v, got error: %v", tc.err, err)
			}
			for _, player := range players {
				for _, team := range player.Teams {
					if !utils.Contains(tc.teams, team) {
						t.Errorf("expected player team in the list %v, got team: %s", tc.teams, team)
					}
				}
			}
		})
	}
}

func TestGetTeam(t *testing.T) {
	setup()
	fmt.Println("api: ", config.CONFIG.TeamApi)
	tt, err := getTeam(0)
	fmt.Printf("team: %+v", tt)
	if err == nil {
		t.Error("expected error but got nil", tt)
	}

	team, err := getTeam(1)
	if err != nil {
		t.Error("expected team but got error: ", err)
	}

	if len(team.Players) == 0 {
		t.Error("expected team players, but got zero players")
	}
}

func TestGetTeamsFaster(t *testing.T) {
	setup()
	testCases := []struct {
		input []string
		err   error
	}{
		{
			input: []string{"Spain", "Real Madrid", "Germany"},
			err:   nil,
		},
		{
			input: []string{"abcde"},
			err:   errTeamNotExist,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(strings.Join(testCase.input, "_"), func(t *testing.T) {
			t.Parallel()
			fmt.Println("get teams: ", tc.input)
			teams, err := getTeamsFaster(tc.input)
			if err != tc.err {
				t.Fatalf("expected error: %v, got error: %v", tc.err, err)
			}
			if tc.err != nil {
				return
			}
			if len(tc.input) != len(teams) {
				t.Fatalf("expected %d teams, but got %d teams", len(tc.input), len(teams))
			}
			for _, team := range teams {
				if !utils.Contains(tc.input, team.Name) {
					t.Fatalf("got team: %s which is not expected", team.Name)
				}
			}
		})
	}
}

func TestGetTeamsBenchmark(t *testing.T) {
	setup()
	teamsNames := []string{"Spain", "Real Madrid", "Germany", "Brazil", "France", "England", "Barcelona"}

	t1 := time.Now()
	time.Sleep(1 * time.Second)
	teams, err := getTeamsFaster(teamsNames)
	fasterTime := time.Since(t1)
	if err != nil {
		t.Fatal(err)
	}
	if len(teams) == 0 {
		t.Fatal("expected teams but got nil")
	}

	t2 := time.Now()
	teams, err = getTeams(teamsNames)
	normalTime := time.Since(t2)
	if err != nil {
		t.Fatal(err)
	}
	if len(teams) == 0 {
		t.Fatal("expected teams but got nil")
	}

	if fasterTime > normalTime {
		t.Fatalf("expected faster method %v to be faster than normal method %v", fasterTime, normalTime)
	}

	t.Logf("faster time: %v, normal time: %v", fasterTime, normalTime)

}
