package fonts

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func errHandle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func ArtPrint(aStr ArtString) {
	// the empty string must comprise only 1 line
	if aStr[0] == "" {
		fmt.Println()
		return
	}
}

func BenchmarkReadArtLetter(b *testing.B) {
	STEP := 24
	f, err := os.Open("fonts/standard.txt")
	errHandle(err)
	var aLetter ArtString
	scanner := bufio.NewScanner(f)

	for i := 0; i <= 9*STEP; i++ {
		scanner.Scan()
	}
	aLetter, err = readArtChar(scanner)
	fmt.Println()
	ArtPrint(aLetter)
	if err != nil {
		b.Fatal(err)
	}
}
