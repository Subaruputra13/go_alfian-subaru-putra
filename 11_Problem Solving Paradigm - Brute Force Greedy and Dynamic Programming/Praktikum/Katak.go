package main

import "fmt"

func Frog(jumps []int) int {
	dp := make([]int, len(jumps))
	dp[0] = 0

	for i := 1; i < len(jumps); i++ {
		cost2 := 0
		cost1 := dp[i-1] + Abs(jumps[i]-jumps[i-1])

		if i > 1 {
			cost2 = dp[i-2] + Abs(jumps[i]-jumps[i-2])
		} else {
			cost2 = cost1
		}
		dp[i] = Min(cost1, cost2)
	}
	return dp[len(jumps)-1]
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
