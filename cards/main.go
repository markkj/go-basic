package main

type Animal struct {
	name string
}

type Dog Animal

func main() {
	// cards := newDeck()
	// hand, remindCards := deal(cards, 5)
	// hand.print()
	// remindCards.print()

	// hand.saveToFile("inHand.txt")

	// greeting := "Hi there !"
	// fmt.Println([]byte(greeting))

	cards := loadFromFile("inHand.txt")
	cards.print()

}
