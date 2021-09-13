package main

import "fmt"

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d wins \n", g.WinningPlayer())
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

func NewGameOfChess() Game {
	return &chess{1, 10, 1}
}

func (c *chess) Start() {
	fmt.Println("Chess Starting ")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d\n", c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func main() {
	g := NewGameOfChess()
	PlayGame(g)
}
