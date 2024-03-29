package main

import "time"

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch)
	for {
		select {
		case <-ch:
			println("Got Message.")
		case <-timeout:
			println("Time Out")
			return
		default:
			println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("Send and closed")
}
