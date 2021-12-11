package blackjack


type Card string

type Deck []Card

type Hand struct {
    Name string
    Cards []Card
    Score string
    Stand, Bust bool
}

type Dealer Hand

type Player struct {
    hands []Hand
}
