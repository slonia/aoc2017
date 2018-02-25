package main

import "fmt"

type Generator struct {
	value      int64
	factor     int64
	multipleOf int64
}

const div int64 = 2147483647

func (g *Generator) next() int64 {
	newValue := g.value * g.factor % div
	for newValue%g.multipleOf != 0 {
		newValue = newValue * g.factor % div
	}
	g.value = newValue
	return newValue
}

func main() {
	genA := Generator{783, 16807, 4}
	genB := Generator{325, 48271, 8}
	matchedPairs := 0
	for i := 0; i < 5000000; i++ {
		bitsA := fmt.Sprintf("%b", genA.next())
		bitsB := fmt.Sprintf("%b", genB.next())
		indA := len(bitsA) - 16
		if indA < 0 {
			indA = 0
		}
		indB := len(bitsB) - 16
		if indB < 0 {
			indB = 0
		}
		bitsA = bitsA[indA:]
		bitsB = bitsB[indB:]
		if bitsA == bitsB {
			matchedPairs++
		}
	}
	fmt.Println(matchedPairs)
}
