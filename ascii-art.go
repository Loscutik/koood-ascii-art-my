package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	ARGS          = 1
	F_STANDART    = "standard.txt"
	SYMBOL_HEIGHT = 8
)

type (
	ArtChar = [SYMBOL_HEIGHT]string
	ArtFont   = map[rune]ArtChar
	ArtString = [SYMBOL_HEIGHT]string
)
//TODO add handles of errors
func main() {
	if len(os.Args) != ARGS+1 {
		log.Fatalf("the programms needs the strict %d argument", ARGS)
	}

	// get the beautiful font for art-printing
	aFont, err := GetArtFont("fonts/" + F_STANDART)
	if err != nil {
		log.Fatalln(err)
	}
	/*test print - whole font */
	// for char := (' '); char <= '~'; char++ {
	// 	for i := 0; i < SYMBOL_HEIGHT; i++ {
	// 		fmt.Print(char)
	// 		fmt.Print(": ")
	// 		fmt.Println((aFont[char][i]))
	// 	}
	// }

	// strs := bytes.Split([]byte(os.Args[1]), []byte("\n"))
	// for _, str := range strs {
	// 	artStr := stringToArt(str, artFont)
	// 	artPrint(artStr)
	// }
}
/* 
reads a file with an ascii graphic fonts representing characters from ' '(space) to '~' and returns a map which keeps the characters
maps's key is the character (type rune), map's element is a value of the type ArtChar= [SYMBOL_HEIGHT]string
*/
func GetArtFont(fileName string) (aFont ArtFont, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	aFont = make(map[rune]ArtChar)
	scanner := bufio.NewScanner(file)

	//reading all characters to the map aFont
	for char := (' '); char <= '~'; char++ {
		scanner.Scan() //skip the first empty line

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
func readArtChar(scanner *bufio.Scanner) (aChar ArtChar, err error) {
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

// func stringToArt(str []rune, afont ArtFont) (ArtString, error) {
// 	//TODO check rune has to be from ' ' to '~'
// 	return nil
// }

// func artPrint(astr ArtString) {
// }
