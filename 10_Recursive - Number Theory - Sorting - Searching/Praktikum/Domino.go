package main

import "fmt"

func playDomino(cards [][]int, deck []int) interface{} {
	for _, v := range cards {
		if v[0] == deck[0] || v[1] == deck[0] || v[0] == deck[1] || v[1] == deck[1] {
			return v
		}
	}
	return "Skip kartu"
}

func main() {
	fmt.Println(playDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playDomino([][]int{{6, 5}, {3, 3}, {2, 1}, {3, 6}}, []int{3, 6}))
	fmt.Println(playDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}
