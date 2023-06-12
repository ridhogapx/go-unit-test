package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertion(t *testing.T) {
	result := Hello("RageNeko26")
	assert.Equal(t, "Hello RageNeko26", result, "Result must be 'Hello RageNeko26' ")
	fmt.Println("Testing...")
}
