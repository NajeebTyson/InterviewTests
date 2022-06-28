package utils

import (
	"testing"
)

func TestGetMapValues(t *testing.T) {
	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"

	output := []string{"one", "two", "three"}

	values := GetMapValues(m)

	if !IsSliceEqual(output, values) {
		t.Errorf("Expected output: %v, got: %v", output, values)
	}

	m1 := make(map[string]int)
	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3

	output1 := []int{1, 2, 3}

	values1 := GetMapValues(m1)

	if !IsSliceEqual(output1, values1) {
		t.Errorf("Expected output: %v, got: %v", output1, values1)
	}
}
