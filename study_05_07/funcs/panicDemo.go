package funcs

import "fmt"

func PanicDemo() {
	x := 90
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic and recover")
			fmt.Println(p)
		}
	}()
	if x > 80 {
		panic("x > 80")
	}
}


func RecoverDemo() {
	x := 90
	defer func() {
		switch p := recover(); p {
		case nil :
		case "x > 80" :
			fmt.Println("x > 80, recover")
		default :
 			panic("必须终止程序了")
		}		
	}()

	if x > 80 {
		panic("x > 80")
	}
}
