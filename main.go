package main

func main() {
	carder := newDeck()
	carder = append(carder, "ace of spades")
	//	carder.printcards()

	hand, remainCards := deal(carder, 2)
	hand.printcards()
	remainCards.printcards()
}
