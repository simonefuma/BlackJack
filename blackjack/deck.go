package blackjack

import "math/rand"
import "strconv"
import "time"

// Create a new deck
func NewDeck(numberDecks int) (deck Deck) {
    var seeds []string = []string{"♥", "♦", "♣", "♠"}
    var values []string = fillValues([]string{0:"A", 10:"J", 11:"Q", 12:"K"})

    // Create deck
    for i := 0; i < len(seeds); i++ {
        for j := 0; j < len(values); j++ {
            deck = append(deck, values[j]+seeds[i])
        }
    }

    // Concatenate decks
    for i := 0; i < numberDecks-1; i++ {
        deck = append(deck, deck...)
    }
    return shuffle(deck)
}

// Fill slice with all values
func fillValues(values []string) []string {
    for i := 1; i < 10; i++ {
        values[i] = strconv.Itoa(i+1)
    }
    return values
}

// Shuffle deck
func shuffle(deck Deck) Deck {
    rand.Seed(int64(time.Now().Nanosecond()))
    rand.Shuffle(len(deck), func(i, j int) {
        deck[i], deck[j] = deck[j], deck[i]
    })
    return deck
}
