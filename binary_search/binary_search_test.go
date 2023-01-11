package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {

	testCases := []struct {
		name     string
		arr      []int
		x        int
		expected int
		ok       bool
	}{
		{
			name:     "Test Case 1: Positive Test Case",
			arr:      []int{1, 2, 3, 4, 5},
			x:        3,
			expected: 2,
			ok:       true,
		},
		{
			name:     "Test Case 2: Negative Test Case - element not in array",
			arr:      []int{1, 2, 3, 4, 5},
			x:        8,
			expected: 5,
			ok:       false,
		},
		{
			name:     "Test Case 3: Empty Array",
			arr:      []int{},
			x:        3,
			expected: 0,
			ok:       false,
		},
		{
			name:     "Test Case 4: Array of length 1",
			arr:      []int{5},
			x:        5,
			expected: 0,
			ok:       true,
		},
		{
			name:     "Test Case 5: Positive Test Case",
			arr:      []int{1, 2, 3, 4, 5},
			x:        7,
			expected: 5,
			ok:       false,
		},
		{
			name:     "Test Case 6: Large array",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			x:        15,
			expected: 14,
			ok:       true,
		},
		{
			name:     "Test Case 7: Large array with dublicat element",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			x:        15,
			expected: 14,
			ok:       true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			index, ok := BinarySearch(test.arr, test.x)
			if index != test.expected || ok != test.ok {
				t.Errorf("Test case failed: expected index=%d, ok=%t but got index=%d, ok=%t", test.expected, test.ok, index, ok)
			}
		})
	}
}

func TestBinarySearchUpperBound(t *testing.T) {

	testCases := []struct {
		name     string
		arr      []int
		x        int
		expected int
		ok       bool
	}{
		{
			name:     "Test Case 1: Positive Test Case",
			arr:      []int{1, 2, 3, 4, 5},
			x:        3,
			expected: 2,
			ok:       true,
		},
		{
			name:     "Test Case 2: Negative Test Case - element not in array",
			arr:      []int{1, 2, 3, 4, 5},
			x:        8,
			expected: 4,
			ok:       false,
		},
		{
			name:     "Test Case 3: Empty Array",
			arr:      []int{},
			x:        3,
			expected: -1,
			ok:       false,
		},
		{
			name:     "Test Case 4: Array of length 1",
			arr:      []int{5},
			x:        5,
			expected: 0,
			ok:       true,
		},
		{
			name:     "Test Case 5: Large array",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			x:        15,
			expected: 14,
			ok:       true,
		},
		{
			name:     "Test Case 6: Possitive Test Case with dublicates elements",
			arr:      []int{1, 2, 3, 3, 4, 4, 4, 5},
			x:        4,
			expected: 6,
			ok:       true,
		},
		{
			name:     "Test Case 7: Negative Test Case with dublicates elements",
			arr:      []int{1, 2, 3, 3, 4, 4, 4, 5},
			x:        7,
			expected: 7,
			ok:       false,
		},
		{
			name:     "Test Case 8: Positive Test Case search element located in last index.",
			arr:      []int{1, 2, 3, 3, 4, 4, 4, 5, 5, 5},
			x:        5,
			expected: 9,
			ok:       true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			index, ok := BinarySearchUpperBound(test.arr, test.x)
			if index != test.expected || ok != test.ok {
				t.Errorf("Test case failed: expected index=%d, ok=%t but got index=%d, ok=%t", test.expected, test.ok, index, ok)
			}
		})
	}
}
