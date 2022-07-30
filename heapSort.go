package main

// import "fmt"

func findLargest(arr []int) (int, int) {
	
	largest := arr[0]
	largestIndex := 0

	for index, value := range arr {
		if largest < value {
			largest = value
			largestIndex = index
		}
	}
	return largest, largestIndex
}

func heapify(arr *[]int) []int {

	// Arbitrairly large number
	var _, largestIndex int = findLargest(*arr)

	// I decided to merge three steps:
	// 1. rootIndex := arr[0]
	// 2. swap arr[rootIndex] with arr[largestIndex]
	// 3. swap arr[rootIndex] with arr[len(arr) - 1]
	// into two steps step:
	rootIndex := len(*arr) - 1
	(*arr)[rootIndex], (*arr)[largestIndex] = (*arr)[largestIndex], (*arr)[rootIndex]

	return (*arr)[0:len((*arr)) - 1]
}

func heapSort(arr []int) []int {

	new_array := arr

	for {
		new_array = heapify(&new_array)
		// fmt.Println(arr, new_array)

		if len(new_array) <= 1 {
			break
		}
	}
	return arr
}