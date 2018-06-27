package main

import "fmt"

type deck []string

func (d deck) print() {
	for index, deck := range d {
		fmt.Println(index, deck)
	}
}
