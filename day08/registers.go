package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var register map[string]int

func main() {
	register = make(map[string]int)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var condition bool
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		target, op, value_s, cond_reg, cond_op, cond_val_s := s[0], s[1], s[2], s[4], s[5], s[6]
		check_register(target)
		check_register(cond_reg)
		cond_val, _ := strconv.Atoi(cond_val_s)
		value, _ := strconv.Atoi(value_s)
		condition = check_cond(cond_reg, cond_op, cond_val)
		if condition {
			if op == "inc" {
				register[target] += value
			} else {
				register[target] -= value
			}
		}
	}
	max := -1
	for _, el := range register {
		if el > max {
			max = el
		}
	}
	fmt.Println(max)
}

func check_register(reg string) {
	_, ok := register[reg]
	if !ok {
		register[reg] = 0
	}
}

func check_cond(reg string, op string, val int) bool {
	switch op {
	case "<":
		return register[reg] < val
	case ">":
		return register[reg] > val
	case ">=":
		return register[reg] >= val
	case "<=":
		return register[reg] <= val
	case "==":
		return register[reg] == val
	case "!=":
		return register[reg] != val
	}
	return false
}
