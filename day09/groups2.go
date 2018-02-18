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
	score := filterText(text)
	fmt.Println(score)
}

func filterText(input string) int {
	garbage := false
	garbageCount := 0
	for i := 0; i < len(input); {
		c := input[i]
		if garbage && c == '!' {
			i++
		} else if c == '<' && !garbage {
			garbage = true
		} else if c == '>' {
			garbage = false
		} else if garbage {
			fmt.Printf("%c", c)
			garbageCount++
		}
		i++
	}
	fmt.Println()
	return garbageCount
}
