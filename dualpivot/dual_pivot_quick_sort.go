package main

import (
	"fmt"
	"math/rand"
	"time"
)

//雙軸快速排序
//同快速排序 只是分成三塊
func DualPivotQuickSort(array []int, left int, right int) {
	if left < right {
		newLeftPivot, newRightPivot := partition(array, left, right)
		DualPivotQuickSort(array, left, newLeftPivot-1)
		DualPivotQuickSort(array, newLeftPivot+1, newRightPivot-1)
		DualPivotQuickSort(array, newRightPivot+1, right)
	}
}

//選擇最右邊跟最左邊為新的中心點 分三塊整理
func partition(array []int, left int, right int) (int, int) {
	//選擇最右邊跟最左邊為新的中心點,並且讓左邊小於右邊
	if array[left] > array[right] {
		swap(array, right, left)
	}

	newLeftPivot := left + 1        //最左邊為新的中心點 並且先留下自己的位置
	newRightPivot := right - 1      //最右邊為新的中心點 並且先留下自己的位置
	leftPivotValue := array[left]   //新的中心點值
	rightPivotValue := array[right] //新的中心點值

	k := left //從左邊慢慢排序,直到跟右邊接觸
	for k <= newRightPivot {
		if array[k] < leftPivotValue { //同快速排序
			swap(array, k, newLeftPivot)
			newLeftPivot++
		} else if array[k] > rightPivotValue { //值需要分配到最右邊
			//最右邊新的中心點先向左邊移動,假設值比較小
			for array[newRightPivot] > rightPivotValue && k < newRightPivot {
				newRightPivot--
			}
			//移動完同快速排序 進行右邊
			swap(array, k, newRightPivot)
			newRightPivot--
			//如果出現值要分配到最左邊 就要再丟到最左邊
			if array[k] < leftPivotValue {
				swap(array, k, newLeftPivot)
				newLeftPivot++
			}
		}
		k++
	}
	//因為一開始有預留中心點位置 所以左右各退一步並交換
	newLeftPivot--
	newRightPivot++
	swap(array, right, newRightPivot)
	swap(array, left, newLeftPivot)

	return newLeftPivot, newRightPivot
}
func swap(array []int, a int, b int) {
	array[a], array[b] = array[b], array[a]
}

func main() {
	rand.Seed(time.Now().Unix())
	test := rand.Perm(30)
	fmt.Println(test)
	DualPivotQuickSort(test, 0, len(test)-1)
	fmt.Println(test)
}
