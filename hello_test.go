package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	emptyRes := Hello("")

	expected := "Hello From: "
	if emptyRes != expected {
		t.Errorf("Unexpected emptyRes,expected %s, got %s", expected, emptyRes)
	}

	result := Hello("Mike")

	expected = "Hello From: Mike"
	if result != expected {
		t.Errorf("Unexpected emptyRes,expected %s, got %s", expected, emptyRes)
	}
}
