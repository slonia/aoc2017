package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var maxInBlock int = -1
var maxBlock int = -1
var numbers []int = []int{}
var keys []string = []string{}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
		if number > maxBlock {
			maxBlock = number
			maxInBlock = i
		}
		i++
	}
	set := map[string]bool{}
	key := arrayToStr()
	keys = append(keys, key)
	set[key] = true
	steps := 1
	for {
		rearrange()
		key = arrayToStr()
		// fmt.Println(key)
		_, ok := set[key]
		if ok {
			break
		} else {
			keys = append(keys, key)
			set[key] = true
			steps++
		}
	}
	// answer for part#1
	fmt.Println(steps)
	for i, k := range keys {
		if k == key {
			// answer for part#2
			fmt.Println(len(keys) - i)
			break
		}
		i++
	}
}
func rearrange() {
	numbers[maxInBlock] = 0
	ind := incInd(maxInBlock)
	for maxBlock > 0 {
		numbers[ind]++
		ind = incInd(ind)
		maxBlock--
	}
	findMax()
}
func findMax() {
	maxInBlock = -1
	maxBlock = -1
	for i, el := range numbers {
		if el > maxBlock {
			maxBlock = el
			maxInBlock = i
		}
	}
}
func incInd(ind int) int {
	l := len(numbers)
	ind++
	if ind > l-1 {
		ind = 0
	}
	return ind
}

func arrayToStr() string {
	str := ""
	ln := len(numbers)
	for i, el := range numbers {
		str += strconv.Itoa(el)
		if i < ln-1 {
			str += ","
		}
	}
	return str
}
