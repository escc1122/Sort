package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		insertionValue := array[i]
		insertionIndex := i - 1
		for insertionIndex >= 0 && array[insertionIndex] > insertionValue {
			array[insertionIndex+1] = array[insertionIndex]
			insertionIndex--
		}
		array[insertionIndex+1] = insertionValue
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	test := rand.Perm(10)
	fmt.Println(test)
	InsertionSort(test)
	fmt.Println(test)
}
