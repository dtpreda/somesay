package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		println("The command is intended to work with pipes.")
		println("Example usage: command | go run main.go")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		char, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}

		output = append(output, char)
	}

	print(string(output))
}
