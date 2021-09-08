package main

import (
	"fmt"
	"sync"
)

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	MonsterName string
	WhatToQuery Argument
	Value       int
}

type Obsever interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Obsever)
	UnSubscribe(o Obsever)
	Fire(q *Query)
}

type Game struct {
	observer sync.Map
}

func (g *Game) Subscribe(o Obsever) {
	g.observer.Store(o, struct{}{})
}

func (g *Game) UnSubscribe(o Obsever) {
	g.observer.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Obsever).Handle(q)
		return true
	})

}

type Monster struct {
	game            *Game
	Name            string
	attact, defense int
}

func NewMonster(game *Game, name string, attack, defense int) *Monster {
	return &Monster{game: game, Name: name, attact: attack, defense: defense}
}

func (m *Monster) Attack() int {
	q := Query{m.Name, Attack, m.attact}
	m.game.Fire(&q)
	return q.Value
}

func (m *Monster) Defense() int {
	q := Query{m.Name, Defense, m.defense}
	m.game.Fire(&q)
	return q.Value
}

func (m *Monster) String() string {
	return fmt.Sprintf("%s (%d/%d)", m.Name, m.Attack(), m.Defense())
}

type MonsterModifier struct {
	game    *Game
	monster *Monster
}

func (m *MonsterModifier) Handle(q *Query) {

}

type DoubleAttackModifier struct {
	MonsterModifier
}

func NewDoubleAttackModifier(game *Game, monster *Monster) *DoubleAttackModifier {
	d := &DoubleAttackModifier{MonsterModifier: MonsterModifier{game: game, monster: monster}}
	game.Subscribe(d)
	return d
}

func (m *DoubleAttackModifier) Handle(q *Query) {
	if q.MonsterName == m.monster.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (m *DoubleAttackModifier) Close() error {
	m.game.UnSubscribe(m)
	return nil
}

func main() {
	game := &Game{sync.Map{}}
	goblin := NewMonster(game, "Goblin", 1, 1)
	// bossGoblin := NewMonster(game, "bossGoblin", 10, 20)
	fmt.Println(goblin.String())
	// fmt.Println(bossGoblin.String())

	m := NewDoubleAttackModifier(game, goblin)
	fmt.Println(goblin.String())
	m.Close()
	fmt.Println(goblin.String())

}
