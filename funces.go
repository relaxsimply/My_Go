package main

import (
	"fmt"
	"unicode/utf8"
)

// var name = "Andrey"

func main() {

	fmt.Println("Сейчас будем работать с функциями")

	var name = "Andrey Kotov"
	sayHello(name)

	fmt.Println(sum(100, 200))

	helloStr := "Привет, друг!"
	getFull(helloStr)
	bytes_len, runes_len := getFull(helloStr)
	fmt.Printf("Строка \"%s\" имеет длину в байтах %d и количество символов %d", helloStr, bytes_len, runes_len)

	fmt.Println()
	val := 10

	md_vl(val)
	fmt.Println(val)

	md_ptr(&val)
	fmt.Println(val)

	divide(10, 0)

}

func sayHello(username string) {
	fmt.Println("Привет", username)
}

func sum(n1, n2 int) int {
	return n1 + n2
}

func getFull(str string) (int, int) {
	bytes := len(str)
	runes := utf8.RuneCountInString(str)
	return bytes, runes
}

func md_vl(a int) {
	a = a + 10
}

func md_ptr(a *int) {
	*a = *a + 10
}

func divide(n1, n2 int) {
	defer fmt.Println("Конец функции divide")
	fmt.Println(n1 / n2)
	//
}
