package main

import "fmt"

func main() {
	card := []string{"ace of diaomonds", newCard()}
	card = append(card, "six of spades")
	for ind, single_card := range card {
		fmt.Println(ind, single_card)
	}

}

func newCard() string {
	return "Five of diamonds/пять бубей"
}
