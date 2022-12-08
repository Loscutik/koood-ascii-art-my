package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	F_STANDART    = "../fonts/shadow.txt"
	SYMBOL_HEIGHT = 8
	FIRST=' '
	LAST = 'Z'
)

type (
	ArtLetter = [SYMBOL_HEIGHT][]byte //when use []byte - map keeps wrong bytes, when use string - ok
	ArtFont   = map[rune]ArtLetter
	ArtString = [SYMBOL_HEIGHT]string
)

func main() {
	/*file, _ := os.Open("testfont.txt")
	aFont := make(map[rune]ArtLetter)
	scanner := bufio.NewScanner(file)
	var aLetter ArtLetter

	scanner.Scan()
	for line := 0; line < SYMBOL_HEIGHT; line++ {
		if scanner.Scan() {
			aLetter[line] = (scanner.Bytes())
		}
	}

	aFont[' '] = aLetter
	fmt.Println(aFont[' '])

	scanner.Scan()
	for line := 0; line < SYMBOL_HEIGHT; line++ {
		if scanner.Scan() {
			aLetter[line] = (scanner.Bytes())
		}
	}

	aFont['!'] = aLetter
	fmt.Println(aFont['!'])

	scanner.Scan()
	for line := 0; line < SYMBOL_HEIGHT; line++ {
		if scanner.Scan() {
			aLetter[line] = (scanner.Bytes())
		}
	}

	aFont['"'] = aLetter
	fmt.Println(aFont['"'])
	*/
	 aFont,_:=GetArtFont(F_STANDART)

	 for char := (FIRST); char <= LAST; char++ {
	for i := 0; i < SYMBOL_HEIGHT; i++ {
		fmt.Print(char)
		fmt.Print(": ")
		fmt.Println(string(aFont[char][i]))
	}
}
}

func GetArtFont(fileName string) (aFont ArtFont, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	aFont = make(map[rune]ArtLetter)
	scanner := bufio.NewScanner(file)

	for char := (FIRST); char <= LAST; char++ {
		scanner.Scan()
		aFont[char], err = readArtLetter(scanner)
		//fmt.Println(aFont[char])
		// for i := 0; i < SYMBOL_HEIGHT; i++ {
		// fmt.Println(string(aFont[char][i]))
		// }
		if err != nil {
			err = fmt.Errorf("error reading the character from a file: %s; character: %c; file:  %s", err, char, fileName)
			break
		}

		// scanner.Bytes()
	}

	return
}

func readArtLetter(scanner *bufio.Scanner) (aLetter ArtLetter, err error) {
	for line := 0; line < SYMBOL_HEIGHT; line++ {
		if scanner.Scan() {
			aLetter[line] = (scanner.Bytes())
		} else if err = scanner.Err(); err != nil {
			break
		} else {
			err = fmt.Errorf("unexpected break of scanning")
			break
		}
	}
	return aLetter, err
}

// func stringToArt(str []byte, afont ArtFont) ArtString {
// 	return nil
// }

// func artPrint(astr ArtString) {
// }
