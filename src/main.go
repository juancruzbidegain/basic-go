package main

import "fmt"

func normalFunction(msg string) {
	fmt.Println(msg)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Buenas tardes")
	tripeArgument(1, 2, "Hola String")
	value := returnValue(100)
	fmt.Println(value)

	value1, _ := doubleReturn(2)

	fmt.Println("Value1 & Value2: ", value1)
}
