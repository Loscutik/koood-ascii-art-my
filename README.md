There are two methods of how you can color letters:

Using the terminal, please go to the destination folder, and type the following to the command line:
go run . --color=red Yrx "Your text here" standard

go run . --color=red "Yo*r t*xt *ere" "Your text here" standard


"go run ." will launch the program.

"--color=red" is a flag for color. Many terminals support only 8 colors including the black background color - white, yellow, red, blue, cyan, purple (magenta), and green. So, the program can accept only these colors, written as words, and their RGB and HEX color codes.

"Yrx" - an example of the letters to be colored. Just choose some letters from the text to color them.
Another option: You can write the same string as your color request, but replace the letters, that you want to color, with stars. Then, the letters relaced with stars will be colored.

"Your text here" can be replaced with a random text. If you want to write more than one word, please put the whole phrase in quotes.

"standard" is a font choice. There are 3 possible options - "standard", "shadow" and "thinkertoy". If you leave it empty, the program automatically chooses the standard font.
