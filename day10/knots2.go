package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	const list_size int = 256
	list := make([]int, list_size, list_size)
	chars := make([]int, 0, 20)
	for i, _ := range list {
		list[i] = i
	}
	pos := 0
	skip := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		for _, el := range line {
			char := int(el)
			chars = append(chars, char)
		}
	}
	chars = append(chars, 17, 31, 73, 47, 23)
	for round := 0; round < 64; round++ {
		for _, length := range chars {
			for i := length/2 - 1; i >= 0; i-- {
				s := pos + i
				opp := pos + length - 1 - i
				for s > list_size-1 {
					s -= list_size
				}
				for opp > list_size-1 {
					opp -= list_size
				}
				list[s], list[opp] = list[opp], list[s]
			}
			pos += length + skip
			for pos > list_size-1 {
				pos -= list_size
			}
			skip++
		}
	}
	for block := 0; block < 16; block++ {
		num := list[16*block]
		for i := 1; i < 16; i++ {
			num = num ^ list[16*block+i]
		}
		fmt.Printf("%x", num)
	}
	fmt.Println("")
}
