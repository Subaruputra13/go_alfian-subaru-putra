package main

import "fmt"

func main() {
	fmt.Println("Start Main")
	result := make(chan int, 10) // buffer channel

	go func() {
		for i := 1; i < 20; i++ {
			result <- i
		}
		close(result)
	}()

	for value := range result {
		if value%3 == 0 {
			fmt.Println(value)
		}
	}

	fmt.Println("End Main")
}
