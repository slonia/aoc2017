package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text string
	for scanner.Scan() {
		line := scanner.Text()
		text += line
	}
	text = filterText(text)
	score := countScore(text)
	fmt.Println(score)
}

func countScore(text string) int {
	score := 0
	curScore := 0
	for _, c := range text {
		if c == '{' {
			curScore++
		} else if c == '}' {
			score += curScore
			curScore--
		}
	}
	return score
}

func filterText(input string) string {
	var output string
	garbage := false
	for i := 0; i < len(input); {
		c := input[i]
		if garbage && c == '!' {
			i++
		} else if c == '<' {
			garbage = true
		} else if c == '>' {
			garbage = false
		} else if !garbage { //&& (c == '{' || c == '}') {
			output += string(c)
		}
		i++
	}
	fmt.Println(output)
	return output
}
