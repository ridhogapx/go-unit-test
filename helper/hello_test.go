package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {

}

func TestMain(m *testing.M) {
	// Before unit test
	fmt.Println("Before unit test...")
	m.Run()
	// After unit test
	fmt.Println("After unit test...")
}

func TestAssertion(t *testing.T) {
	result := Hello("RageNeko26")
	assert.Equal(t, "Hello RageNeko26", result, "Result must be 'Hello RageNeko26' ")
}

func TestOther(t *testing.T) {
	var result = Hello("Budi")
	assert.Equal(t, "Hello Budi", result, "Result must be same as expected 'Hello Budi!' ")
}
