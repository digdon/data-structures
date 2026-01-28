package quicksort

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty array", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"unsorted", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
		{"duplicates", []int{3, 3, 3, 3}, []int{3, 3, 3, 3}},
		{"negative numbers", []int{-3, 1, -5, 4, 0}, []int{-5, -3, 0, 1, 4}},
		{"two elements", []int{2, 1}, []int{1, 2}},
		{"large array", []int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := quicksort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("quicksort() = %v, want %v", result, tt.expected)
			}
		})
	}
}
