package main

func main() {
	cards := newDeck()
	//cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("<<<<Hand>>>>>>")
	// hand.print()
	// fmt.Println("<<<<Remaining Deck>>>>")
	// remainingDeck.print()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
