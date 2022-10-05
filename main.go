package main

import "log"

func SayHello(name string) string {
	return "Hello " + name
}

func main() {
	response := SayHello("vijeth")
	log.Println(response)
}
