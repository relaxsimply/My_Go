package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Hello, world")

	var i int = 5555
	var j float64 = 234.1235

	i_str := strconv.Itoa(i)
	j_str := strconv.FormatFloat(j, 'f', 6, 64)

	fmt.Println(i_str, "\n", j_str)
	fmt.Println()
	fmt.Println(string(i_str[1]))
	fmt.Println(string(j_str[3]))

	i_int, err := strconv.Atoi(i_str)
	if err != nil {
		fmt.Println("Ошибка")
	} else {
		fmt.Println(i_int)
	}

	j_int, err := strconv.ParseFloat(j_str, 64)
	if err != nil {
		fmt.Println("Ошибка")
	} else {
		fmt.Println(j_int)
	}

	str := "Привет Андрей йцвйцв"
	fmt.Println(strings.Contains(str, "При"))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.HasPrefix(str, "Gh"))
	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.EqualFold(str, "привет андрей йцвйцв"))

	name := "Andrey"
	age := 123
	// str1 := fmt.Sprint("My name is ", name, " i am ", age)
	// fmt.Println(str1)

	str1 := fmt.Sprintf("My name is %s, i am %d years old", name, age)
	fmt.Println(str1)

	//Спецификаторы
	// %s - string
	// %d - int
	// %f - float
	// %t - bool

	// %d%% - процент
	// %b - двоичная
	// %o - восьмеричная
	// %d - десятичная
	// %x - шестанадцатеричная

	// %.2f - два знака после запятой
	// %v - любое значение

	var a1 rune = 'A'
	fmt.Printf("индекс аски %d, символ %c\n", a1, a1)

	var a2 rune = '\u1234'
	fmt.Printf("индекс аски %d, символ %c\n", a2, a2)

	str5 := "Hello"
	fmt.Printf("число %d, символ %c\n", str5[0], str5[0])

	x := 42
	fmt.Println("Значение x", x)
	fmt.Println("адрес x", &x)

	px := &x
	fmt.Println(*px)
	*px = 100
	fmt.Println(*px)
	fmt.Println(x)

	// ----------------------------------
	var p *int
	fmt.Println(p)

}
