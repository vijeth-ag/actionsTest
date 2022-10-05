package main

import (
	"log"
	"testing"
)

func TestSayHello(t *testing.T) {
	expected := "Hello vijeth"
	actual := SayHello("vijeth")

	log.Println("ex", expected)
	log.Println("actual", actual)
}
