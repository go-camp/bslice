package bslice

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestBatch(t *testing.T) {
	var arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var arr3 []int
	for bsize := -20; bsize < 20; bsize++ {
		arr3 = arr3[:0]
		Batch(len(arr1), bsize, func(i, j int) bool {
			arr3 = append(arr3, arr1[i:j]...)
			return false
		})
		sort.Ints(arr3)
		if !reflect.DeepEqual(arr1, arr3) {
			t.Errorf("Batch(%d, %d, f) want %v but get %v", len(arr1), bsize, arr1, arr3)
		}
	}
}

func ExampleBatch() {
	var s = []int{1, 2, 3, 4, 5}
	Batch(len(s), 2, func(i, j int) bool {
		fmt.Println(s[i:j])
		return false
	})
	// Output:
	// [1 2]
	// [3 4]
	// [5]
}

func ExampleBatch_reverse() {
	var s = []int{1, 2, 3, 4, 5}
	Batch(len(s), -2, func(i, j int) bool {
		fmt.Println(s[i:j])
		return false
	})
	// Output:
	// [4 5]
	// [2 3]
	// [1]
}
