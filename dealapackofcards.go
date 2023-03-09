package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	block := len(deck) / 4

	for i := 0; i < len(deck); i += block {
		player := i/block + 1
		end := i + block
		if end > len(deck) {
			end = len(deck)
		}

		DisplayCards(player, deck[i:end])
	}
}

func DisplayCards(player int, cards []int) {
	fmt.Printf("Player %d: ", player)
	for i, card := range cards {
		fmt.Printf("%d", card)

		if i != len(cards)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n")
}
