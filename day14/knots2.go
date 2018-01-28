package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

var field [][]int

func main() {
	field = make([][]int, 0, 0)
	for i := 0; i < 128; i++ {
		field = append(field, knot("flqrgnkx-", i))
	}
	fmt.Println(field[0][0])
}

func knot(str string, i int) []int {
	a := strconv.Itoa(i)
	str += a
	const list_size int = 256
	list := make([]int, list_size, list_size)
	for i, _ := range list {
		list[i] = i
	}
	var pos, skip int
	charsStr := strings.Split(str, "")
	chars := make([]int, 0, 0)
	for _, el := range charsStr {
		chars = append(chars, int(el[0]))
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
	arr := make([]int, 0, 0)
	for block := 0; block < 16; block++ {
		num := list[16*block]
		for i := 1; i < 16; i++ {
			num = num ^ list[16*block+i]
		}
		leadingZeros := bits.LeadingZeros8(uint8(num))
		str := fmt.Sprintf("%b", num)
		for i := 0; i < leadingZeros; i++ {
			arr = append(arr, 0)
		}
		if leadingZeros != 8 {
			for _, el := range strings.Split(str, "") {
				if el == "1" {
					arr = append(arr, 1)
				} else {
					arr = append(arr, 0)
				}
			}
		}
	}
	return arr
}
