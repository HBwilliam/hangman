package hangman

import (
	"fmt"
	"strings"
)

// ========================GAME============================//

type Hangman struct {
	Word      []rune
	Tofind    string
	Guesses   []rune
	W_Guesses []string
	Tries     int
}

// Game is the main function of the game
func Game(input string, hang *Hangman) string {

	inputs := input

	if len(inputs) == 1 {
		found := false
		for index, letters := range hang.Tofind {
			if letters == rune(inputs[0]) {
				hang.Word[index] = rune(inputs[0])
				found = true
			}
		}
		if strings.ContainsRune(string(hang.Guesses), rune(inputs[0])) {
			return "Already guessed"
		} else if !found {
			hang.Guesses = append(hang.Guesses, ' ')
			hang.Guesses = append(hang.Guesses, rune(inputs[0]))
			hang.Tries++
		}
	} else if len(inputs) > 1 {
		if inputs == hang.Tofind {
			hang.Word = []rune(inputs)
		} else if strings.Contains(strings.Join(hang.W_Guesses, ""), inputs) {
			return "Already guessed"
		} else {
			hang.W_Guesses = append(hang.W_Guesses, inputs)
			hang.Tries += 2
		}
	}

	if hang.Tries >= 10 {
		return "Lose"
	} else if string(hang.Word) == hang.Tofind {
		return "Win"
	}
	return "Continue"
}

func Reload(hang *Hangman, file string) {
	dico, err := Reader_word(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	Init_game(hang, dico)
}
