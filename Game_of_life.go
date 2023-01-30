package main

import (
	"fmt"
	"math/rand"
)

var dimension int
var seed int
var generations int

func outOfBorder(x int) int {
	if x == -1 {
		x = dimension - 1
	} else if x == dimension {
		x = 0
	}
	return x
}

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
	fmt.Scan(&dimension, &seed, &generations)
	rand.Seed(int64(seed))

	currentUniverse := make([][]int, dimension)
	for element := range currentUniverse {
		currentUniverse[element] = make([]int, dimension)
	}

	//nextUniverse := make([][]int, dimension)
	//for element := range nextUniverse {
	//	nextUniverse[element] = make([]int, dimension)
	//}

	createUniverse(currentUniverse)

	for x := 1; x <= generations; x++ {
		nextUniverse := make([][]int, dimension)
		for element := range nextUniverse {
			nextUniverse[element] = make([]int, dimension)
		}
		for i := 0; i < dimension; i++ {
			for j := 0; j < dimension; j++ {
				if currentUniverse[i][j] == 1 {
					count := currentUniverse[i][outOfBorder(j-1)] + currentUniverse[i][outOfBorder(j+1)] + currentUniverse[outOfBorder(i-1)][outOfBorder(j-1)] + currentUniverse[outOfBorder(i-1)][j] + currentUniverse[outOfBorder(i-1)][outOfBorder(j+1)] + currentUniverse[outOfBorder(i+1)][outOfBorder(j-1)] + currentUniverse[outOfBorder(i+1)][j] + currentUniverse[outOfBorder(i+1)][outOfBorder(j+1)]
					if count > 3 || count < 2 {
						nextUniverse[i][j] = 0
						//fmt.Print(" ")
					} else {
						nextUniverse[i][j] = 1
						//fmt.Print("O")
					}
				} else if currentUniverse[i][j] == 0 {
					count := currentUniverse[i][outOfBorder(j-1)] + currentUniverse[i][outOfBorder(j+1)] + currentUniverse[outOfBorder(i-1)][outOfBorder(j-1)] + currentUniverse[outOfBorder(i-1)][j] + currentUniverse[outOfBorder(i-1)][outOfBorder(j+1)] + currentUniverse[outOfBorder(i+1)][outOfBorder(j-1)] + currentUniverse[outOfBorder(i+1)][j] + currentUniverse[outOfBorder(i+1)][outOfBorder(j+1)]
					if count == 3 {
						nextUniverse[i][j] = 1
					}
				}
			}
			//fmt.Println()
		}
		//printUniverse(currentUniverse)
		//fmt.Println("--------")
		//printUniverse(nextUniverse)
		//fmt.Println("--------")
		currentUniverse = nextUniverse
		//printUniverse(currentUniverse)
		//fmt.Println("--------")
		//printUniverse(nextUniverse)
		//fmt.Println("--------")
	}
	printUniverse(currentUniverse)
}
