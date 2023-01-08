package main

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	colors := []string{
		"red",
		"rgb(125, 241, 25)",
		"#12f52a",
		"#2545f2",
		"#c139cc",
		"#ac6fb1",
		"#3dd33rr",
	}

	for i, c := range colors {
		code := getColorCode(c)
		fmt.Println(c)
		fmt.Println(code)
		fmt.Println(i)
		if code == "\033[0m" {
			t.Fatalf("error")
		}
	}
}
