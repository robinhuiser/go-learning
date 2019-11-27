package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type based upon slice of string
type deck []string

// New deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal a hand
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// Convert the deck to a sting
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Load from file
func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// We print the error and quit with errorcode 1
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)

}

// Shuffle!
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		ri := r.Intn(len(d) - 1)
		d[i], d[ri] = d[ri], d[i]
	}
}
