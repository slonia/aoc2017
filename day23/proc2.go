package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var commands [][]string
var registers map[string]int
var lastPlayed int
var recovered int
var currentInstruction int

func main() {
	commands = [][]string{}
	registers = make(map[string]int)
	registers["a"] = 1
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		commands = append(commands, strings.Split(line, " "))
	}
	sz := len(commands)
	for currentInstruction < sz {
		command := commands[currentInstruction]
		fmt.Println(registers)
		// fmt.Println(command)
		switch command[0] {
		case "set":
			set(command[1], command[2])
		case "sub":
			sub(command[1], command[2])
		case "mul":
			mul(command[1], command[2])
		case "jnz":
			jnz(command[1], command[2])
		}
	}
	fmt.Println(registers["h"])
}

func set(reg string, source string) {
	// set X Y sets register X to the value of Y.
	registers[reg] = getValue(source)
	currentInstruction++
}

func sub(reg string, source string) {
	// sub X Y decreases register X by the value of Y.
	registers[reg] -= getValue(source)
	currentInstruction++
}

func mul(reg string, source string) {
	// mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
	registers[reg] *= getValue(source)
	currentInstruction++
}

func jnz(source string, jumps string) {
	// jnz X Y jumps with an offset of the value of Y, but only if the value of X is not zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
	if getValue(source) != 0 {
		currentInstruction += getValue(jumps)
	} else {
		currentInstruction++
	}
}

func getValue(reg string) int {
	val, err := strconv.Atoi(reg)
	if err == nil {
		return val
	}
	val, ok := registers[reg]
	if ok {
		return int(val)
	} else {
		registers[reg] = 0
		return 0
	}
}
