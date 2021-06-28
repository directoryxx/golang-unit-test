package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Angga")
	if result != "Hello Angga" {
		t.Fatal("Result must 'Hello Angga'")
	}

	//fmt.Println("TestHelloWorld Success")
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Angga")
	if result != "Hello Angga" {
		t.Error("Result must 'Hello Angga'")
	}
	//fmt.Println("TestHelloWorld2 Success")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Angga")
	// Dibawah require tidak dijalankan
	// Ex: TestHelloWorld3
	require.Equal(t, "Hello Angga",result,"Result must equal 'Hello Angga'")
	//fmt.Println("TestHelloWorld2 Success")
}

func TestHelloWorld3(t *testing.T) {
	result := HelloWorld("Angga")
	assert.Equal(t, "Hello Angga",result,"Result must equal 'Hello Angga'")
}

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "darwin"{
		t.Skip("Test tidak dijalankan di mac")
	}

	result := HelloWorld("Angga")
	assert.Equal(t, "Hello Angga",result,"Result must equal 'Hello Angga'")
}

func TestMain(m *testing.M) {
	// BIsa digunakan init db
	fmt.Println("Before Unit Test")
	m.Run()

	// BIsa digunakan cleanup db
	fmt.Println("After Unit Test")
}


