package rotationminnum11

func minNumberInRotateArray(array []int) int {
	low, high := 0, len(array)-1

	//one exception,rotate 0 elements, the array-self
	if array[low] < array[high] {
		return array[low]
	}

	mid := (low + high) / 2
	for array[low] >= array[high] {
		if high-low == 1 {
			mid = high
			break
		}
		mid = (low + high) / 2
		//if cannot confirm the middle elements are belong to the front or end
		//only search by order
		if array[low] == array[mid] && array[mid] == array[high] {
			return MinOrder(array, low, high)
		}

		if array[mid] >= array[low] {
			low = mid
		} else {
			high = mid
		}
	}
	return array[mid]

}

//search by order
func MinOrder(array []int, low, high int) int {
	min := array[low]
	for i := low; i < high; i++ {
		if array[i] < min {
			min = array[i]
		}
	}
	return min
}
