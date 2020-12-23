package main

import "fmt"

func main() {

	// cards := []string{"Ace Of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// var cards = deck{
	// 	"Ace Of Diamonds",
	// 	newCard(),
	// 	"Six of Spades"}
	// var cards = newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("Hand------------------------------")
	// hand.print()
	// fmt.Println("Remaining Deck------------------------------")
	// remainingDeck.print()
	// cards.saveToFile("my_cards")
	readCards := newDeckFromFile("my_cards")
	fmt.Println(readCards)
	readCards.shuffle()
	fmt.Println(readCards)

}

func newCard() string {
	return "Five of Diamonds"
}
