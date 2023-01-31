package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"01.kood.tech/git/obudarah/ascii-art-my/asciiart"
)

const (
	USAGE_MESSAGE = "usage: go run . [OPTIONS] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> --color=<color> <letters to be colored> something standard"
)

/*
parses arguments. It returns a struct filled with values of flags and parametrs.
If a flag isn't given the function returns for it the default value of the flag.
*/
func parseArgs() (input, error) {
	var args input
	flag.StringVar(&args.output, "output", "", "--output=FILE_NAME\t save output into the flile")
	flag.StringVar(&args.align, "align", "left", "--align=WORD\t align by WORD: left, right, center, justify")
	flag.StringVar(&args.color, "color", "", "--color=WORD\t color output by  WORD: black, red,green, yellow, blue, purple, cyan, white")
	flag.StringVar(&args.reverse, "reverse", "", "--reverse=path/to/file\t convert the graphic representation by the standard ascii banner into a text")
	flag.Parse()

	flags := os.Args[1 : flag.NFlag()+1]
	for _, flag := range flags {
		if !strings.HasPrefix(flag, "--") || !strings.Contains(flag, "=") {
			return args, fmt.Errorf(USAGE_MESSAGE)
		}
	}

	nArgs := flag.NArg()
	var arg []string
	if nArgs>1 && flag.Arg(0) == "-" {
		arg = flag.Args()[1:]
		nArgs--
	} else {
		arg = flag.Args()
	}

	args.banner = asciiart.F_STANDART
	args.lettersToColor = ""

	if args.reverse != "" {
		if nArgs > 0 {
			return args, fmt.Errorf("you cannot use the reverse option with others. In that case\nUsage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		}

		return args, nil
	}

	if nArgs == 1 {
		args.text = arg[0]
		return args, nil
	}

	/* in this case with color 1 arg=letters,2=text, no banner
	if nArgs == 2 {
		if args.color != "" {
			args.lettersToColor = arg[0)
			args.text = arg[1)
			return args, nil
		} else { // the color did not define
			args.text = arg[0)
			err := args.checkBanner(arg[1))
			return args, err
		}
	}
	*/

	if nArgs == 2 {
		err := args.checkBanner(arg[1])
		if err != nil && args.color != "" { // the banner is not valid but there is color, so it wasn't a banner
			args.lettersToColor = arg[0]
			args.text = arg[1]
			return args, nil
		}

		args.text = arg[0]
		return args, err
	}

	if nArgs == 3 && args.color != "" {
		args.lettersToColor = arg[0]
		args.text = arg[1]
		err := args.checkBanner(arg[2])
		return args, err
	}

	return args, fmt.Errorf(USAGE_MESSAGE)
}

/*
check is the banner valid
*/
func (args *input) checkBanner(banner string) error {
	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
		args.banner = banner
		return nil
	} else {
		return fmt.Errorf("available banners are: shadow, standard, thinkertoy.\n\nPlease try again")
	}
}
