package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{8, 7, 6, 5, 4, 3, 2, 1}

	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Intn(30))
	}

	fmt.Printf("Array at start: %v\n", arr)
	fmt.Printf("Array sorted  : %v\n", heapSort(arr))
}