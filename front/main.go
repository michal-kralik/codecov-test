package main

import "fmt"

func renderScreen() bool {
	fmt.Println("Rendering")
	fmt.Println("Rendering more")
	return true
}

func main() {
	fmt.Println("Front")
}
