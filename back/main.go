package main

import (
	"fmt"
	"math/rand"
)

func calculateDiff() int {
	if rand.Int() > 50 {
		fmt.Println("Found ya!")
	}

	return 1
}

func main() {
	fmt.Println("Back")
}
