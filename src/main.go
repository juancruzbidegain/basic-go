package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// llave valor del map

	for i, value := range m {
		fmt.Println(i, value)
	}

	// value := m["1Jose"]
	// if value == 0 {
	// 	fmt.Println("No se encontro la llave solicitada")
	// } else {
	// 	fmt.Println(value)
	// }

	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
