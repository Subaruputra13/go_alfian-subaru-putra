package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	//membuat slice kosong
	var result []int

	//membuat map kosong
	var temp = make(map[int]int)

	//perulangan untuk setiap value di slice
	for _, value := range angka {
		// mengubah tipe data string ke int
		n, _ := strconv.Atoi(string(value))
		temp[n]++
	}

	//perulangan untuk setiap key dan value di map
	for key, value := range temp {
		//jika value sama dengan 1 maka akan di append ke slice result
		if value == 1 {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // Output :[4]
	fmt.Println(munculSekali("76523752"))   // Output :[6 3]
	fmt.Println(munculSekali("12345"))      // Output :[1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // Output :[]
	fmt.Println(munculSekali("0872504"))    // Output :[8 7 2 5 4]
}
