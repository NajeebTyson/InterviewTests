package onefootball

import "fmt"

type TeamId uint32
type PlayerId string

// Team is a data structer to hold the team data got from API
type Team struct {
	ID      TeamId       `json:"id"`
	Name    string       `json:"name"`
	Players []TeamPlayer `json:"players"`
}

// TeamPlayer is a data structure to hold the player data, got from API
type TeamPlayer struct {
	ID        PlayerId `json:"id"`
	Firstname string   `json:"firstName"`
	Lastname  string   `json:"lastName"`
	Name      string   `json:"name"`
	Age       string   `json:"age"`
}

// FullName returns the full name of the player
func (p *TeamPlayer) FullName() string {
	if len(p.Firstname) != 0 && len(p.Lastname) != 0 {
		return fmt.Sprintf("%s %s", p.Firstname, p.Lastname)
	} else if len(p.Firstname) != 0 {
		return p.Firstname
	} else if len(p.Lastname) != 0 {
		return p.Lastname
	}
	return p.Name
}

// Player is a data structure to hold the player data in the application
type Player struct {
	ID       uint32
	FullName string
	Age      string
	Teams    []string
}
