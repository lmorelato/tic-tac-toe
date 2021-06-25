package main

import (
	"log"
	"github.com/lmorelato/tic-tac-toe/games"
)


func main() {
	log.Println("Starting tic-tac-toe-api v0.1")
	games.HandleRequests()
}
