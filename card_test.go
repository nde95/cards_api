package deck

import (
	"fmt"
	"testing"
)

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

func TestNew (t *testing.T) {
	cards := NewCard()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong amount of cards created")
	}
}

func TestDefaultSort (t *testing.T) {
	cards := NewCard(DefaultSort)
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("Expected Ace of Spades as first result, instead got:", cards[0])
	}
}

func TestSort (t *testing.T) {
	cards := NewCard(Sort(Less))
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("Expected Ace of Spades as first result, instead got:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := NewCard(Jokers(3))
	count := 0 
	for _, c:= range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 jokers, instead got:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three 
	}
	cards := NewCard(Filter(filter))
	for _, c:= range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all Twos and Threes to be filtered out of the deck")
		}
	}
}