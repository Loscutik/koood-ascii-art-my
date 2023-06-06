#!/bin/bash
# Script to run the Ascii-art-fs audit questions

read -rsn1 -p"Press any key to start/continue. To quit click Ctrl+C";echo
go run . "Hello\n" | cat -e

read -rsn1 -p"Press any key to start/continue. To quit click Ctrl+C";echo
go run . "\n\n" | cat -e

read -rsn1 -p"Press any key to start/continue. To quit click Ctrl+C";echo
go run . "Hello\nThere" | cat -e

read -rsn1 -p"Press any key to start/continue. To quit click Ctrl+C";echo
go run . "Hello\n\nThere" | cat -e

read -rsn1 -p"Press any key to start/continue. To quit click Ctrl+C";echo
go run . "\n\n\nHello" | cat -e