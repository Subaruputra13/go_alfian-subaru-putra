package main

import "fmt"

func GetPrime(n, i int) bool {
	if i == 1 {
		return true
	} else if n%i == 0 {
		return false
	} else {
		return GetPrime(n, i-1)
	}
}

func getNthPrime(ang, c int) int {
	if GetPrime(c, c-1) {
		if ang == 1 {
			return c
		} else {
			return getNthPrime(ang-1, c+1)
		}
	} else {
		return getNthPrime(ang, c+1)
	}
}

func primeX(n int) int {
	return getNthPrime(n, 2)
}

func main() {

	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
