package main

import "fmt"

func main() {
	// ang := 100
	// count := 0
	// for i := 1; i <= ang; i++ {
	// 	if i%2 == 1 {
	// 		fmt.Println(i)
	// 		count++
	// 	}
	// }

	// fmt.Printf("Bilangan ganjil 1 - 100 ada %d", count)

	//Bilangan Prima 1- 100
	var ang int
	fmt.Print("Masukan Angka : ")
	fmt.Scanln(&ang)

	for i := 2; i < ang; i++ {

		if i%ang == 0 {
			fmt.Println()
		}
	}
}
