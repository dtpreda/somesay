package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"strings"
)

func generateCow() string {
	return `       \  ^__^
	\ (oo)\_______
	  (__)\       )\/\
	      ||----w |
	      ||     ||
  `
}

func tabToWhitespace(lines []string) []string {
	var result []string

	for _, line := range lines {
		result = append(result, strings.Replace(line, "\t", " ", -1))
	}

	return result
}

func maxLineWidth(lines []string) int {
	var max int

	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}

	return max
}

func padLines(lines []string, width int) []string {
	var result []string

	for _, line := range lines {
		result = append(result, line+strings.Repeat(" ", width-len(line)))
	}

	return result
}

func buildBalloon(lines []string, width int) string {
	var result string

	result += "o" + strings.Repeat("-", width+2) + "o\n"
	for _, line := range lines {
		result += "| " + line + " |"
		result += "\n"
	}
	result += "o" + strings.Repeat("-", width+2) + "o\n"
	return result
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

	cow := generateCow()

	lines = tabToWhitespace(lines)
	width := math.Max(float64(maxLineWidth(lines)), float64(maxLineWidth(strings.Split(cow, "\n"))))
	lines = padLines(lines, int(width))
	ballon := buildBalloon(lines, maxLineWidth(lines))
	print(ballon)
	print(generateCow())
}
