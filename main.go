package main

import (
	"fmt"

	"github.com/punigonzalez/go-pruebas/ejercicios"
	"github.com/punigonzalez/go-pruebas/funciones"
	"github.com/punigonzalez/go-pruebas/iteraciones"
	"github.com/punigonzalez/go-pruebas/teclado"
)

func main() {
	// Prueba la función con diferentes valores
	valor1 := "101"
	ejercicios.CompararConCien(valor1)

	valor2 := "q00"
	ejercicios.CompararConCien(valor2)

	//ingreso por teclado
	teclado.IngresoNumeros()

	iteraciones.IterarXveces(1)

	// saber la tabla de un numero del 1 al 10
	var input int
	fmt.Println("Ingrese numero para saber su tabla: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return
	} else {
		ejercicios.TablaDel(input)
	}

	// probando funciones

	funciones.Calculos()
}
