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
	_, err := fmt.Scan(&dimension /*, &seed, &generations*/)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//rand.Seed(int64(seed))

	currentUniverse := make([][]int, dimension)
	for element := range currentUniverse {
		currentUniverse[element] = make([]int, dimension)
	}

	createUniverse(dimension)

	for x := 1; x <= generations; x++ {
		fmt.Printf("Generation #%v\n", x)
		//initialize temporary universe
		nextUniverse := make([][]int, dimension)
		for element := range nextUniverse {
			nextUniverse[element] = make([]int, dimension)
		}
		for i := 0; i < dimension; i++ {
			for j := 0; j < dimension; j++ {
				if currentUniverse[i][j] == 1 {
					count := countNeighbors(currentUniverse, i, j)
					if count > 3 || count < 2 {
						nextUniverse[i][j] = 0
					} else {
						nextUniverse[i][j] = 1
					}
				} else if currentUniverse[i][j] == 0 {
					count := countNeighbors(currentUniverse, i, j)
					if count == 3 {
						nextUniverse[i][j] = 1
					}
				}
			}
			//fmt.Println()
		}
		fmt.Printf("Alive: %v\n", countAlive(nextUniverse))
		currentUniverse = nextUniverse
		printUniverse(currentUniverse)

	}
	//printUniverse(currentUniverse)
}
