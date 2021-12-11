package blackjack

import "strings"
import "strconv"

// Give one card
func (m *Hand) Hit(deck Deck) Deck {
    m.Cards = append(m.Cards, deck[0])
    return deck[1:]
}

// Calculate the new score
func (m *Hand) CalculateScore() {
    // New card value
    lastCard := m.Cards[len(m.Cards)]
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
        softNewScoreInt = softOldScoreInt + softValueInt
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
