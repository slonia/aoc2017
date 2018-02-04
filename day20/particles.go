package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const bigNumber = 10000000000.0

type Coords struct {
	x, y, z int
}

type Particle struct {
	point    Coords
	speed    Coords
	velocity Coords
}

var particles []Particle

func main() {
	particles = make([]Particle, 0)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		extractParticle(line)
	}
	findMin()
}

func findMin() {
	minInd := -1
	var minDist float64
	for i, particle := range particles {
		point := math.Sqrt(float64(particle.point.x*particle.point.x + particle.point.y*particle.point.y + particle.point.z*particle.point.z))
		speed := math.Sqrt(float64(particle.point.x*particle.point.x + particle.point.y*particle.point.y + particle.point.z*particle.point.z))
		velocity := math.Sqrt(float64(particle.point.x*particle.point.x + particle.point.y*particle.point.y + particle.point.z*particle.point.z))
		dist := math.Abs(point + speed*bigNumber + velocity*bigNumber*bigNumber/2)
		fmt.Printf("%v, %f\n", i, dist)
		if minInd == -1 || dist < minDist {
			minInd = i
			minDist = dist
		}
	}
	fmt.Println(minInd)
}

func extractParticle(line string) {
	// p=<-2978,1564,-1301>, v=<-427,223,-186>, a=<31,-15,18>
	strs := strings.Split(line, ", ")
	point := extractCoords(strs[0])
	speed := extractCoords(strs[1])
	velocity := extractCoords(strs[2])
	particles = append(particles, Particle{point, speed, velocity})
}

func extractCoords(line string) Coords {
	line = line[3 : len(line)-1]
	coords := strings.Split(line, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])
	return Coords{x, y, z}
}
