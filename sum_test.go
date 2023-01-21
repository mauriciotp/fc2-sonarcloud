package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}

	result2 := sub(4, 2)

	if result != 2 {
		t.Error("The result must be 2")
	}
}
