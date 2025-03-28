package iteraciones

import (
	"fmt"
)

func IterarXveces(i int) {
	for ; i < 10; i += 2 {
		fmt.Println(i)
	}
}
