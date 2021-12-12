package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {

	car1 := car{brand: "Ford", year: 2020}
	fmt.Println(car1)

	// Otra manera
	var car2 car
	car2.brand = "Ferrari"
	fmt.Println(car2)
}
