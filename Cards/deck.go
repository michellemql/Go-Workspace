package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// convert the d to []string type: []string(d)
	// join each string in []stirng into a big stirng by adding "," in between
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// ioutile.ReadFile() will return a []byte slice and an error
	// bs = byte slice, err = error
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

/* shuffle() - shuffle the deck of cards
 * Note: need to generate different randomized sequence of number every single time,
 *       so need to get different seed source every time, and generate random numbers
 *       based on it.
 * Solution: We can use the local time as the seed source because everytime we start the program, the local time will be different.
 * Step 1. generate different seed source everytime
        * source := rand.NewSource(time.Now().UnixNano())
			- rand.NewSource(int64) will genreate new seed source everytime
		  	  based on a int64 number passed in as parameter
			- time.Now().UnixNano() will convert the local time to a int64 number sequence,
		  	  which will then pass into rand.NewSource()
		* r := rand.New(source)
			- generate a random sequence based on the source
	Step 2. Loop through the deck, for each card,
	        generate a random index from range 0-len(deck)
			swap the current card with the card at index_random
			keep going until all cards have been shuffled.

*/
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
