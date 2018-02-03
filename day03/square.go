package main

import (
	"fmt"
	"math"
)

func main() {
	i := 361527
	fmt.Printf("%v - %v\n", i, stepsTo(i))
}

func stepsTo(n int) int {
	squareSize := int(math.Ceil(math.Sqrt(float64(n))))
	if squareSize%2 == 0 {
		squareSize++
	}
	maxN := squareSize * squareSize
	minRow := 0
	minCol := 0
	maxRow := squareSize - 1
	maxCol := squareSize - 1
	curRow := squareSize - 1
	curCol := squareSize - 1
	direction := "left"
	for maxN != n {
		switch direction {
		case "left":
			if curCol > minCol {
				curCol--
				maxN--
			} else {
				direction = "up"
				minCol++
			}
		case "up":
			if curRow > minRow {
				curRow--
				maxN--
			} else {
				direction = "right"
				minRow++
			}
		case "right":
			if curCol < maxCol {
				curCol++
				maxN--
			} else {
				direction = "down"
				maxCol--
			}
		case "down":
			if curRow < maxRow {
				curRow++
				maxN--
			} else {
				direction = "left"
				maxRow--
			}
		}
	}
	return int(math.Abs(float64(curRow-squareSize/2))) + int(math.Abs(float64(curCol-squareSize/2)))
}
