package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const list_size int = 255
	list := make([]int, list_size, list_size)
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
		s := strings.Split(line, ",")
		for _, el := range s {
			length, _ := strconv.Atoi(el)
			for i := length/2 - 1; i >= 0; i-- {
				s := pos + i
				opp := pos + length - 1 - i
				if s > list_size-1 {
					s -= list_size
				}
				if opp > list_size-1 {
					opp -= list_size
				}
				list[s], list[opp] = list[opp], list[s]
			}
			pos += length + skip
			if pos > list_size-1 {
				pos -= list_size
			}
			skip++
		}
	}
	fmt.Println(list[0] * list[1])
}
