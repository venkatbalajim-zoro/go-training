package day_one

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

var FuncLiteral func() = sayHi
