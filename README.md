# ASCII-art

Ascii-art is a program which receives a string as an argument and outputs the string in a graphic representation using ASCII. 
Given string must only contain ascii printable characters. If the argument contains spaces or/and shell's special characters, it must be quated. 
Note that the `!` in double quotes must be escaped on Mac and must not escaped on Linux. It means on Mac `"\!"` will be printed  as `!` and on Linux `"\!"` will be printed  as `\!`. Single quates work the same on Mac and on Linux so `'\!'` will be pinted as `\!`. On both system in double quated string you must escape  \` ( the backtick) , `"` and `\` .
This project decode a escape sequences `\n` in given string as a new line.