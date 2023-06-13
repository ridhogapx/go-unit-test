package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	t.Run("First", func(t *testing.T) {
		result := Hello("Kirito")
		assert.Equal(t, "Hello Kirito", result, "Value is must be Hello Kirito")
	})
	t.Run("Second", func(t *testing.T) {
		result := Hello("John")
		assert.Equal(t, "Hello John", result, "Value is must be Hello Kirito")
	})
	// Trying to failed it
	t.Run("Third", func(t *testing.T) {
		result := Hello("Doe")
		assert.Equal(t, "Hi Doe", result, "Value is equal to Hello Doe")
	})
}

//func TestMain(m *testing.M) {
// Before unit test
//fmt.Println("Before unit test...")
//m.Run()
// After unit test
//fmt.Println("After unit test...")
// }

func TestAssertion(t *testing.T) {
	result := Hello("RageNeko26")
	assert.Equal(t, "Hello RageNeko26", result, "Result must be 'Hello RageNeko26' ")
}

func TestOther(t *testing.T) {
	var result = Hello("Budi")
	assert.Equal(t, "Hello Budi", result, "Result must be same as expected 'Hello Budi!' ")
}

func TestExample(t *testing.T) {
	result := Hello("Anon")
	assert.Equal(t, "Hi Anon!", result, "Result is not as expected!")
}
