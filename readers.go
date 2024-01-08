package hangman

import (
	"bufio"
	"os"
)

//=======================READER_WORD==========================//

// Reader_word reads the file with the words
func Reader_word(file_word string) ([]string, error) {
	var word []string

	readFile, err := os.Open(file_word)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile) // Creates a scanner to read the file.

	fileScanner.Split(bufio.ScanLines) // Divides the file into lines.

	for fileScanner.Scan() {
		// Adds the text of the current line to the FindWord slice.
		word = append(word, fileScanner.Text())
	}

	readFile.Close()

	return word, nil
}

// =========================================================//
