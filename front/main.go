package main

import "fmt"

func renderScreen() bool {
	fmt.Println("Rendering")
	fmt.Println("Rendering more")
	fmt.Println("Rendering even more")
	return true
}

func main() {
	fmt.Println("Front")
}
