package algorithm

import (
	"fmt"
	"testing"
)

func quickSort(array []int, left, reight int) {
	fmt.Println(array, left, reight)
	if len(array) <= 1 {
		return
	}
	mid := array[left]
	min := left
	max := reight
	forward := false
	for min < max {
		if forward {
			if array[min] > mid {
				array[max] = array[min]
				max--
				forward = false
			} else {
				min++
			}
		} else {
			if array[max] < mid {
				array[min] = array[max]
				min++
				forward = true
			} else {
				max--
			}
		}
	}
	array[min] = mid
	if min > left+1 {
		quickSort(array, left, min)
	}
	if max < reight-1 {
		quickSort(array, max+1, reight)
	}
}
func TestQuicklSort(t *testing.T) {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
