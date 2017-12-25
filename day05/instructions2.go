package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	numbers := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	ind := 0
	steps := 0
	for ind < len(numbers) {
		new_ind := ind + numbers[ind]
		if numbers[ind] > 2 {
			numbers[ind]--
		} else {
			numbers[ind]++
		}
		ind = new_ind
		steps++
	}
	fmt.Printf("%v\n", steps)
}

func print(numbers []int, ind int) {
	fmt.Printf("%v - ", ind)
	for i, el := range numbers {
		if i == ind {
			fmt.Printf("[%v] ", el)
		} else {
			fmt.Printf("%v ", el)
		}
	}
	fmt.Printf("\n")

}
