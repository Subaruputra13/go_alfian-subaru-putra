package main

import "fmt"

func generateTriangle(numsRows int) [][]int {
	segi := make([][]int, numsRows)

	for i := range segi {
		segi[i] = make([]int, i+1)
		segi[i][0], segi[i][i] = 1, 1

		for j := 1; j < i; j++ {
			segi[i][j] = segi[i-1][j-1] + segi[i-1][j]
		}
	}
	return segi
}

func main() {
	fmt.Println(generateTriangle(5))
}
