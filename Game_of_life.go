package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var dimension int
	var seed int64
	fmt.Scan(&dimension, &seed)

	rand.Seed(seed)

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
