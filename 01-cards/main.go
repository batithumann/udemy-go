package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println(" ")
	// remainingCards.print()
	// fmt.Print(cards.toString())
	cards.saveToFile("cards.txt")
	newCards := newDeckFromFile("cards.txt")
	newCards.shuffle()
	_, newCards = deal(newCards, 5)
	hand, newCards := deal(newCards, 5)
	hand.print()

}
