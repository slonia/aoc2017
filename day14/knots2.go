package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

var field [][]int
var labels [][]int
var queue [][]int
var currentLabel int

func main() {
	field = make([][]int, 0, 0)
	labels = make([][]int, 128)
	queue = make([][]int, 0)
	for i := 0; i < 128; i++ {
		field = append(field, knot("jzgqcdpd-", i))
		labels[i] = make([]int, 128)
	}
	currentLabel = 1
	pixelPos := 0
	labels[0][0] = currentLabel
	for pixelPos < 128*128 {
		col := pixelPos % 128
		row := pixelPos / 128
		if field[row][col] == 1 && labels[row][col] == 0 {
			labels[row][col] = currentLabel
			coords := []int{row, col}
			queue = append(queue, coords)
			checkNeighbours()
			currentLabel++
		}
		pixelPos++
	}
	fmt.Println(currentLabel - 1)
}

func checkNeighbours() {
	for len(queue) > 0 {
		// fmt.Println(queue)
		var el []int
		el, queue = queue[len(queue)-1], queue[:len(queue)-1]
		neighbours := getNeighbours(el[0], el[1])
		for _, el := range neighbours {
			x, y := el[0], el[1]
			if field[x][y] == 1 && labels[x][y] == 0 {
				labels[x][y] = currentLabel
				queue = append(queue, []int{x, y})
			}
		}
	}
}

func getNeighbours(x int, y int) [][]int {
	arr := make([][]int, 0)
	if x > 0 {
		arr = append(arr, []int{x - 1, y})
	}
	if x < 127 {
		arr = append(arr, []int{x + 1, y})
	}
	if y > 0 {
		arr = append(arr, []int{x, y - 1})
	}
	if y < 127 {
		arr = append(arr, []int{x, y + 1})
	}
	return arr
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
