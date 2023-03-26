package main

import (
	"fmt"

	"github.com/apetsko/course/romanConvert"
	"github.com/apetsko/course/sorts"
	"github.com/apetsko/course/twoDslice"
)

func main() {
	unsorted := []int{92, 22, 81, 67, 87, 65, 49, 25, 20, 6, 72, 70, 47, 26, 45, 40, 100, 73, 95, 19, 41, 1, 59, 38, 71, 50, 99, 84, 42, 80, 52, 54, 79, 18, 48, 85, 9, 14, 28, 55, 7, 75, 46, 27, 90, 43, 64, 29, 2, 53, 10, 66, 56, 74, 93, 11, 39, 97, 94, 62, 17, 61, 23, 63, 98, 91, 88, 68, 3, 16, 32, 44, 24, 34, 5, 89, 96, 86, 76, 82, 58, 13, 15, 21, 31, 83, 35, 51, 30, 60, 69, 57, 37, 33, 78, 12, 8, 4, 36, 77}
	tosort := make([]int, len(unsorted))

	//BubbleSort
	copy(tosort, unsorted)
	fmt.Println(sorts.BubbleSort(tosort))

	//MergeSort
	copy(tosort, unsorted)
	fmt.Println(sorts.MergeSort(tosort))

	//InsertionSort
	copy(tosort, unsorted)
	fmt.Println(sorts.InsertionSort(tosort))

	//random 2d slice generator & find max slice
	randArr := twoDslice.GenArray(5, 3, 100)
	fmt.Println(randArr)
	fmt.Println(twoDslice.GetMax(randArr))

	//convert to roman
	fmt.Println(romanConvert.Arab2roman(33333))
}
