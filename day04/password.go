package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	valid1 := 0
	valid2 := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		valid1 += validate1(line)
		valid2 += validate2(line)
	}
	file.Close()
	fmt.Println(valid1)
	fmt.Println(valid2)
}

func validate1(line string) int {
	words := strings.SplitAfter(line, " ")
	sort.Strings(words)
	length := len(words)
	for i, word := range words {
		if i < length-1 {
			word = strings.Trim(word, " \n")
			word2 := strings.Trim(words[i+1], " \n")
			if word == word2 {
				return 0
			}
		}
	}
	return 1
}

func validate2(line string) int {
	words := strings.SplitAfter(line, " ")
	length := len(words)
	newWords := make([]string, length, length)
	for i, word := range words {
		newWords[i] = orderedWord(word)
	}
	sort.Strings(newWords)
	for i, word := range newWords {
		if i < length-1 {
			if word == newWords[i+1] {
				return 0
			}
		}
	}
	return 1
}

func orderedWord(word string) string {
	chars := strings.Split(strings.Trim(word, " \n"), "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
