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
		x = 1
	}
	return x
}

func main() {

	fmt.Println("Insert dimension and seed of the universe:")
	fmt.Scan(&dimension, &seed, &generations)

	rand.Seed(int64(seed))

	universe := make([][]int, dimension)
	for element := range universe {
		universe[element] = make([]int, dimension)
	}

	gen2 := make([][]int, dimension)
	for element := range gen2 {
		gen2[element] = make([]int, dimension)
	}

	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			x := rand.Intn(2)
			universe[i][j] = x
			if x == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}

	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if universe[i][j] == 1 {
				count := universe[i][outOfBorder(j-1)] + universe[i][outOfBorder(j+1)] + universe[outOfBorder(i-1)][outOfBorder(j-1)] + universe[outOfBorder(i-1)][j] + universe[outOfBorder(i-1)][outOfBorder(j+1)] + universe[outOfBorder(i+1)][outOfBorder(j-1)] + universe[outOfBorder(i+1)][j] + universe[outOfBorder(i+1)][outOfBorder(j+1)]
				if count > 3 || count < 2 {
					gen2[i][j] = 0
					fmt.Print(" ")
				} else {
					gen2[i][j] = 1
					fmt.Print("O")
				}
			}
		}
		fmt.Println()
	}

}
