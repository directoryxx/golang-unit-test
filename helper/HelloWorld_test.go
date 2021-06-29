package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Angga")
	}
}

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("Angga", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Angga")
		}
	})

	b.Run("Wijaya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Wijaya")
		}
	})

}

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

func TestSubTest(t *testing.T) {
	t.Run("Angga", func(t *testing.T) {
		result := HelloWorld("Angga")
		require.Equal(t, "Hello Angga",result)
	})

	t.Run("Rudi", func(t *testing.T) {
		result := HelloWorld("Rudi")
		require.Equal(t, "Hello Rudi",result)
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name : "angga",
			request: "angga",
			expected: "Hello angga",
		},
		{
			name : "rudi",
			request: "rudi",
			expected: "Hello rudi",
		},
	}

	for _,test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t,test.expected,result)
		})
	}
}