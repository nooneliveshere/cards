package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Пики", "Буби", "Черви", "Крести"}
	cardValues := []string{"Туз", "Король", "Дама", "Валет"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suit)
		}
	}
	return cards
}
func (d deck) printcards() {
	for i, cardoid := range d {
		fmt.Println(i, cardoid)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
