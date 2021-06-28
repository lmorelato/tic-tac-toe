package games

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type Game struct {
	SessionId   string    `json:"sessionId"`
	Player1     Player    `json:"player1"`
	Player2     Player    `json:"player2"`
	Moves       [9]string `json:"moves"`
	Logs        []Log     `json:"logs"`
	Start       time.Time `json:"start"`
	Duration    int64     `json:"duration"`
	Winner      Player    `json:"winner"`
	WinningPath [3]int    `json:"winningPath"`
	NextSymbol  string    `json:"nextSymbol"`
}

type Player struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Log struct {
	Index  int `json:"index"`
	Order  int `json:"order"`
	Player `json:"player"`
}

type PlayerMove struct {
	Index  int    `json:"index"`
	Symbol string `json:"symbol"`
}

type GamesMap map[string]*Game

/*
	one simple string array will store the game moves,
	so we will have the following game board design:
		0|1|2
		3|4|5
		6|7|8

	with that design in mind we can set the winning paths,
	according to the tic-tac-toe rules
*/
var winningPaths = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rols
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // cols
	{0, 4, 8}, {2, 4, 6}} // diagonals

// adds basic info to start the game
func NewGame(name string) Game {
	g := Game{
		SessionId:  newSessionId(),
		Player1:    Player{Name: name, Symbol: SYMBOL_X},
		Start:      time.Now(),
		NextSymbol: "X",
	}

	log.Println("New game created: ", g)
	return g
}

// generates a random string to be used as a sessionId
func newSessionId() string {
	rand.Seed(time.Now().UnixNano())
	id := make([]byte, SESSION_ID_SIZE)

	for i := 0; i < SESSION_ID_SIZE; i++ {
		id[i] = SESSION_ID_CHARS[rand.Intn(len(SESSION_ID_CHARS))]
	}

	return string(id)
}

func (g *Game) SetPlayer2(name string) {
	// if we have a player 2 we don't need to do anything
	if g.Player2.Name != "" {
		log.Println("We already have a player 2!")
		return
	}

	log.Println("PLayer 2 has entered in the game: ", name)
	(*g).Player2 = Player{Name: name, Symbol: SYMBOL_O}
}

// adds a new move and check the board for a winner or a tie
func (g *Game) Play(symbol string, index int) {
	// checks if the current symbol is the one we are expecting
	// otherwise we do not allow this move
	if g.NextSymbol != symbol {
		log.Println("Move not allowed, expected symbol", g.NextSymbol)
		return
	}

	// if we have a winner we don't need to do anything
	if g.Winner.Name != "" {
		log.Println("This game is already over!")
		return
	}

	name := g.getPlayerNameBySymbol(symbol)
	g.addPlayerMove(index, name, symbol)

	if g.checkWinner(symbol) {
		g.setWinner(name, symbol)
	} else if len(g.Logs) == 9 {
		// if we have 9 moves and no winner so it is a tie!
		log.Println("Oh no it is a tie!")
		g.setWinner(NO_WINNER, "")
	}
}

// retrieves the player name by his symbol,
// because the players can have the same name
func (g Game) getPlayerNameBySymbol(symbol string) string {
	name := g.Player1.Name
	if symbol == g.Player2.Symbol {
		name = g.Player2.Name
	}

	return name
}

func (g *Game) addPlayerMove(index int, name string, symbol string) {
	(*g).Moves[index] = symbol

	(*g).NextSymbol = SYMBOL_X
	if symbol == SYMBOL_X {
		(*g).NextSymbol = SYMBOL_O
	}

	// logging the move
	l := Log{
		Index: index,
		Order: len((*g).Logs),
		Player: Player{
			Name:   name,
			Symbol: symbol}}
	(*g).Logs = append((*g).Logs, l)

	log.Println("New game move:", l)
}

func (g *Game) checkWinner(symbol string) bool {
	log.Println("Looking for a winner")

	// if we have less than 5 moves we don't have any winner yet
	if len(g.Logs) < 5 {
		log.Println("Not enough moves yet!")
		return false
	}

	// checks if any winning path is satisfied by symbol
	for _, p := range winningPaths {
		if g.Moves[p[0]] == symbol &&
			g.Moves[p[1]] == symbol &&
			g.Moves[p[2]] == symbol {
			log.Println("Wining path found for symbol", symbol, p)
			g.WinningPath = p
			return true
		}
	}

	log.Println("We don't have any winner yet!")
	return false
}

func (g *Game) setWinner(name string, symbol string) {
	// logs the elapsed since the start of the game
	(*g).Duration = time.Since(g.Start).Milliseconds()
	(*g).Winner = Player{Name: name, Symbol: symbol}

	log.Println("The winner is", g.Winner)
	log.Println("Game duration", g.Duration)
}

func LoadGameFromJson(source string) Game {
	var g Game
	json.Unmarshal([]byte(source), &g)
	return g
}
