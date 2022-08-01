package main

import (
	"fmt"
	"math/rand"
	"time"
)

//先執行一次partition分成左右邊之後 左右邊遞迴
func quickSort(array []int, left int, right int) {
	if left < right {
		newPivot := partition(array, left, right)
		quickSort(array, left, newPivot-1)
		quickSort(array, newPivot+1, right)
	}
}

//選擇最右邊為新的中心點 然後比他小的放左邊 大的放右邊
func partition(array []int, left int, right int) int {
	pivotValue := array[right]
	newPivot := left

	for i := left; i < right; i++ {
		if array[i] < pivotValue {
			array[i], array[newPivot] = array[newPivot], array[i]
			newPivot++
		}
	}
	array[right], array[newPivot] = array[newPivot], array[right]
	return newPivot
}

//https://www.techiedelight.com/zh-tw/quick-sort-using-hoares-partitioning-scheme/
//https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/645577/golang-with-lomutohoare-partition-scheme
//Hoare
func quickSortHoare(array []int, left int, right int) {
	if left < right {
		newPivot := partitionHoare(array, left, right)
		quickSortHoare(array, left, newPivot-1)
		quickSortHoare(array, newPivot+1, right)
	}
}

func partitionHoare(array []int, left int, right int) int {
	pivotValue := array[left]
	i, j := left, right
	for {
		for array[i] < pivotValue {
			i++
		}

		for array[j] > pivotValue {
			j--
		}

		if i >= j {
			return j
		}
		swap(array, i, j)
		i++
		j--
	}
}

func swap(array []int, a int, b int) {
	array[a], array[b] = array[b], array[a]
}

func main() {
	rand.Seed(time.Now().Unix())
	test := rand.Perm(30)
	testHoare := make([]int, len(test))
	copy(testHoare, test)
	fmt.Println(test)

	quickSort(test, 0, len(test)-1)
	fmt.Println(test)

	quickSortHoare(testHoare, 0, len(testHoare)-1)
	fmt.Println(testHoare)

}
