package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var field [][]string
var szX, szY, curX, curY, prevX, prevY int
var path []string = make([]string, 0)
var direction string

func main() {
	field = [][]string{}
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		field = append(field, chars)
		szY++
		if len(chars) > szX {
			szX = len(chars)
		}
	}
	curY, curX = findStartingPoint()
	prevX, prevY = -1, -1
	for prevX != curX || prevY != curY {
		direction = getDirection()
		prevX, prevY = curX, curY
		switch direction {
		case "up":
			moveUp()
		case "down":
			moveDown()
		case "right":
			moveRight()
		case "left":
			moveLeft()
		}
		// printField()
		// time.Sleep(400 * time.Millisecond)
	}
	fmt.Println(path)
}

func printField() {
	clear()
	for i := 0; i < szY; i++ {
		for j, el := range field[i] {
			if curY == i && curX == j {
				fmt.Printf("*")
			} else {
				fmt.Printf(el)
			}
		}
		fmt.Println("")
	}
}

func moveUp() {
	if curY > 0 {
		curY--
	}
}

func moveDown() {
	if curY < szY-1 {
		curY++
	}
}

func moveRight() {
	if curX < szX-1 {
		curX++
	}
}

func moveLeft() {
	if curX > 0 {
		curX--
	}
}

func findStartingPoint() (int, int) {
	direction = "down"
	for i, el := range field[0] {
		if el != " " {
			return 0, i
		}
	}
	return 0, 0
}

func getDirection() string {
	char := field[curY][curX]
	if char == "|" {
		return direction
	} else if char == "-" {
		return direction
	} else if char >= "A" && char <= "Z" {
		path = append(path, char)
		return direction
	} else if char == "+" {
		if direction == "down" || direction == "up" {
			curRowSize := len(field[curY])
			if curX > 0 && (field[curY][curX-1] == "-" || (field[curY][curX-1] >= "A" && field[curY][curX-1] <= "Z")) {
				return "left"
			} else if curX < curRowSize-1 && (field[curY][curX+1] == "-" || (field[curY][curX+1] >= "A" && field[curY][curX+1] <= "Z")) {
				return "right"
			} else {
				return "stop"
			}
		} else {
			if curY > 0 && len(field[curY-1]) > curX && (field[curY-1][curX] == "|" || (field[curY-1][curX] >= "A" && field[curY-1][curX] <= "Z")) {
				return "up"
			} else if curY < szY-1 && len(field[curY+1]) > curX && (field[curY+1][curX] == "|" || (field[curY+1][curX] >= "A" && field[curY+1][curX] <= "Z")) {
				return "down"
			} else {
				return "stop"
			}
		}
	} else {
		return "stop"
	}
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
