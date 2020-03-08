package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// A new type of variable
// which is nothing but a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//Reciever
// d for deck is the convention of reciver for go.
// In go recievers need to be declared with a single or two lettered variable...
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Deal function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Helper function for converting deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Save to file function
func (d deck) saveToFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error occured")
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	for i := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
