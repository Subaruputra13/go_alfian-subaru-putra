package main

import (
	"fmt"
)

func kelipatan(kel chan int) {
	ang := <-kel

	for i := 1; i <= ang; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("Start Main")
	c := make(chan int)

	go kelipatan(c)
	c <- 25

	fmt.Println("End Main")
}
