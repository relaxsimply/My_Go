package main

import (
	"fmt"
)

type Direction int

const (
	_ Direction = iota
	North
	West
	South
	East
)

func main() {
	dit := 5
	action(dir)

}

func action(d Direction) {
	fmt.Println("Действие в направлении:", d)
}
