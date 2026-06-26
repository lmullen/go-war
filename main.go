package main

import (
	"fmt"

	"github.com/lmullen/go-war/war"
)

var titleGraphic = `
 _       _____    ____  ______
| |     / /   |  / __ \/ / / /
| | /| / / /| | / /_/ / / / / 
| |/ |/ / ___ |/ _, _/_/_/_/  
|__/|__/_/  |_/_/ |_(_|_|_)   
`

func main() {
	fmt.Println("Paul and Daddy's cool game: War!!!")
	fmt.Println(titleGraphic)
	fmt.Println("")
	fmt.Println("A fun card game remade on a computer")
	fmt.Println("")
	fmt.Println("Here is the deck of cards:")

	for _, card := range war.Deck {
		fmt.Printf("%s, ", card)
	}
	fmt.Println()

}
