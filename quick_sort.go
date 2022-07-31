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

func main() {
	rand.Seed(time.Now().Unix())
	test := rand.Perm(30)
	fmt.Println(test)
	quickSort(test, 0, len(test)-1)
	fmt.Println(test)
}
