package main

import "fmt"

func Fibo(n int, list map[int]int) int {
	if n < 2 {
		return n
	}
	if val, ok := list[n]; ok {
		return val
	}

	list[n] = Fibo(n-1, list) + Fibo(n-2, list)
	return list[n]
}

func sumFibTopDown(n int, list map[int]int) int {
	var result int
	for i := 0; i <= n; i++ {
		if i == 8 {
			continue
		}
		result += Fibo(i, list)
	}
	return result
}

func main() {
	fmt.Print("fibTopDown: \n")
	fmt.Println(Fibo(0, map[int]int{}))
	fmt.Println(Fibo(1, map[int]int{}))
	fmt.Println(Fibo(2, map[int]int{}))
	fmt.Println(Fibo(3, map[int]int{}))
	fmt.Println(Fibo(4, map[int]int{}))
	fmt.Println(Fibo(5, map[int]int{}))
	fmt.Println(Fibo(6, map[int]int{}))
	fmt.Println(Fibo(7, map[int]int{}))
	fmt.Println(Fibo(9, map[int]int{}))
	fmt.Println(Fibo(10, map[int]int{}))

	fmt.Println()

	fmt.Print("sumFibTopDown: \n")
	fmt.Println(sumFibTopDown(10, map[int]int{}))

}
