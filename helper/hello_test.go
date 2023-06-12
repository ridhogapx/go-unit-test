package helper

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Ridho")
	if result != "Hello Ridho" {
		panic("Result is not Hello Ridho")
	}

}
