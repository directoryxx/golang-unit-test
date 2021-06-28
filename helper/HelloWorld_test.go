package helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Anga")
	// Dibawah require tidak dijalankan
	// Ex: TestHelloWorld3
	require.Equal(t, "Hello Angga",result,"Result must equal 'Hello Angga'")
	//fmt.Println("TestHelloWorld2 Success")
}

func TestHelloWorld3(t *testing.T) {
	result := HelloWorld("Angga")
	assert.Equal(t, "Hello Angga",result,"Result must equal 'Hello Angga'")
}


