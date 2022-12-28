package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"01.kood.tech/git/obudarah/ascii-art-my/asciiart"
)

type input struct {
	output         string
	align          string
	color          string
	lettersToColor string
	banner         string
	text           string
}

// TODO color
// TODO justify
// TODO reverse
func main() {
	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	col, err := cmd.Output()

	fmt.Println(string(col))
	ncol, err := strconv.Atoi((strings.ReplaceAll(string(col), "\n", "")))
	fmt.Println(ncol)

	fmt.Println("\x1b[33m")

	// check that argument only contains ascii symbols from ' ' to '~'
	if ok, runes := asciiart.IsAsciiString(args.text); !ok {
		log.Fatalf("the programme only works with ascii symbols. Next symbols cannot be handle %v ", runes)
	}

	var color string
	var aStrs []asciiart.ArtString
	if args.color != "" {
		color = getColorCode(args.color)
		fmt.Print(color)
	}

	if args.lettersToColor != "" {
		aStrs, err = asciiart.GetStyleArtText(args.text, "banners/"+args.banner+".txt", color, args.lettersToColor)
		if err != nil {
			log.Fatalln(err)
		}
	} else { // no letters to color or no color at all
		aStrs, err = asciiart.GetArtText(args.text, "banners/"+args.banner+".txt")
		if err != nil {
			log.Fatalln(err)
		}
	}

	if args.output == "" {
		asciiart.ArtPrint(aStrs)
	} else {
		file, err:=os.Create(args.output)
		if err != nil {
			log.Fatalln( err)
		}
		defer file.Close()

		asciiart.ArtFprint(file,aStrs)
	}
}

/*
parses arguments. It returns a struct filled with values of flags and parametrs.
If a flag isn't given the function returns for it the default value of the flag.
*/
func parseArgs() (input, error) {
	var args input
	flag.StringVar(&args.output, "output", "", "--output=FILE_NAME\t save output into the flile")
	flag.StringVar(&args.align, "align", "left", "--align=WORD\t align by WORD: left, right, center, justify")
	flag.StringVar(&args.color, "color", "", "--color=WORD\t colored output by  WORD: black, red,green, yellow, blue, purple, cyan, white")
	flag.Parse()

	flags := os.Args[1 : flag.NFlag()+1]
	for _, flag := range flags {
		if !strings.HasPrefix(flag, "--") || !strings.Contains(flag, "=") {
			return args, fmt.Errorf("Usage: go run . [OPTIONS] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> --color=<color> <letters to be colored> something standard\n")
		}
	}

	nArgs := flag.NArg()
	args.banner = asciiart.F_STANDART
	args.lettersToColor = ""

	if nArgs == 1 {
		args.text = flag.Arg(0)
		return args, nil
	}

	/* in this case with color 1 arg=letters,2=text, no banner
	if nArgs == 2 {
		if args.color != "" {
			args.lettersToColor = flag.Arg(0)
			args.text = flag.Arg(1)
			return args, nil
		} else { // the color did not define
			args.text = flag.Arg(0)
			err := args.checkBanner(flag.Arg(1))
			return args, err
		}
	}
	*/

	if nArgs == 2 {
		err := args.checkBanner(flag.Arg(1))
		if err != nil && args.color != "" { // the banner is not valid but there is color, so it wasn't a banner
			args.lettersToColor = flag.Arg(0)
			args.text = flag.Arg(1)
			return args, nil
		}

		args.text = flag.Arg(0)
		return args, err
	}

	if nArgs == 3 && args.color != "" {
		args.lettersToColor = flag.Arg(0)
		args.text = flag.Arg(1)
		err := args.checkBanner(flag.Arg(2))
		return args, err
	}

	return args, fmt.Errorf("Usage: go run . [OPTIONS] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> --color=<color> <letters to be colored> something standard\n")
}

/*
check is the banner valid
*/
func (args *input) checkBanner(banner string) error {
	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
		args.banner = banner
		return nil
	} else {
		return fmt.Errorf("Available banners are: shadow, standard, thinkertoy.\n\nPlease try again.")
	}
}

/*
returns an ANSI escape sequence for color
*/
func getColorCode(color string) string {
	color = strings.ToLower(color)
	switch color {
	case "white", "#ffffff", "rgb(255, 255, 255)":
		return "\033[37m"
	case "cyan", "#00ffff", "rgb(0, 255, 255)":
		return "\033[36m"
	case "purple", "magenta", "#ff00ff", "rgb(255, 0, 255)":
		return "\033[35m"
	case "blue", "#0000ff", "rgb(0, 0, 255)":
		return "\033[34m"
	case "yellow", "#ffff00", "rgb(255, 255, 0)":
		return "\033[33m"
	case "green", "#00ff00", "rgb(0, 255, 0)":
		return "\033[32m"
	case "red", "#ff0000", "rgb(255, 0, 0)":
		return "\033[31m"
	case "":
		return ""
	default:

		if strings.HasPrefix(color, "rgb(") {
			r := regexp.MustCompile(`rgb\((\d{1,3}), (\d{1,3}), (\d{1,3})\)`)
			colors := r.FindStringSubmatch(color)
			if colors == nil {
				return "\033[0m" // Reset color
			}
			res := "\033[38;2"
			for i := 1; i < len(colors); i++ {
				res += ";" + colors[i]
			}
			return res + "m"
		}

		if strings.HasPrefix(color, "#") {
			r := regexp.MustCompile(`#([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})`)
			colors := r.FindStringSubmatch(color)
			if colors == nil {
				return "\033[0m" // Reset color
			}

			res := "\033[38;2"
			for i := 1; i < len(colors); i++ {
				c, err := strconv.ParseInt(colors[i], 16, 32)
				if err != nil {
					return "\033[0m" // Reset color
				}
				res += ";" + strconv.FormatInt(c, 10)
			}
			return res + "m"
		}

		return "\033[0m" // Reset color
	}
}
