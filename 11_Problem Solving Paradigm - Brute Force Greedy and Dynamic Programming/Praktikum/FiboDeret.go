package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func sumFibo(n int) int {
	var result int
	for i := 0; i <= n; i++ {
		if i == 8 {
			continue
		}
		result += fibo(i)
	}
	return result
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(4))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
	fmt.Println()

	fmt.Print("Sum of Fibo: \n")
	fmt.Println(sumFibo(10))
}
