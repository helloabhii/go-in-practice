package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":1902")
	if err != nil {
		panic("Failed to connect to localhost: 1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Ltime
	logger := log.New(conn, "example", f)

	logger.Println("This is a regular message")
	logger.Panicln("This is a panic")

	// should run like - nc -lk 1902
}
