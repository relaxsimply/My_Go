package main

import (
	"errors"
	"fmt"
)

//"runtime/debug"
// "unicode/utf8"

func main() {

	result, err := divide(5, 0)
	if err != nil {
		fmt.Printf("%s ", err)
		//log.Fatalf("divide error: %s", err)
	}
	fmt.Println(result)
}

func someFunction() error {
	return fmt.Errorf("make %s", "fn1")
	//return errors.New(fmt.Sprintf("make %s", "fn1"))
}

func divide(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("divide zero")
	}
	return n1 / n2, nil
}
