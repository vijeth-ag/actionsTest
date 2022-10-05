package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	expected := "Hellox vijeth"
	actual := SayHello("vijeth")

	if expected != actual {
		t.Errorf("Failing")
	}
}
