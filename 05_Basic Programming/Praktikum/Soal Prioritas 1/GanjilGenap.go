package main

import "fmt"

func main() {
	bil := 5

	if bil%2 == 0 {
		fmt.Println(bil, "adalah Bilangan Genap")
	} else {
		fmt.Println(bil, "adalah Bilangan Ganjil")
	}
}
