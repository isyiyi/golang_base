package main

import "fmt"

func main() {
	var funcs = [10]func(int){}
	var i = 0
	for i = 0; i < 10; i ++ {
		funcs[i] = func(x int) {
			fmt.Println(x)
		}
	}
	
}


func PrintBool(num int) bool {
	if num > 90 {
		return true
	} else {
		return false
	}
}
