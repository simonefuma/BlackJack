package blackjack

import "strings"
import "strconv"

// Create a new player
func NewPlayer(name string) Player {
    return Player{[]Hand{Hand{name, []Card{}, "", false, false}}}
}

// Create a new dealer
func NewHand(name string) Hand {
    return Hand{name, []Card{}, "", false, false}
}

// Give one card
func (m *Hand) Hit(deck Deck) Deck {
    m.Cards = append(m.Cards, deck[0])
    m.CalculateScore()
    if len(m.Cards) > 2 {
        m.IsBust()
    }
    return deck[1:]
}

// Give the first 4 cards
func Distribute(player, dealer *Hand, deck Deck) (Card, Deck) {
    deck = player.Hit(deck)
    deck = dealer.Hit(deck)
    deck = player.Hit(deck)
    dealer.Cards = append(dealer.Cards, Card("--"))
    return deck[0], deck[1:]
}

// Calculate the new score
func (m *Hand) CalculateScore() {
    // New card value
    lastCard := m.Cards[len(m.Cards)-1]
    hardValueInt, softValueInt := giveValue(lastCard)
    hardValueString, softValueString := strconv.Itoa(hardValueInt), strconv.Itoa(softValueInt)

    // First card
    if m.Score == "" {
        if softValueInt == 0 {
            m.Score = hardValueString
        } else {
            m.Score = hardValueString+"/"+softValueString
        }
        return
    }

    // Old score value
    oldScore := strings.Split(m.Score, "/")
    hardOldScoreInt, _ := strconv.Atoi(oldScore[0])
    softOldScoreInt := 0
    if len(oldScore) == 2 {
        softOldScoreInt, _ = strconv.Atoi(oldScore[1])
    }

    // Calculate new hard score
    hardNewScoreInt := hardOldScoreInt + hardValueInt
    hardNewScoreString := strconv.Itoa(hardNewScoreInt)
    softNewScoreInt := 0
    softNewScoreString := ""

    // Calculate new soft score
    if softOldScoreInt == 0 && softValueInt != 0 {
        softNewScoreInt = hardOldScoreInt + softValueInt
        softNewScoreString = strconv.Itoa(softNewScoreInt)
        m.Score = hardNewScoreString+"/"+softNewScoreString
    } else if softOldScoreInt != 0 && softValueInt == 0 {
        softNewScoreInt = softOldScoreInt + hardValueInt
        softNewScoreString = strconv.Itoa(softNewScoreInt)
        m.Score = hardNewScoreString+"/"+softNewScoreString
    } else {
        m.Score = hardNewScoreString
    }

    // Check if it's BlackJack
    if len(m.Cards) == 2 && softNewScoreInt == 21 {
        m.Score = "BlackJack"
    }
}

// Give the value of the card
func giveValue(card Card) (hardValue, softValue int) {
    valueString := card[:len(card)-3]
    switch valueString {
    case "A":
        return 1, 11
    case "J", "Q", "K":
        return 10, 0
    default:
        cardValue, _ := strconv.Atoi(string(valueString))
        return cardValue, 0
    }
}

// Check if it's bust
func (m *Hand) IsBust() {
    score := strings.Split(m.Score, "/")

    // hard score
    hardScoreInt, _ := strconv.Atoi(score[0])
    if hardScoreInt > 21 {
        m.Score = "Bust"
        m.Bust = true
        return
    }

    // soft score
    if len(score) == 2 {
        softScoreInt, _ := strconv.Atoi(score[1])
        if softScoreInt > 21 {
            m.Score = score[0]
        }
    }
}
