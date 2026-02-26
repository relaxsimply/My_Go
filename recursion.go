package main

import (
	"fmt"
)

func main() {
	fn(0)
}

func fn(n int) {
	if n > 100 {
		return
	}
	fmt.Println(n, "Привет")
	fn(n + 1)
}
