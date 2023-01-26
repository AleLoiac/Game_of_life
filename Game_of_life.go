package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Insert dimension and seed of the universe:")
	var dimension int
	var seed int
	fmt.Scan(&dimension, &seed)

	rand.Seed(int64(seed))

	universe := make([][]int, dimension)
	for element := range universe {
		universe[element] = make([]int, dimension)
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

}
