package main

import "fmt"

func main() {
	ang := 5

	for i := 1; i <= ang; i++ {
		for s := i; s < ang; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}
