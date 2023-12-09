package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/i-DarkLight/memari/binary"
	"github.com/i-DarkLight/memari/hex"
)

func main() {
	println("Enter your assembly code (to finish type '~' into the terminal on a new line and press enter):")
	input := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		input.Scan()
		line := input.Text()
		if line == "~" {
			break
		}
		lines = append(lines, line)
	}
	fmt.Printf("\033[1A\033[K")
	println("===============================")
	println("Here is your code converted to Hex form:")
	hex.Run(lines)
	println("And here is your code in Binary form:")
	binary.ToBinary()
}
