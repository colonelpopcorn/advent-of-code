package main

import "testing"

func TestDay1(t *testing.T) {
	result := Day1("R2, L3")
	if result != 5 {
		t.Errorf("Returned incorrect value! Expected %d got %d", 5, result)
	}
}