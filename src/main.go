package main

import "fmt"

func main() {

	//Constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("Pi", pi)
	fmt.Println("Pi2", pi2)

	// Variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area Cuadrado BASE * ALTURA

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

}
