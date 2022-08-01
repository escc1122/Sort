package main

import (
	"fmt"
	"math/rand"
	"time"
)

//選擇排序法
func SelectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	test := rand.Perm(10)
	fmt.Println(test)
	SelectionSort(test)
	fmt.Println(test)
}
