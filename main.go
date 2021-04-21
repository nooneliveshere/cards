package main

func main() {
	card := deck{"ace of diaomonds", newCard()}
	card = append(card, "six of spades")
	card.printcards()
}

func newCard() string {
	return "Five of diamonds/пять бубей"
}
