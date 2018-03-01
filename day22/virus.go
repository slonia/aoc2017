package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var infected [][]int
var posX, posY, bursts int
var direction string = "up"

const steps int = 10000

func main() {
	infected = make([][]int, 0, 0)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lineNum int
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		if lineNum == 0 {
			posX = len(split) / 2
		}
		for i, el := range split {
			if el == "#" {
				infected = append(infected, []int{lineNum, i})
			}
		}
		lineNum++
	}
	posY = lineNum / 2
	for i := 0; i < steps; i++ {
		ind := isInfected(posY, posX)
		if ind == -1 {
			infected = append(infected, []int{posY, posX})
			direction = turnLeft()
			bursts++
		} else {
			infected = append(infected[:ind], infected[ind+1:]...)
			direction = turnRight()
		}
		switch direction {
		case "up":
			posY--
		case "right":
			posX++
		case "down":
			posY++
		case "left":
			posX--
		}
	}
	fmt.Println(bursts)
}

func isInfected(y, x int) int {
	for i, el := range infected {
		if el[0] == y && el[1] == x {
			return i
		}
	}
	return -1
}

func turnRight() string {
	switch direction {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	default:
		return "up"
	}
}

func turnLeft() string {
	switch direction {
	case "up":
		return "left"
	case "left":
		return "down"
	case "down":
		return "right"
	default:
		return "up"
	}
}
