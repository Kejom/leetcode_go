package main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {

	var numbersInSequence = make(map[int]struct{})

	var currentNumber = n
	numbersInSequence[currentNumber] = struct{}{}

	for {
		var newNumber = 0

		for currentNumber > 0 {
			digit := currentNumber % 10
			newNumber += digit * digit
			currentNumber /= 10
		}

		if newNumber == 1 {
			return true
		}

		if _, exists := numbersInSequence[newNumber]; exists {
			return false
		}

		numbersInSequence[newNumber] = struct{}{}
		currentNumber = newNumber

	}
}
