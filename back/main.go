package main

import (
	"fmt"
	"math/rand"
)

func calculateDiff() int {
	if rand.Int() > 50 {
		fmt.Println("Found ya!")
		fmt.Println("Or not?!")
	}

	return 1
}

func main() {
	fmt.Println("Back")
}
