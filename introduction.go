package main

import (
	"fmt"
	"math/cmplx"
)

var goIsFun bool = true                       //declaring a variable of type bool
var maxInt uint64 = 1<<64 - 1                 //declaring a variable of type uint64
var complex complex128 = cmplx.Sqrt(-5 + 12i) //declaring a variable of type complex128

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex, complex)
}
