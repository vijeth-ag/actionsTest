package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	expected := "Hello vijeth"
	actual := SayHello("vijeth")

	if expected != actual {
		t.Errorf("Failing")
	}
}
