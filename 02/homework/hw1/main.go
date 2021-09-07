package main

import "fmt"

func main() {
	var inputSlice = []int32{7500000, 3650000, 3550000, 5250000, 7500000, 3550000}

	if len(inputSlice) < 2 {
		fmt.Print("Invalid input array!")
	} else {
		second := max2Numbers(inputSlice)
		fmt.Print("The second biggest element is: ", second)
	}
}

// to find the second biggest number
func max2Numbers(inputSlice []int32) (second int32) {
	second = -2147483648
	first := inputSlice[0]

	for i := 1; i < len(inputSlice); i++ {
		if inputSlice[i] > first {
			second = first
			first = inputSlice[i]
		} else if inputSlice[i] > second && inputSlice[i] != first {
			second = inputSlice[i]
		}
	}
	return second
}
