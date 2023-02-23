package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// your code here
	merged := append(arrayA, arrayB...) // Menggabungkan 2 array
	result := make([]string, 0)         // Membuat array kosong
	isExist := make(map[string]bool)    // Membuat map kosong

	// Mengisi array kosong dengan array yang sudah digabung
	for _, v := range merged {
		// Jika map kosong, maka isi map dengan nilai true
		if !isExist[v] {
			isExist[v] = true
			// mengisi array kosong dengan nilai dari array yang sudah digabung
			result = append(result, v)
		}
	}
	return result
}

func main() {

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))

}
