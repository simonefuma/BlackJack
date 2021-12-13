package blackjack

import "strings"
import "os/exec"
import "runtime"
import "fmt"
import "os"

var clear map[string]func() //create a map for storing clear funcs

func PrintTable(player Player, dealer Hand) {
    var table []string
    var b strings.Builder

    // Add dealer string
    table = append(table, dealer.Name+": "+colorSeeds(dealer.Cards))

    // Add player string
    for _, hand := range player.Hands {
        table = append(table, hand.Name+": "+colorSeeds(hand.Cards))
    }

    // Get the longest string
    longestString := 0
    for _, handString := range table {
        lenghtHandString := getLenght(handString)
        if lenghtHandString > longestString {
            longestString = lenghtHandString
        }
    }

    // Print table
    for i, handString := range table {
        lenghtHandString := getLenght(handString)
        b.WriteString(handString)
        for lenghtHandString < longestString {
            b.WriteString(" ")
            lenghtHandString++
        }
        b.WriteString("\tScore: ")

        if i == 0 {
            b.WriteString(dealer.Score)
        } else {
            b.WriteString(player.Hands[i-1].Score)
        }
        fmt.Println(b.String())
        b.Reset()
    }
}

// Color cards seeds
func colorSeeds(cards []Card) string {
    var b strings.Builder
    for _, card := range cards {
        if card == "--" {
            b.WriteString("--")
            break
        }
        value := card[:len(card)-3]
        seed := card[len(card)-3:]

        b.WriteString(string(value))

        // Color seed
        if seed == "♥" || seed == "♦" {
            b.WriteString("\033[31m")
        } else {
            b.WriteString("\033[30m")
        }
        b.WriteString(string(seed)+"\033[0m ")
    }
    return b.String()
}

// Return the lenght of a string
func getLenght(s string) (lenght int) {
    for _, element := range s {
        lenght++
        if element == '♥' || element == '♦' || element == '♣' || element == '♠' {
            lenght -= 9
        }
    }
    return
}

// Function to clear shell
//******************************************************************************
func Init() {
	clear = make(map[string]func()) //Initialize it
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") // Mac
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc. darwin = mac
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
//******************************************************************************
