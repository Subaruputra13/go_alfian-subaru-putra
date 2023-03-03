package main

import "fmt"

func fibonacci(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	fmt.Println(fibonacci(6))
}
