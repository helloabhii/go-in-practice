package main

import (
	"log"
	"net"
	"time"
)

func main() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", ":1902", timeout)
	if err != nil {
		panic("Failed to connect  to localhost: 1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Ltime
	logger := log.New(conn, "example", f)

	logger.Println("This is a regular message")
	logger.Panicln("This is a panic.")
}
