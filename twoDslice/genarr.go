package twoDslice

import (
	"math/rand"
)

func GenArray(x, y, random int) [][]int {
	arr := make([][]int, x)
	for a := range arr {
		arr[a] = make([]int, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			arr[i][j] = rand.Intn(random)
		}
	}
	return arr
}

func GetMax(arr [][]int) int {
	var max int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if j == 0 || arr[i][j] > max {
				max = arr[i][j]
			}
		}
	}
	return max
}
