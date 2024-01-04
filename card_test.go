package deck

import "fmt"

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Spade})
	fmt.Println(Card{Rank: Four, Suit: Heart})
	fmt.Println(Card{Rank: King, Suit: Diamond})
	fmt.Println(Card{Rank: Ten, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output: 
	// Ace of Spades
	// Four of Hearts
	// King of Diamonds
	// Ten of Clubs
	// Joker
}