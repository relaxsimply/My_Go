package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	Monday = iota
	Tuesday
	Wensday
	Thusday
)

const (
	a = iota + 6
	b
	c
	d
)

func main() {
	var age int = 30
	var name string = "Andrey"
	fmt.Println(a, b, c, age)
	age = 35
	fmt.Println(a, b, c, age)

	const PI = 3.123128719891287234702394702394720394720934720934
	fmt.Println(PI)
	fmt.Printf("Hello, %s you are %d", name, age)
	fmt.Println()

	fmt.Println(a, b, c, d)

	// random_num := rand.IntN(100)
	// fmt.Println(random_num)

	// min := 10
	// max := 500
	// rand := rand.IntN(max - min)
	// fmt.Print(rand)

	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Fatal("Ошибка", err.Error())
	}

	fmt.Println(n.Int64())
	fmt.Println(math.Pow(2, 3))

	fmt.Println(math.Round(2.123123))
	// fmt.Println(int(2.123123))
	// var totalPrice float64 = 2.12312323
	// fmt.Println(int(totalPrice))

	// var str1 string
	// fmt.Println("str1", str1)

	// var str2 string = `"Andrey
	// awdawd
	// awdawd
	// awd
	// Kotov"`
	// fmt.Println("str2", str2)
	var str string = "тцшгашгуашгр"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
	str5 := "awdawdadwad \uFAFA"
	fmt.Println(str5)

	str6 := "awdwadawd"
	fmt.Println(strings.ToUpper(str6))
	//strings.ToLower
	//strings.TrimSpace
	//strings.Split

	//Содержит ли строка нужные символы strings.Contains(str, "нужные символы")

	// Разбить строку по символу в слайс strings.Split(str, " ")

	//Имеет ли строка префикс strings.HasPrefix(str, "Hello, world")
	//Заканчивается на что-то строка strings.HasSuffix(str, "world!")

	//равны ли строчки не учитывая регистра strings.EqualFold(str, "awdawdaw")

	//Вывести число как строку

	i := 555
	s1 := strconv.Itoa(i)
	fmt.Println(s1)

	var i2 float64 = 444.123123
	var i3 = strconv.FormatFloat(i2, 'f', 6, 64)
	fmt.Println(i3)

	//Приведение строки к числу

	str10 := "123123.123"
	//str10 := "13123"
	//num, err := strconv.Atoi(str10)
	//num, err := strconv.ParseInt(str10, 10, 64)
	num, err := strconv.ParseFloat(str10, 64)
	if err != nil {
		fmt.Println("Ошибка", err)
	} else {
		fmt.Println(num)
	}
}
