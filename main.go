package main

import (
	"fmt"

	"github.com/punigonzalez/go-pruebas/ejercicios"
)

func main() {
	// Prueba la función con diferentes valores
	valor1 := "150"
	resultado1 := ejercicios.CompararConCien(valor1)
	fmt.Printf("Resultado para %s: %s\n", valor1, resultado1)

	valor2 := "75"
	resultado2 := ejercicios.CompararConCien(valor2)
	fmt.Printf("Resultado para %s: %s\n", valor2, resultado2)
}
