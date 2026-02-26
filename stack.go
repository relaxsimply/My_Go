package main

import (
	"fmt"
	"runtime/debug"
	// "unicode/utf8"
)

// var name = "Andrey"

func main() {
	fn1()

}

func fn1() {
	fn2()
	fmt.Println("fn1")
}

func fn2() {
	fmt.Println("fn2")
	debug.PrintStack()
}
