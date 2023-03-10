package main

import "fmt"

func simpleEquation(a, b, c int) {
	for z := 1; z <= 10000; z++ {
		for x := 1; x <= 10000; x++ {
			y := a - z - x

			if x*x+y*y+z*z == c && x*y*z == b {
				fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
				return
			}
		}
	}
	fmt.Println("No Solution Found")
}

func main() {
	fmt.Print("Persamaan: \n")
	simpleEquation(1, 2, 3)
	simpleEquation(6, 6, 14)
}
