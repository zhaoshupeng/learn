package main

import "fmt"

func main() {
	// arr := []int{3, 8, 7, 5, 12}
	arr := []int{3, 38, 11, 659, 162} // 基数示例
	//BucketSortSecond(arr)
	// CountingSort(arr)
	RadixSort(arr)
	//BucketSort(arr, 3)
	//QuickSort(arr)
	//MergeSort(arr)
	fmt.Println("after sort", arr)
}
