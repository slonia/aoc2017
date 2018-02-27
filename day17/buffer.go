package main

import "fmt"

func main() {
	buffer := []int{0}
	stepSize := 376
	steps := 2017
	pos := 0
	for i := 1; i <= steps; i++ {
		pos += stepSize + 1
		ln := len(buffer)
		for pos >= ln {
			pos -= ln
		}
		// fmt.Println(pos)
		buffer = append(buffer, 0)
		copy(buffer[pos+1:], buffer[pos:])
		buffer[pos+1] = i
	}
	fmt.Println(buffer[pos+2])
}
