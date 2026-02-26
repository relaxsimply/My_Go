package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	// fmt.Println("Hello, world")

	// var i int = 5555
	// var j float64 = 234.1235

	// i_str := strconv.Itoa(i)
	// j_str := strconv.FormatFloat(j, 'f', 6, 64)

	// fmt.Println(i_str, "\n", j_str)
	// fmt.Println()
	// fmt.Println(string(i_str[1]))
	// fmt.Println(string(j_str[3]))

	// i_int, err := strconv.Atoi(i_str)
	// if err != nil {
	// 	fmt.Println("Ошибка")
	// } else {
	// 	fmt.Println(i_int)
	// }

	// j_int, err := strconv.ParseFloat(j_str, 64)
	// if err != nil {
	// 	fmt.Println("Ошибка")
	// } else {
	// 	fmt.Println(j_int)
	// }

	// str := "Привет Андрей йцвйцв"
	// fmt.Println(strings.Contains(str, "При"))
	// fmt.Println(strings.ToUpper(str))
	// fmt.Println(strings.HasPrefix(str, "Gh"))
	// fmt.Println(strings.TrimSpace(str))
	// fmt.Println(strings.EqualFold(str, "привет андрей йцвйцв"))

	// name := "Andrey"
	// age := 123
	// // str1 := fmt.Sprint("My name is ", name, " i am ", age)
	// // fmt.Println(str1)

	// str1 := fmt.Sprintf("My name is %s, i am %d years old", name, age)
	// fmt.Println(str1)

	// //Спецификаторы
	// // %s - string
	// // %d - int
	// // %f - float
	// // %t - bool

	// // %d%% - процент
	// // %b - двоичная
	// // %o - восьмеричная
	// // %d - десятичная
	// // %x - шестанадцатеричная

	// // %.2f - два знака после запятой
	// // %v - любое значение

	// var a1 rune = 'A'
	// fmt.Printf("индекс аски %d, символ %c\n", a1, a1)

	// var a2 rune = '\u1234'
	// fmt.Printf("индекс аски %d, символ %c\n", a2, a2)

	// str5 := "Hello"
	// fmt.Printf("число %d, символ %c\n", str5[0], str5[0])

	// x := 42
	// fmt.Println("Значение x", x)
	// fmt.Println("адрес x", &x)

	// px := &x
	// fmt.Println(*px)
	// *px = 100
	// fmt.Println(*px)
	// fmt.Println(x)

	// // ----------------------------------
	// var p *int
	// fmt.Println(p)

	// // commit 25.02.2026

	// var buf strings.Builder
	// buf.WriteString("Hello")
	// buf.WriteRune('f')
	// buf.WriteString("world")

	// line := buf.String()
	// fmt.Println(line)

	// // 	&& - и
	// // 	|| - или
	// //  ! - не

	// num := 40
	// fmt.Printf("%b\n", num)
	// fmt.Printf("%b\n", num<<2)

	// if num14 := rand.IntN(100); num > 50 {
	// 	fmt.Printf("Выпало число %d!\n", num14)
	// } else {
	// 	fmt.Printf("lol\n")
	// }

	// // switch
	// day := 2

	// switch day {
	// case 1:
	// 	fmt.Println("Понедельник")
	// case 2:
	// 	fmt.Println("Вторник")
	// case 3:
	// 	fmt.Println("Среда")
	// default:
	// 	fmt.Println("Неизвестный день")
	// }

	// switch {
	// case day < 1:
	// 	fmt.Println("Non correct")
	// case day >= 1 && day < 4:
	// 	fmt.Println("correct")
	// }

	// switch day {
	// case 1, 2, 3, 4, 5, 6:
	// 	fmt.Println("correct")
	// case 10:
	// 	fmt.Println("Non correct")
	// }

	// switch day {

	// case 1:
	// 	fmt.Println("Понедельник")
	// case 2:
	// 	fmt.Println("Вторник после чего будет")
	// 	fallthrough
	// case 3:
	// 	fmt.Println("Среда")
	// }

	// // Определение времени суток
	// var m int
	// fmt.Println("Введите аремя суток")

	// _, err := fmt.Scan(&m)
	// if (err != nil) || (m < 0) || (m > 23) {
	// 	fmt.Println("введены некорректные данные")
	// }

	// switch {
	// case 6 >= m && m >= 12:
	// 	fmt.Println("Утро", m)
	// case 12 >= m && m >= 18:
	// 	fmt.Println("День", m)
	// case 18 >= m && m >= 23:
	// 	fmt.Println("Вечер", m)
	// case 23 >= m && m >= 6:
	// 	fmt.Println("Ночь", m)
	// }

	// Калькулятор ИМТ
	var mass int
	var len float64

	fmt.Println("Введите ваш вес (кг)")
	_, err1 := fmt.Scan(&mass)
	if err1 != nil || mass < 0 || mass > 200 {
		fmt.Println("Некорректные данные")
		return
	}

	fmt.Println("Введите ваш рост (см)")
	_, err2 := fmt.Scan(&len)
	if err2 != nil || len < 0 || len > 210 {
		fmt.Println("Некорректные данные")
		return
	}
	len /= 100
	var imt float64 = (float64(mass) / (math.Pow(len, 2)))
	fmt.Print("Ваш ИМТ: ", imt)

	var cat string = "non"

	switch {
	case imt < 18.5:
		cat = "Недостаточный вес"
	case (18.5 <= imt && imt < 25):
		cat = "Нормальный вес"
	case imt >= 30:
		cat = "Ожирение"
	}
	fmt.Print(" Категория: ", cat)

	// Поиск по названию товара
	fmt.Println("Введите название товара")
	var name string
	_, err3 := fmt.Scan(&name)
	if err3 != nil || name == "" || name == " " {
		fmt.Println("Ошибка")
		return
	}

	var klava string = "Клавиатура JZ9"
	klava = strings.ToLower(klava)
	var headphons string = "Наушники N45"
	klava = strings.ToLower(headphons)
	var smartphone string = "Смартфон S10"
	klava = strings.ToLower(smartphone)

	name = strings.ToLower(name)
	switch {

	case (strings.Contains(klava, name)):
		fmt.Println(19200)
	case (strings.Contains(headphons, name)):
		fmt.Println(9600)
	case (strings.Contains(smartphone, name)):
		fmt.Println(55000)
	default:
		fmt.Printf("Товар %s не найден", name)
	}
}
