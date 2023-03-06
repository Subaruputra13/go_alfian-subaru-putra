package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	nama   string
	jumlah int
}

func MostAppearItem(items []string) []Pair {
	var result []Pair
	var temp []string

	// Fungsi untuk mengecek apakah item sudah ada di slice atau belum
	IsExist := func(slice []string, item string) bool {
		for _, s := range slice {
			if s == item {
				return true
			}
		}
		return false
	}

	//perulang untuk mengecek apakah item sudah ada di slice temp atau belum
	for _, item := range items {
		if !IsExist(temp, item) {
			temp = append(temp, item)
		}
	}

	//perulangan untuk menghitung jumlah kemunculan item
	for _, item1 := range temp {
		var pair Pair
		pair.nama = item1
		pair.jumlah = 0
		for _, item2 := range items {
			if item1 == item2 {
				pair.jumlah++
			}
		}
		result = append(result, pair)
	}

	//fungsi sort.Slice untuk mengurutkan slice dari ascending
	sort.Slice(result, func(i, j int) bool {
		return result[i].jumlah < result[j].jumlah
	})

	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})) // [golang 1 ruby 2 js 4]
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})) // [C 1 D 2 A 3 B 3]
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
