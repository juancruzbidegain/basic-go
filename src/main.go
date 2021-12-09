package main

import "fmt"

func main() {

	// Paquete FMT
	helloMenssage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMenssage, worldMessage)
	fmt.Println(helloMenssage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

	//Spintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("HelloMessage: %T \n", helloMenssage)
	fmt.Printf("cursos: %T \n", cursos)
}
