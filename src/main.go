package main

import (
	"fmt"
)

func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		fmt.Println(string(text[i]))
		textReverse += string(text[i])
	}

	fmt.Println(text, textReverse)

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo...")
	}
}

func main() {

	// slice := []string{"Hola", "Que", "Hace"}
	// for i, value := range slice {
	// 	fmt.Println(i, value)
	// }

	// for _, x := range slice {
	// 	fmt.Println(x)
	// }

	// Funcion para un palindromo

	isPalindromo("amor a roma")

}
