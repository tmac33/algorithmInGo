package rangeofrobot

import "fmt"

func movingCount(threshold, rows, cols int) int {
	bound := []int{rows, cols}
	//save all steps
	steps := [][]int{}

	//Store all the steps currently explored
	queue := [][]int{
		{0, 0},
	}

	//up,down,right,left
	dirs := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for len(queue) > 0 {
		//move one step
		current := queue[0]
		queue = queue[1:]
		if outOfBounds(current, bound) {
			continue
		}
		if biggerThanThreshold(current, threshold) {
			continue
		}
		if alreadyPassed(current, steps) {
			continue
		}
		steps = append(steps, current)
		for _, dir := range dirs {
			newStep := []int{current[0] + dir[0], current[1] + dir[1]}
			queue = append(queue, newStep)
		}
	}
	return len(steps)
}

func outOfBounds(step []int, bound []int) bool {
	if step[0] < 0 || step[1] < 0 || step[0] >= bound[0] || step[1] >= bound[1] {
		return true
	}
	return false
}

func biggerThanThreshold(step []int, threshold int) bool {
	i, j := step[0], step[1]
	k := 0
	for i > 0 {
		k += i % 10
		i /= 10
	}
	for j > 0 {
		k += j % 10
		j /= 10
	}
	if k > threshold {
		return true
	}

	return false
}

func alreadyPassed(step []int, steps [][]int) bool {
	for _, item := range steps {
		if step[0] == item[0] && step[1] == item[1] {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(movingCount(5, 10, 10), "[21]")
	fmt.Println(movingCount(15, 20, 20), "[359]")
	fmt.Println(movingCount(10, 1, 100), "[29]")
	fmt.Println(movingCount(10, 1, 10), "[10]")
	fmt.Println(movingCount(15, 100, 1), "[79]")
	fmt.Println(movingCount(15, 10, 1), "[10]")
	fmt.Println(movingCount(5, 10, 10), "[21]")
	fmt.Println(movingCount(12, 1, 1), "[1]")
	fmt.Println(movingCount(-10, 10, 10), "[0]")
}
