package main

import (
	"fmt"
	"strings"
)

func (d deck) toString() string {
	return strings.Join([]string(d), " ")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
