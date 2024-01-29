package main

import "log"

func main() {
	log.Println("this is a regular message.")
	log.Fatalln("this is a fatal error.")
	log.Println("this is the end of the function") // this will never be print
}
