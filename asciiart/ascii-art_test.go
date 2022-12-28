package asciiart

import (
	"fmt"
	"testing"

)

func TestIsAsciiString(t *testing.T) {
	str := "fggh"
	res, runes := IsAsciiString(str)
	if !res {
		t.Fatalf("wrong runes %c", runes)
	}
}

func TestAddChar (t *testing.T){
	var aLine ArtString
	artFont, err := GetArtFont("../banners/"+F_STANDART)
	if err != nil {
		t.Fatal(err)
	}
	aLine.AddChar('|',artFont)
	if len(aLine[0])!=4{
		t.Fatalf("1len(aLine)=%d", len(aLine[0]))
	}
	aLine.ArtPrint()
	
	aLine.AddChar('|',artFont)
	if len(aLine[0])!=2*4{
		t.Fatalf("2len(aLine)=%d", len(aLine))
	}
	fmt.Println()
	aLine.ArtPrint()

	aLine2:=StringToArt("||||",artFont)
	if len(aLine2[0])!=4*4{
		t.Fatalf("3len(aLine)=%d", len(aLine2[0]))
	}
	fmt.Println()
	aLine2.ArtPrint()

	aLine2.AddChar('j',artFont)
	if len(aLine2[0])!=5*4+2{
		t.Fatalf("4len(aLine)=%d", len(aLine2[0]))
	}
	fmt.Println()
	aLine2.ArtPrint()

	if len(aLine2)!=1{
		t.Fatalf("end=%d", len(aLine2))
	}
}

func BenchmarkStringToArtEmpty(t *testing.B) {
	str := ""
	artFont, err := GetArtFont("../banners/"+F_STANDART)
	if err != nil {
		t.Fatal(err)
	}
	aStr:=StringToArt(str,artFont)
	fmt.Println(len(aStr))
	fmt.Printf("0:%s:\n",aStr[0])
	// if aStr !=nil{
	// 	t.Fatal("not nil")
	// }
}
