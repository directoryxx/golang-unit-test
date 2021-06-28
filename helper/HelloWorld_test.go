package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Angga")
	if result != "Hello Angga" {
		panic("Result is not fulfilled")
	}
}
