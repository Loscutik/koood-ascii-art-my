package main

import (
	"regexp"
	"strconv"
	"strings"
)

/*
returns an ANSI escape sequence for color
*/
func getColorCode(color string) string {
	color = strings.ToLower(color)
	switch color {
	case "white":
		return "\033[37m"
	case "cyan":
		return "\033[36m"
	case "purple", "magenta":
		return "\033[35m"
	case "blue":
		return "\033[34m"
	case "yellow":
		return "\033[33m"
	case "green":
		return "\033[32m"
	case "red":
		return "\033[31m"
	default:
		if strings.HasPrefix(color, "rgb(") {
			return parseColorRgb(color)
		}

		if strings.HasPrefix(color, "#") {
			return parseColorHex(color)
		}

		return "\033[0m" // Reset color
	}
}

/*
parses given rgb color and returns an ANSI escape sequence for color
*/
func parseColorRgb(rgb string) string {
	r := regexp.MustCompile(`rgb\((\d{1,3}), ?(\d{1,3}), ?(\d{1,3})\)`)
	colors := r.FindStringSubmatch(rgb)
	if colors == nil {
		return "\033[0m" // Reset color
	}

	res := "\033[38;2"
	for i := 1; i < len(colors); i++ {
		res += ";" + colors[i]
	}
	return res + "m"
}

/*
parses given Hex color and returns an ANSI escape sequence for color
*/
func parseColorHex(hex string) string {
	r := regexp.MustCompile(`#([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})`)
	colors := r.FindStringSubmatch(hex)
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
