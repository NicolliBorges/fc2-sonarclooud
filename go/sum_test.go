package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Errorf("Expected sum(2, 3) to be 5, but got %d", result)
	}

}