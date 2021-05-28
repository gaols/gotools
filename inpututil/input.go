package inpututil

import (
	"bufio"
	"fmt"
	"github.com/gaols/gotools"
	"os"
)

func Confirm(prompt string) {
	fmt.Print(prompt + " > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if gotools.IsEqualsAny(input, []string{"yes", "y", "Y"}...) {
			break
		} else if gotools.IsEqualsAny(input, []string{"no", "n", "N"}...) {
			os.Exit(0)
		} else {
			fmt.Println("invalid input")
			fmt.Print("> ")
		}
	}
}

func MakeChoice(prompt string, choices []string) (retChoiceIdx int, retChoice string) {
	fmt.Println(prompt)
	validChoices := make([]string, 0)
	for i, choice := range choices {
		fmt.Printf("%s. %s\n", gotools.LeftPad(gotools.Int2Str(i+1), 2, ' '), choice)
		validChoices = append(validChoices, gotools.Int2Str(i+1))
	}
	fmt.Print("Make your choice > ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if gotools.IsEqualsAny(input, validChoices...) {
			idx, _ := gotools.Str2Int(input)
			retChoiceIdx = idx - 1
			retChoice = choices[idx-1]
			return
		} else if input == "q" {
			os.Exit(0)
		} else {
			fmt.Println("invalid choice, press q to quit")
			fmt.Print("> ")
		}
	}
	os.Exit(1)
	return
}
