package utils

import "testing"

func TestContains(t *testing.T) {
	arr1 := []string{"one", "two", "three"}
	if !Contains(arr1, "two") {
		t.Fatalf("expcted two in the list: %v, but not found", arr1)
	}
	if Contains(arr1, "five") {
		t.Fatalf("not expected five in the list: %v, but returns true", arr1)
	}

	arr2 := []int{1, 2, 3}
	if !Contains(arr2, 2) {
		t.Fatalf("expcted 2 in the list: %v, but not found", arr2)
	}
	if Contains(arr2, 4) {
		t.Fatalf("not expected 5 in the list: %v, but returns true", arr2)
	}

}

func TestIsSliceEqual(t *testing.T) {
	if !IsSliceEqual([]string{"one", "two", "three"}, []string{"one", "two", "three"}) {
		t.Fatal("expected slices to be equal, but got false")
	}

	if IsSliceEqual([]string{"one", "two", "three"}, []string{"one", "two", "thre"}) {
		t.Fatal("expected slices to be not equal, but got true")
	}

	if IsSliceEqual([]string{"one", "two", "three"}, []string{"one", "two", "three", "foure"}) {
		t.Fatal("expected slices to be not equal, but got true")
	}

	if !IsSliceEqual([]int{1, 2, 3}, []int{3, 2, 1}) {
		t.Fatal("expected slices to be equal, but got false")
	}
}
