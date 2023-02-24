package main

import "fmt"

func main() {
	MatrixArray := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	DiagonalRigth := 0
	DiagonalLeft := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				DiagonalRigth += MatrixArray[i][j]
			}
			if i+j == 2 {
				DiagonalLeft += MatrixArray[i][j]
			}
		}
	}

	fmt.Println("Diagonal :", DiagonalRigth)
	fmt.Println("AntiDiagonal :", DiagonalLeft)
	fmt.Println("Perbedaan Mutlak :", DiagonalLeft-DiagonalRigth)
}
