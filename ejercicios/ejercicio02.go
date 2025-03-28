package ejercicios

import "fmt"

func TablaDel(numero int) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}
}
