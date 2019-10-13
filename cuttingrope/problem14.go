package cuttingrope

import (
	"math"
)

//DP
func maxProductAfterCutting(length int) (max int) {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}

	products := make([]int, length+1)

	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	for i := 4; i <= length; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if max < product {
				max = product
			}
			products[i] = max
		}
	}
	max = products[length]
	return
}

//greedy algorithm
func maxProductAfterCutting2(length int) (max int) {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}

	//cut 3m rope as many as possible
	timeOf3 := length / 3

	//in the end when the length is 4 ,should divide into 2 parts 2*2
	if length-timeOf3*3 == 1 {
		timeOf3 -= 1
	}

	timeOf2 := (length - timeOf3*3) / 2
	max1 := math.Pow(3, timeOf3)
	max2 := math.Pow(2, timeOf2)
	max = int(max1)
	return
}
