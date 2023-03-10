package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func getBilBulat(x int) {
	for i := 1; i < 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%d at time %v\n", i, time.Since(start))
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("main execution started at time", time.Since(start))
	fmt.Println()

	go getBilBulat(14)

	time.Sleep(3 * time.Second)
	fmt.Println("\nmain execution ended at time", time.Since(start))
}
