package mergesort

import "sync"

func MergeSort(arr []int) (sortedArr []int) {

	l, r := 0, len(arr)-1
	mergeSort(&arr, l, r)
	return arr
}

func mergeSort(arr *[]int, l int, r int) {

	if l >= r {
		return
	}
	m := l + (r-l)/2

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		mergeSort(arr, l, m)
	}()

	go func() {
		defer wg.Done()
		mergeSort(arr, m+1, r)
	}()

	wg.Wait()
	merge(arr, l, r, m)
}

func merge(arr *[]int, l int, r int, m int) {

	var (
		i    = l                      // left part
		j    = m + 1                  // right part
		k    = l                      // arr index
		temp = make([]int, len(*arr)) // temp array
	)

	// add min elements
	for i <= m && j <= r {
		if (*arr)[i] < (*arr)[j] {
			temp[k] = (*arr)[i]
			k++
			i++
		} else {
			temp[k] = (*arr)[j]
			k++
			j++
		}
	}

	// add remaining elements from left_sub_array
	for i <= m {
		temp[k] = (*arr)[i]
		k++
		i++
	}

	// add remaining elements from right_sub_array
	for j <= r {
		temp[k] = (*arr)[j]
		k++
		j++
	}

	// copy elements
	for x := l; x <= r; x++ {
		(*arr)[x] = temp[x]
	}
}
