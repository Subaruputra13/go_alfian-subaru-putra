package main

import "fmt"

func main() {
	fmt.Print("Hello, ")
	fmt.Println(sum(1, 2))
}

func sum(a int, b int) int {
	return a + b
}
