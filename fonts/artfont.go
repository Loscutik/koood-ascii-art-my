package fonts

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	SYMBOL_HEIGHT = 8
	FIRST_SYMBOL  = ' '
	LAST_SYMBOL   = '~'
)

type (
	ArtFont   = map[rune]ArtString
	ArtString = [SYMBOL_HEIGHT]string
)

/*
reads a file with an ascii graphic fonts representing characters from ' '(space) to '~' and returns a map which keeps the characters
maps's key is the character (type rune), map's element is a value of the type ArtString= [SYMBOL_HEIGHT]string
*/
func GetArtFont(fileName string) (aFont ArtFont, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	aFont = make(map[rune]ArtString)
	scanner := bufio.NewScanner(file)

	// reading all characters to the map aFont
	for char := (FIRST_SYMBOL); char <= LAST_SYMBOL; char++ {
		scanner.Scan() // skip the first empty line

		aFont[char], err = readArtChar(scanner)
		if err != nil {
			err = fmt.Errorf("error reading the character from a file: %s; character: %c; file:  %s", err, char, fileName)
			break
		}
	}

	return
}

/*
reads the next graphic character from given scanner, will return an error if scanner.Scan finishes with an error
*/
func readArtChar(scanner *bufio.Scanner) (aChar ArtString, err error) {
	for line := 0; line < SYMBOL_HEIGHT; line++ {
		if scanner.Scan() {
			aChar[line] = (scanner.Text())
		} else if err = scanner.Err(); err != nil {
			break
		} else {
			err = fmt.Errorf("unexpected break of scanning")
			break
		}
	}
	return aChar, err
}
