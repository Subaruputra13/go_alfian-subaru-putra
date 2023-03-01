package main

import "fmt"

func PointerMaxMin(n ...*int) (max, min int) {
	max = *n[0]
	min = *n[0]

	for _, v := range n {
		if *v < max {
			max = *v
		} else if *v > min {
			min = *v
		}
	}
	return max, min
}

func main() {
	var a, b, c, d, e, f int = 1, 2, 3, 9, 8, 7
	var max, min int
	// fmt.Scan(&a)
	// fmt.Scan(&b)
	// fmt.Scan(&c)
	// fmt.Scan(&d)
	// fmt.Scan(&e)
	// fmt.Scan(&f)
	min, max = PointerMaxMin(&a, &b, &c, &d, &e, &f)
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
	fmt.Printf("%d\n", e)
	fmt.Printf("%d\n", f)
	fmt.Println("Nilai terendah :", min)
	fmt.Println("Nilai tertinggi :", max)
}
