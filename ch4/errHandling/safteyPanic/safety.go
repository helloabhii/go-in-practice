package main

import (
	"errors"
	"time"

	"github.com/Masterminds/cookoo/safely"
)

func message() {
	println("Inside Goroutine")
	panic(errors.New("Oops !!!"))
}

func main() {
	safely.Go(message)
	println("Outside goroutine")
	time.Sleep(1000)
}
