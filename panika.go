package main

import (
	"fmt"
)

//"runtime/debug"
// "unicode/utf8"

func main() {
	//panic("Что-то пошло не так AAAAAAAA!!!")
	defer handlePanic()
	riskyFunction()
	fmt.Println("Эта строка не будет достигнута, если произойдет паника.")
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Обработана паника ", r)
	}

}

func riskyFunction() {
	panic("что-то пошло не так ааааааааааааа!!!!\n")

}
