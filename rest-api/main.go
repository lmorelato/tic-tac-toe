package main

import (
	"github.com/lmorelato/tic-tac-toe/games"
	"log"
)

func main() {
	log.Println("Starting tic-tac-toe-api v0.1")
	games.HandleRequests()
}
