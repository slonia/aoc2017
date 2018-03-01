package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var infected [][]int
var weakend [][]int
var flagged [][]int
var posX, posY, bursts int
var direction string = "up"

const steps int = 10000000

func main() {
	infected = make([][]int, 0, 0)
	weakend = make([][]int, 0, 0)
	flagged = make([][]int, 0, 0)
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
		if i%500000 == 0 {
			fmt.Println("Step ", i)
		}
		if infInd := searchInArray(infected, posY, posX); infInd != -1 {
			infected = append(infected[:infInd], infected[infInd+1:]...)
			flagged = append(flagged, []int{posY, posX})
			direction = turnRight()
		} else if weakInd := searchInArray(weakend, posY, posX); weakInd != -1 {
			weakend = append(weakend[:weakInd], weakend[weakInd+1:]...)
			infected = append(infected, []int{posY, posX})
			bursts++
		} else if flagInd := searchInArray(flagged, posY, posX); flagInd != -1 {
			flagged = append(flagged[:flagInd], flagged[flagInd+1:]...)
			direction = turnRight()
			direction = turnRight()
		} else {
			weakend = append(weakend, []int{posY, posX})
			direction = turnLeft()
		}
		move()
	}
	fmt.Println(bursts)
}

func move() {
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

func searchInArray(array [][]int, y int, x int) int {
	for i, el := range array {
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
