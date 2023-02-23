package main

import "fmt"

func Mapping(slice []string) map[string]int {
	// buat map kosong
	result := make(map[string]int)

	// melakukan perulangan untuk setiap value di slice
	for _, value := range slice {
		// menjumlahkan value yang sama di map
		result[value]++
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))

	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))

	fmt.Println(Mapping([]string{}))
}
