package main

import "fmt"

type Monster struct {
	Name            string
	Attack, Defense int
}

func NewMonster(name string, attack, defense int) *Monster {
	return &Monster{name, attack, defense}
}

func (c *Monster) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type MonsterModifier struct {
	Monster *Monster
	next    Modifier
}

func NewMonsterModifier() *MonsterModifier {
	return &MonsterModifier{}
}

func (c *MonsterModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *MonsterModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifier struct {
	MonsterModifier
}

func NewDoubleAttackModifier(c *Monster) *DoubleAttackModifier {
	return &DoubleAttackModifier{MonsterModifier: MonsterModifier{Monster: c}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling, ", d.Monster.Name, "\b's attack")
	d.Monster.Attack *= 2
	d.MonsterModifier.Handle()
}

func main() {
	goblin := NewMonster("Goblin", 1, 1)
	bossGoblin := NewMonster("bossGoblin", 10, 20)

	fmt.Println(goblin.String())
	fmt.Println(bossGoblin.String())

	root := NewMonsterModifier()
	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewDoubleAttackModifier(bossGoblin))

	root.Handle()
	root.Handle()

	fmt.Println(goblin.String())
	fmt.Println(bossGoblin.String())

}
