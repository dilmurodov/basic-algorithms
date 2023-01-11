package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"Test 1: Positive Test case",
			[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5},
		},
		{
			"Test 2: Empty Input ",
			[]int{}, []int{},
		},
		{
			"Test 3: Already sorted ",
			[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5},
		},
		{
			"Test 4: Large Input ",
			[]int{12, 3, 2, 7, 5, 1, 4, 9, 11, 20, 15, 8, 6},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 20}},
		{
			"Test 5: Negative Numbers",
			[]int{5, -4, 3, 2, -1}, []int{-4, -1, 2, 3, 5},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			output := MergeSort(test.input)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Test case failed: expected %v, got %v", test.expected, output)
			}
		})
	}
}
