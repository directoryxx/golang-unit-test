package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Anga")
	if result != "Hello Angga" {
		t.Fail()
	}

	fmt.Println("TestHelloWorld Success")
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Anga")
	if result != "Hello Angga" {
		t.FailNow()
	}
	// t.FailNow program dibawahnya tidak dijalankan
	fmt.Println("TestHelloWorld2 Success")
}
