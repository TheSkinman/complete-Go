package main

func main() {
	cards := newDeck()

	//cards.saveToFile("MyCards")

	//cards := newDeckFromFile("MyCards")
	cards.print()
	cards.shuffle()
	cards.print()
}
