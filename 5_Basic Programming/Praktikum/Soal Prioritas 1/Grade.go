package main

import "fmt"

func main() {
	nilai := 101

	if nilai < 100 && nilai >= 80 {
		fmt.Println("Nilai Anda A")
	} else if nilai <= 79 && nilai >= 65 {
		fmt.Println("Nilai Anda B")
	} else if nilai <= 64 && nilai >= 50 {
		fmt.Println("Nilai Anda C")
	} else if nilai <= 49 && nilai >= 35 {
		fmt.Println("Nilai Anda D")
	} else if nilai <= 34 && nilai >= 0 {
		fmt.Println("Nilai Anda E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
