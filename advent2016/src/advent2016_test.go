package main

import (
	"testing"
	"io/ioutil"
)

func TestDay1(t *testing.T) {
	result := Day1("R2, L3")
	if result != 5 {
		t.Errorf("Returned incorrect value! Expected %d got %d", 5, result)
	}

	result = Day1("R2, R2, R2")
	if result != 2 {
		t.Errorf("Returned incorrect value! Expected %d got %d", 2, result)
	}

	result = Day1("R5, L5, R5, R3")
	if result != 12 {
		t.Errorf("Returned incorrect value! Expected %d got %d", 12, result)
	}

	data, err := ioutil.ReadFile("../input/day1.txt")
    if err != nil {
    	t.Errorf("Unable to read file!")
    } else {
        result = Day1(string(data))
    }
}