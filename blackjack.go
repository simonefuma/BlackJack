package main

import "github.com/simonefuma/BlackJack/blackjack"
import "fmt"

/*

Classico, Europeo, Spagnolo 21, Vegas Strip, Atlantic City, Blackjack Switch, Multi-mano

1- aggiungere azioni (stand, hit, double, split)
2- aggiungere assicurazione
3- aggiungere puntate e soldi
4- aggiungere stampa e clear della shell
5- aggiungere varianti e regole
6- aggiungere un controllore per le statistiche
7- aggiungere un contatore delle carte (varie metodologie)
8- in caso di blackjack o 21 stand automaticamente

*/

func main() {
    blackjack.Init()
    blackjack.CallClear()

    deck := blackjack.NewDeck(6)
    dealer := blackjack.NewHand("Dealer")
    player := blackjack.NewPlayer("Hand 1")

    hiddenCard, deck := blackjack.Distribute(&player.Hands[0], &dealer, deck)
    _ = hiddenCard


    blackjack.PrintTable(player, dealer)

    deck = player.Hands[0].Hit(deck)
    fmt.Println()

    blackjack.PrintTable(player, dealer)

    deck = player.Hands[0].Hit(deck)
    fmt.Println()

    blackjack.PrintTable(player, dealer)
}
