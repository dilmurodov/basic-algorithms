package main

import (
	"fmt"

	mergesort "github.com/dilmurodov/basic-algorithms/merge_sort"
)

func main() {
	arr := mergesort.MergeSort([]int{5, -4, 3, 2, -1})

	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
