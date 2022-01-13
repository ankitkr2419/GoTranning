package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "elephant"
	entries := map[string]bool{}
	placeholder := []string{}
	for i := 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}
	fmt.Println("Welcome to hangman game ")
	chances := 8
	for {
		arr := []string{}

		fmt.Println(strings.Join(arr, ", "))
		fmt.Println("Lives left ", chances)
		fmt.Println(placeholder)
		userInput := strings.Join(placeholder, "")
		if chances == 0 && userInput != word {
			fmt.Println("Game Over Try again")
			break
		}
		if strings.Compare(word, userInput) == 0 {
			fmt.Print("You won")
			break
		}
		var input string
		fmt.Println("Enter the word :: ")
		fmt.Scanln(&input)

		if !entries[input[0:1]] {
			entries[input[0:1]] = true
			if strings.Contains(word, input[0:1]) {
				// append(arr, input[0:1])
				firstIndex := strings.Index(word, input)
				for i := firstIndex; i < len(word); i++ {
					if string(word[i]) == input[0:1] {
						placeholder[i] = input[0:1]
					}
				}
			}
		}
		if !strings.Contains(word, input[0:1]) {
			chances--
		}
	}
}
