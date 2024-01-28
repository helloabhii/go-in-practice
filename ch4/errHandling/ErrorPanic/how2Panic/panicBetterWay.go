package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() { // annynomyous function
		if err := recover(); err != nil {
			fmt.Printf("Trapped Panic: %s (%T)\n", err, err)
		}
	}()
	yikes()
}

func yikes() {
	panic(errors.New("Something bad happend."))
}
