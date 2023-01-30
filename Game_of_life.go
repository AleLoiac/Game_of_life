package main

import (
	"fmt"
	"log"
	"math/rand"
)

var dimension int
var seed int
var generations int

func createUniverse(universe [][]int) {
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			x := rand.Intn(2)
			universe[i][j] = x
			if x == 0 {
				//fmt.Print(" ")
			} else {
				//fmt.Print("O")
			}
		}
		//fmt.Println()
	}
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

func main() {

	fmt.Println("Insert dimension and seed of the universe:")
	_, err := fmt.Scan(&dimension, &seed, &generations)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	rand.Seed(int64(seed))

	currentUniverse := make([][]int, dimension)
	for element := range currentUniverse {
		currentUniverse[element] = make([]int, dimension)
	}

	createUniverse(currentUniverse)

	for x := 1; x <= generations; x++ {
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
		currentUniverse = nextUniverse
	}
	printUniverse(currentUniverse)
}
