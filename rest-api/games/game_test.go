package games

import (
	"testing"
	"time"
)

func TestNewGame(t *testing.T) {
	// arrange
	name := "player1"
	now := time.Now()

	// act
	g := NewGame(name)

	// assert
	if len(g.SessionId) != SESSION_ID_SIZE {
		t.Errorf("Expected len of %d, but got len of %d", SESSION_ID_SIZE, len(g.SessionId))
	}
	if g.Player1.Name != name {
		t.Errorf("Expected %s, but got %s", name, g.Player1.Name)
	}
	if g.Player1.Symbol != SYMBOL_X {
		t.Errorf("Expected %s, but got %s", SYMBOL_X, g.Player1.Symbol)
	}
	if g.Start.Before(now) {
		t.Errorf("Expected to be greater than %s, but got %s", now.String(), g.Start.String())
	}
}

func TestSetPlayer2(t *testing.T) {
	// arrange
	name := "player2"
	g := NewGame("player1")

	// act
	g.SetPlayer2(name)

	// assert
	if g.Player2.Name != name {
		t.Errorf("Expected %s, but got %s", name, g.Player2.Name)
	}
	if g.Player2.Symbol != SYMBOL_O {
		t.Errorf("Expected %s, but got %s", SYMBOL_O, g.Player2.Symbol)
	}
}

func TestPlay_NoWinner(t *testing.T) {
	// arrange
	player1 := "player1"
	player2 := "player2"
	g := NewGame(player1)
	g.SetPlayer2(player2)

	// act
	g.Play(SYMBOL_X, 0)
	g.Play(SYMBOL_O, 3)
	g.Play(SYMBOL_X, 1)
	g.Play(SYMBOL_O, 4)
	g.Play(SYMBOL_X, 6)
	g.Play(SYMBOL_O, 7)
	g.Play(SYMBOL_X, 5)
	g.Play(SYMBOL_O, 2)
	g.Play(SYMBOL_X, 8)

	// assert
	if g.Winner != NO_WINNER {
		t.Errorf("Expected %s, but got %s", NO_WINNER, g.Winner)
	}
}

func TestPlay_Player1Wins(t *testing.T) {
	// arrange
	player1 := "player1"
	player2 := "player2"
	g := NewGame(player1)
	g.SetPlayer2(player2)

	// act (diagonal win of X)
	g.Play(SYMBOL_X, 0)
	g.Play(SYMBOL_O, 3)
	g.Play(SYMBOL_X, 4)
	g.Play(SYMBOL_O, 6)
	g.Play(SYMBOL_X, 8)

	// assert
	if g.Winner != player1 {
		t.Errorf("Expected %s, but got %s", NO_WINNER, g.Winner)
	}
}

func TestPlay_Player2Wins(t *testing.T) {
	// arrange
	player1 := "player1"
	player2 := "player2"
	g := NewGame(player1)
	g.SetPlayer2(player2)

	// act (row win of O)
	g.Play(SYMBOL_X, 0)
	g.Play(SYMBOL_O, 3)
	g.Play(SYMBOL_X, 1)
	g.Play(SYMBOL_O, 4)
	g.Play(SYMBOL_X, 8)
	g.Play(SYMBOL_O, 5)

	// assert
	if g.Winner != player2 {
		t.Errorf("Expected %s, but got %s", NO_WINNER, g.Winner)
	}
}
