package hangman

import (
	"math/rand"
)

// =========================================================//
// =======================INIT_GAME==========================//

// Init_game initialize the game
func Init_game(hangman *Hangman, word []string) {
	hangman.Word = []rune{}
	hangman.Tofind = ""
	hangman.Guesses = []rune{}
	hangman.W_Guesses = []string{}

	random := rand.Intn(len(word))
	hangman.Tofind = word[random]

	visableletters := len(hangman.Tofind)/2 - 1

	for range hangman.Tofind {
		hangman.Word = append(hangman.Word, '_')
	}

	again := false
	var place []int

	// search visable letters in the word
	for visableletters > 0 {
		random = rand.Intn(len(hangman.Tofind))
		for _, j := range place {
			if j == random {
				again = true
			}
		}
		if !again {
			place = append(place, random)
			visableletters--
		} else {
			again = false
		}
	}

	// places the visable letters in the word
	WordRune := []rune(hangman.Tofind)
	for _, index := range place {
		hangman.Word[index] = WordRune[index]
		hangman.Guesses = append(hangman.Guesses, WordRune[index])
	}
	for i := range hangman.Guesses {
		for index, letters := range hangman.Tofind {
			if letters == hangman.Guesses[i] {
				hangman.Word[index] = letters
			}
		}
	}

	hangman.Tries = 0
}
