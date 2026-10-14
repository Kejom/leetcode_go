package main

import "fmt"

func main() {
	var example1 = []int{1, 7, 3, 6, 5, 6}
	var example2 = []int{1, 2, 3}
	var example3 = []int{2, 1, -1}

	fmt.Println(findPivot(example1))
	fmt.Println(findPivot(example2))
	fmt.Println(findPivot(example3))
}

func findPivot(nums []int) int {

	var leftSum int = 0
	var rightSum int = 0

	for _, v := range nums {
		rightSum += v
	}

	for i, v := range nums {

		rightSum -= v
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}
	return -1
}
