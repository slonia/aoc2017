package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const sz int = 16

var programs []rune

func main() {
	commands := []string{}
	for i := 0; i < sz; i++ {
		programs = append(programs, rune(i+97))
	}
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		commands = append(commands, strings.Split(line, ",")...)
	}
	// programs are repeated
	// formula is 24+x%24
	// 1000000000%24 = 16
	// 24 + 16 = 40
	for i := 0; i < 40; i++ {
		for _, command := range commands {
			command, args := command[0], command[1:]
			switch string(command) {
			case "s":
				spin(args)
			case "x":
				exchange(args)
			case "p":
				partner(args)
			}
		}
	}
	for i := 0; i < sz; i++ {
		fmt.Printf("%c", programs[i])
	}
	fmt.Println()
}

func spin(command string) {
	num, _ := strconv.Atoi(command)
	programs = append(programs[(sz-num):], programs[:(sz-num)]...)
}

func exchange(command string) {
	strs := strings.Split(command, "/")
	a, _ := strconv.Atoi(strs[0])
	b, _ := strconv.Atoi(strs[1])
	programs[a], programs[b] = programs[b], programs[a]
}

func partner(command string) {
	strs := strings.Split(command, "/")
	a := indOf(strs[0])
	b := indOf(strs[1])
	programs[a], programs[b] = programs[b], programs[a]
}
func indOf(el string) int {
	for i, str := range programs {
		if str == rune(el[0]) {
			return i
		}
	}
	return -1
}
