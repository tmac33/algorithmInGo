package main

import (
	"math"
)

func main() {

}

func JumpSteps(n int) int {
	if n <= 2 {
		return n
	}
	pre1, pre2 := 1, 2
	res := 1
	for i := 2; i < n; i++ {
		res = pre1 + pre2
		pre1, pre2 = res, pre1
	}
	return res
}

//jump2 jump1,2,...,n
func jump2(target int) int {
	return int(math.Pow(2, float64(target-1)))
}
