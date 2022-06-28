package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Contains returns true if the element exist in the slice, else it returns false
func Contains[T comparable](arr []T, item T) bool {
	for _, e := range arr {
		if e == item {
			return true
		}
	}
	return false
}

// IsSliceEqual checks if both slices are same or not
// NOTE: this function first sort both slices and then checks for the elements
// so if the elements are shuffle but same, then the function will return true
func IsSliceEqual[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	sort.SliceStable(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.SliceStable(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
