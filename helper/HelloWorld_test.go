package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Anga")
	if result != "Hello Angga" {
		t.Fatal("Result must 'Hello Angga'")
	}

	//fmt.Println("TestHelloWorld Success")
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Anga")
	if result != "Hello Angga" {
		t.Error("Result must 'Hello Angga'")
	}
	//fmt.Println("TestHelloWorld2 Success")
}
