package main

import (
	"bufio"
	"io"
	"os"
)

func generateCow() string {
	return `  \   ^__^
	\  (oo)\_______
	   (__)\       )\/\
	   ||----w |
	   ||     ||`
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		println("The command is intended to work with pipes.")
		println("Example usage: command | go run main.go")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		lines = append(lines, string(line))
	}

	print(lines[0])
}
