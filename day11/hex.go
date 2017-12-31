package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	x, y, z := 0, 0, 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var distance, max_distance float64
	max_distance = -1
	for scanner.Scan() {
		line := scanner.Text()
		dirs := strings.Split(line, ",")
		for _, dir := range dirs {
			switch dir {
			case "n":
				x++
				z--
			case "ne":
				y--
				x++
			case "se":
				y--
				z++
			case "s":
				x--
				z++
			case "sw":
				y++
				x--
			case "nw":
				y++
				z--
			}
			distance = (math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))) / 2
			if distance > max_distance {
				max_distance = distance
			}
		}
	}
	fmt.Println(distance)
	fmt.Println(max_distance)
}
