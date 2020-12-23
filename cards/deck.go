package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func (cards deck) print() {

	for index, card := range cards {
		fmt.Println(index+1, card)
	}
}

func (cards deck) toString() string {
	fullText := strings.Join(cards, ",")
	return fullText
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Tree", "Four"}

	for cardSuitIndex := 0; cardSuitIndex < len(cardSuits); cardSuitIndex++ {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuits[cardSuitIndex])
		}
	}

	return cards
}

func (cards deck) saveToFile(fileName string) error {
	fullText := cards.toString()
	willWriteData := []byte(fullText)
	error := ioutil.WriteFile(fileName, willWriteData, 0666)
	return error
}

func newDeckFromFile(fileName string) deck {
	readData, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error : ", error)
		os.Exit(1)
	}

	readString := string(readData)
	deckStringList := strings.Split(readString, ",")
	return deck(deckStringList)
}

func (cards deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(source)

	for index := range cards {
		newPosition := newRand.Intn(len(cards))
		cards[index], cards[newPosition] = cards[newPosition], cards[index]
	}

}
