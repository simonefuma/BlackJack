package blackjack


type card string

type Deck []carta

type hand struct {
    name string
    cards []card
    score string
    stand, isBust bool
}

type dealer hand

type player struct {
    hands []hand
}
