package main

import "fmt"

type Generator struct {
	value  int64
	factor int64
}

const div int64 = 2147483647

func (g *Generator) next() int64 {
	g.value = g.value * g.factor % div
	return g.value
}

func main() {
	genA := Generator{783, 16807}
	genB := Generator{325, 48271}
	matchedPairs := 0
	for i := 0; i < 40000000; i++ {
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
