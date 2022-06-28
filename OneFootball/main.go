package main

import (
	"fmt"
	"strings"

	"github.com/NajeebTyson/onefootball/internal/config"
	"github.com/NajeebTyson/onefootball/pkg/onefootball"
)

func main() {
	config.LoadConfig("config.json")

	fmt.Printf("Getting players of team: %v\n", config.CONFIG.Teams)

	players, err := onefootball.GetPlayers(config.CONFIG.Teams)
	if err != nil {
		panic(err)
	}

	printPlayers(players)
}

func printPlayers(players []onefootball.Player) {
	for i, player := range players {
		fmt.Printf("%d. %d; %s; %s; %s\n", i+1, player.ID, player.FullName, player.Age, strings.Join(player.Teams, `, `))
	}
}
