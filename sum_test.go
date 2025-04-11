package main

import "testing"

// sum adds two integers and returns the result.
func TestSum(t *testing.T) {
	a := 1
	b := 2
	expected := 3
	result := sum(a, b)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}


func TestSub(t *testing.T) {
	a := 5
	b := 2
	expected := 3
	result := sub(a, b)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
func TestMul(t *testing.T) {
	a := 3
	b := 4
	expected := 12
	result := mul(a, b)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
func TestDiv(t *testing.T) {	
	a := 8
	b := 2
	expected := 4
	result := div(a, b)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
func TestDivNegative(t *testing.T) {
	a := -8
	b := 2
	expected := -4
	result := div(a, b)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}	