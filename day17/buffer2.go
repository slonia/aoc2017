package main

import "fmt"

func main() {
	stepSize := 376
	steps := 50000000
	pos := 0
	answer := -1
	for i := 1; i <= steps; i++ {
		pos += stepSize + 1
		for pos >= i {
			pos -= i
		}
		if pos == 0 {
			answer = i
		}
	}
	fmt.Println(answer)
}
