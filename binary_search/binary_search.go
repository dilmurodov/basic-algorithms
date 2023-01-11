package binarysearch

func BinarySearch(arr []int, x int) (index int, ok bool) {

	l, r := 0, len(arr)-1

	for l <= r {
		m := l + (r-l)/2

		if arr[m] < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if 0 <= l && l < len(arr) {
		return l, arr[l] == x
	}

	return l, false
}

func BinarySearchUpperBound(arr []int, x int) (index int, ok bool) {

	l, r := 0, len(arr)-1

	for l <= r {
		m := l + (r-l)/2
		if arr[m] <= x {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if 0 <= r && r < len(arr) {
		return r, arr[r] == x
	}

	return r, false
}
