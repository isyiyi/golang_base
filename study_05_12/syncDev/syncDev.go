package syncDev

import (
	"fmt"
)

func Funcs() {
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
}
