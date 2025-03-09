package main

import (
	"fmt"
)

func sumar(a int, b int) int {
	return a + b
}

func main() {

	fmt.Println("¡Hola, Mundo!")

	resultado := sumar(2, 4)
	fmt.Println("Resultado:", resultado)

	edad := 13

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad :/")
	}

	frutas := []string{"Manzana", "Banana", "Naranja"}
	fmt.Println("Array de frutas:", frutas)
	frutas = append(frutas, "Pera")
	fmt.Println("Slice de frutas:", frutas)

}
