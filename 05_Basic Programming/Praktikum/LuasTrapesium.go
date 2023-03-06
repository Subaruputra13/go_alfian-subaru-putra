package main

import "fmt"

func main() {
	var sisi1, sisi2, tinggi float64
	sisi1 = 16
	sisi2 = 10
	tinggi = 8

	Luas := 0.5 * (sisi1 + sisi2) * tinggi

	fmt.Println("Luas Trapesium adalah ", Luas)
}
