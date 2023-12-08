package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println(safeDiv(1, 0))  // 0 runtime error: integer divide by zero
	fmt.Println(safeDiv(18, 3)) // 6 <nil>

}

func safeDiv(a, b int) (res int, err error) {
	// res and err are local vars and automatically returned

	defer func() {
		// The recover built-in function allows a program to manage behavior of a panicking goroutine
		// e's type is `any (or interface{})` and not `error` so we have to format it to error while assigning it to `err` var
		if e := recover(); e != nil {
			log.Print("ERROR:", e)
			err = fmt.Errorf("%v", e)

		}

	}()

	return a / b, nil
}
