package main

import "fmt"

type deck []string

func (d deck) printcards() {

	for ind, single_card := range d {
		fmt.Println(ind, single_card)
	}
}
