package main

import "fmt"

type Car struct {
	typeCar string
	fuelIn  int
}

func (car Car) hitungJarank() float64 {
	return float64(car.fuelIn) / (1.5)
}

func main() {
	car := Car{
		typeCar: "SUV",
		fuelIn:  65,
	}
	fmt.Println(car.hitungJarank())
}
