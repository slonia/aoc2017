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
		// fmt.Println(registers)
		// fmt.Println(command)
		switch command[0] {
		case "snd":
			snd(command[1])
		case "set":
			set(command[1], command[2])
		case "add":
			add(command[1], command[2])
		case "mul":
			mul(command[1], command[2])
		case "mod":
			mod(command[1], command[2])
		case "rcv":
			rcv(command[1])
		case "jgz":
			jgz(command[1], command[2])
		}
	}
}

func snd(reg string) {
	// snd X plays a sound with a frequency equal to the value of X.
	lastPlayed = getValue(reg)
	currentInstruction++
}

func set(reg string, source string) {
	// set X Y sets register X to the value of Y.
	registers[reg] = getValue(source)
	currentInstruction++
}

func add(reg string, source string) {
	// add X Y increases register X by the value of Y.
	registers[reg] += getValue(source)
	currentInstruction++
}

func mul(reg string, source string) {
	// mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
	registers[reg] *= getValue(source)
	currentInstruction++
}

func mod(reg string, source string) {
	// mod X Y sets register X to the remainder of dividing the value contained in register X by the value of Y (that is, it sets X to the result of X modulo Y).
	registers[reg] %= getValue(source)
	currentInstruction++
}

func rcv(source string) {
	// rcv X recovers the frequency of the last sound played, but only when the value of X is not zero. (If it is zero, the command does nothing.)
	if getValue(source) != 0 {
		recovered = lastPlayed
		registers[source] = recovered
		fmt.Println(recovered)
		currentInstruction += 10000 // for exit
	}
	currentInstruction++
}

func jgz(source string, jumps string) {
	// jgz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, a
	if getValue(source) > 0 {
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
