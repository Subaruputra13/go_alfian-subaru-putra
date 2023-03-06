package main

import "fmt"

func PairSum(arr []int, target int) []int {

	//membuat slice kosong untuk menampung hasil
	var result []int

	//looping untuk mengecek setiap elemen di array
	for i := 0; i < len(arr); i++ {
		//Looping untuk mengecek setiap elemen di array
		for j := i; j < len(arr); j++ {

			//jika elemen pertama ditambah elemen kedua sama dengan target
			if arr[i]+arr[j] == target {

				//masukkan kedalam slice result
				result = append(result, i, j)

				//mengembalikan nilai result
				return result
			}
		}
	}
	return result
}

func main() {

	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // Output : [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // Output : [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // Output : [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // Output : [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // Output : [0, 1]

}
