package main

import "fmt"

func main() {
	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println(result)

	// Resta
	result = y - x
	fmt.Println(result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	//Resto
	result = y % x
	fmt.Println("Modulo / Resto:", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

}
