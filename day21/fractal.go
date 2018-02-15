package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Rule struct {
	from, to []string
}

var field []string = []string{".#.", "..#", "###"}
var rules []Rule

const iterations int = 1

func main() {
	rules = make([]Rule, 0)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " => ")
		from := strings.Split(split[0], "/")
		to := strings.Split(split[1], "/")
		rules = append(rules, Rule{from, to})
	}
	for i := 0; i < iterations; i++ {
		expand()
	}
}

func expand() {
	var step int
	curLen := len(field)
	if curLen%2 == 0 {
		step = 2
	} else {
		step = 3
	}
	var newField []string
	newLen := curLen / step * (step + 1)
	for i := 0; i < newLen; i++ {
		str := make([]string, newLen)
		newStr := strings.Join(str, "")
		newField = append(newField, newStr)
	}
	for i := 0; i < curLen/step; i++ {
		square := make([]string, 0)
		for j := 0; j < step; j++ {
			square = append(square, field[j][(step*i):(step*(i+1))])
		}
		rule, err := findRule(square)
		if err == nil {
			for _, newRow := range rule.to {

			}
		}
	}
}

func findRule(square []string) (Rule, error) {
	for _, permutaion := range permute(square) {
		for _, rule := range rules {
			match := true
			for i, rows := range permutaion {
				if rows != rule.from[i] {
					match = false
					break
				}
			}
			if match {
				return rule, nil
			}
		}
	}
	return Rule{square, square}, errors.New("Cannot match")
}

func permute(square []string) [][]string {
	permutations := make([][]string, 0)
	ln := len(square)
	for rot := 0; rot < 4; rot++ {
		newSquare := make([]string, ln)
		for i := 0; i < ln; i++ {
			str := make([]string, 0)
			for j := ln - 1; j > -1; j-- {
				str = append(str, string(square[j][i]))
			}
			newSquare[i] = strings.Join(str, "")
		}
		square = newSquare
		permutations = append(permutations, square)
		horFlip := make([]string, ln)
		verFlip := make([]string, ln)
		for i := 0; i < ln; i++ {
			horFlip[i] = reverse(square[i])
			verFlip[ln-1-i] = square[i]
		}
		permutations = append(permutations, horFlip)
		permutations = append(permutations, verFlip)
	}
	return permutations
}

func reverse(str string) string {
	newStr := ""
	for i := len(str) - 1; i > -1; i-- {
		newStr = newStr + string(str[i])
	}
	return newStr
}
