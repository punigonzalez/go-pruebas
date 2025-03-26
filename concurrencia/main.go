package concurrencia

import (
	"fmt"
	"time"
)

func decirHola(canal chan<- string) {
	time.Sleep(1 * time.Second)
	canal <- "Hola desde la GoRoutine"
}

func imprimirMensaje(canal <-chan string) {
	msg := <-canal
	fmt.Println(msg)
}

func main() {

}
