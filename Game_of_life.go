package main

import (
	"fmt"
	"log"
	"math/rand"
)

var dimension int
var generations int = 10

func createUniverse(dimension int) [][]int {
	universe := make([][]int, dimension)
	for i := range universe {
		universe[i] = make([]int, dimension)
		for j := range universe[i] {
			universe[i][j] = rand.Intn(2) // Random 0 or 1
		}
	}
	return universe
}

func countNeighbors(universe [][]int, i int, j int) int {
	count := 0
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x == i && y == j {
				continue
			}
			count += universe[outOfBorder(x)][outOfBorder(y)]
		}
	}
	return count
}

func outOfBorder(x int) int {
	if x == -1 {
		x = dimension - 1
	} else if x == dimension {
		x = 0
	}
	return x
}

func printUniverse(universe [][]int) {
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			x := universe[i][j]
			if x == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func countAlive(universe [][]int) int {
	count := 0
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if universe[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func nextGeneration(current [][]int) [][]int {
	next := make([][]int, dimension)
	for i := range next {
		next[i] = make([]int, dimension)
		for j := range next[i] {
			neighbors := countNeighbors(current, i, j)
			if current[i][j] == 1 {
				next[i][j] = 1
				if neighbors < 2 || neighbors > 3 {
					next[i][j] = 0 // Underpopulation or overpopulation
				}
			} else if neighbors == 3 {
				next[i][j] = 1 // Reproduction
			}
		}
	}
	return next
}

func main() {
	fmt.Println("Insert dimension of the universe:")
	_, err := fmt.Scan(&dimension)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	currentUniverse := createUniverse(dimension)

	for gen := 1; gen <= generations; gen++ {
		fmt.Printf("Generation #%d\n", gen)
		fmt.Printf("Alive: %d\n", countAlive(currentUniverse))
		printUniverse(currentUniverse)
		currentUniverse = nextGeneration(currentUniverse)
	}
}
