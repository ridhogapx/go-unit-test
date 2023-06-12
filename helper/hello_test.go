package helper

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Ridho")
	if result != "Hello dude" {
		t.Fail()
	}
}

func TestGalih(t *testing.T) {
	result := Hello("Galih")
	if result != "Hello Galih" {
		t.Fail()
	}
}

func TestBudi(t *testing.T) {
	result := Hello("Budi")
	if result != "Hello Budi" {
		t.Fail()
	}
}
