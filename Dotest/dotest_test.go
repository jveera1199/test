package main

import "testing"

func TestDivde(t *testing.T) {
	_, err := divide(100.0, 1.0)
	if err != nil {
		t.Error("Got an error when we sould not get")
	}
}
func TestBadDivide(t *testing.T) {
	_, err := divide(100.0, 0)
	if err == nil {
		t.Error("Did not get an error even I should get it")
	}
}
func TestAdd(t *testing.T) {
	_, err := add(10.0, 2.0)
	if err != nil {
		t.Error("returned error even if there is no error.")

	}
}
func TestBadAdd(t *testing.T) {
	_, err := add(0, 0)
	if err == nil {
		t.Error("did not get an error even if there is an error")
	}
}
