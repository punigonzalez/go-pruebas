package ejercicios

import (
	"fmt"
	"strconv"
)

func CompararConCien(valor string) string {
	valorInt, err := strconv.Atoi(valor)
	if err != nil {
		return "Error: No se pudo convertir a entero"
	}

	if valorInt > 100 {
		fmt.Println("El valor es mayor a 100")
		return "El valor es mayor a 100"
	} else {
		fmt.Println("El valor es menor o igual a 100")
		return "El valor es menor o igual a 100"
	}
}
