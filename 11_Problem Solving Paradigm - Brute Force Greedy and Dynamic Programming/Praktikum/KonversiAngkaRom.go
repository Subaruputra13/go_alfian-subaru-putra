package main

import "fmt"

func convertToRoman(num int) string {
	var result string
	var roman = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var arabic = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i < len(arabic); i++ {
		for num >= arabic[i] {
			result += roman[i]
			num -= arabic[i]
		}
	}
	return result
}

func main() {
	fmt.Println(convertToRoman(4))
	fmt.Println(convertToRoman(9))
	fmt.Println(convertToRoman(23))
	fmt.Println(convertToRoman(2021))
	fmt.Println(convertToRoman(1646))
}
