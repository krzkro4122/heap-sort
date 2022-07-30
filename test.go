package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{}
	// Populate the array with 20 ints from range [0, 10)
	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Intn(10))
	}

	fmt.Printf("Array at start: %v\n", arr)
	fmt.Printf("Array sorted  : %v\n", heapSort(arr))
}