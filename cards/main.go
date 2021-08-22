package main

func main() {
	cards := newDeck()
	// hand, remindCards := deal(cards, 5)
	cards.print()
	cards.shuffle()
	cards.print()
	// remindCards.print()

	// hand.saveToFile("inHand.txt")

	// greeting := "Hi there !"
	// fmt.Println([]byte(greeting))

	// cards := loadFromFile("inHand.txt")
	// cards.print()
	// cards.shuffle()
	// cards.print()

}
