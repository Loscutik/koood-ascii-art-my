package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"01.kood.tech/git/obudarah/ascii-art-my/asciiart"
)

const (
	ARGS       = 1
	F_STANDART = "standard.txt"
)

type input struct {
	output string
	align  string
	banner string
	text   string
}
//TODO color
//TODO justify
//TODO reverse 
func main() {
	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	// var c exec.Cmd
	// fmt.Println(c.Env[4])

	fmt.Println("\x1b[33m")

	// check that argument only contains ascii symbols from ' ' to '~'
	if ok, runes := asciiart.IsAsciiString(args.text); !ok {
		log.Fatalf("the programme only works with ascii symbols. Next symbols cannot be handle %v ", runes)
	}

	aStr, err := asciiart.TextToArt(args.text, "banners/"+args.banner+".txt")
	if err != nil {
		log.Fatalln(err)
	}

	if args.output == "" {
		fmt.Print(aStr)
	} else {
		err = os.WriteFile(args.output, []byte(aStr), 0o644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func parseArgs() (input, error) {
	var args input
	flag.StringVar(&args.output, "output", "", "--output=fileName.txt")
	flag.StringVar(&args.align, "align", "left", "--align=center")
	flag.Parse()

	nArgs := flag.NArg()
	if nArgs == 1 {
		args.text = flag.Arg(0)
		args.banner = "standard"
		return args, nil
	}

	if nArgs == 2 {
		args.text = flag.Arg(0)
		if flag.Arg(1) == "standard" || flag.Arg(1) == "shadow" || flag.Arg(1) == "thinkertoy" {
			args.banner = flag.Arg(1)
		} else {
			return args, fmt.Errorf("Available banners are: shadow, standard, thinkertoy.\n\nPlease try again.")
		}

		return args, nil
	}

	return args, fmt.Errorf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
}
