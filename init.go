package main

import (
	"fmt"
)

var config string

func init() {
	config = "Настройки загружены"
	fmt.Println("Инициализация пакета")
}

func main() {
	fmt.Println(config)
}
