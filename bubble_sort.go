package main

import (
	"fmt"
	"math/rand"
	"time"
)

//冒泡排序
func BubbleSort(array []int) {
	for i := 0; j < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	test := rand.Perm(10)
	fmt.Println(test)
	BubbleSort(test)
	fmt.Println(test)
}
