package main

import "fmt"

func main() {
	// For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println("El numero de este ciclo es: ", i)
	}

	// For while
	counter := 0
	for counter <= 10 {
		fmt.Println("Counter numero: ", counter)
		counter++
	}

	// For forever

	counterForever := 0
	for {
		fmt.Println("Counter Forever Escuchando: ... ", counterForever)
		counterForever++
	}
}
