package main

type Animal struct {
	name string
}

type Dog Animal

func main() {
	cards := newDeck()
	hand, remindCards := deal(cards, 5)
	hand.print()
	remindCards.print()

	// greeting := "Hi there !"
	// fmt.Println([]byte(greeting))

}
