package main

func main() {

}

func RectCover(n int) int {
	if n <= 2 {
		return n
	}
	pre1, pre2 := 2, 1
	result := 0
	for i := 3; i <= n; i++ {
		result = pre1 + pre2
		pre1, pre2 = result, pre1
	}
	return result
}
