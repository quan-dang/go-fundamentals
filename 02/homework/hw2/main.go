package main

import "fmt"

func main() {
	var inputSlice = []string{"aba", "aa", "ad", "c", "vcd"}

	if len(inputSlice) == 0 {
		fmt.Print("Invalid input array!")
	} else {
		fmt.Print("Max length element: ", findMaxLengthElement(inputSlice))
	}
}

// to find maximum length elements
func findMaxLengthElement(inputSlice []string) (outputSlice []string) {
	outputSlice = []string{inputSlice[0]}

	if len(inputSlice) == 1 {
		return
	}

	for i := 1; i < len(inputSlice); i++ {
		if len(inputSlice[i]) > len(outputSlice[0]) {
			outputSlice = []string{inputSlice[i]}
		} else if len(inputSlice[i]) == len(outputSlice[0]) {
			outputSlice = append(outputSlice, inputSlice[i])
		}
	}
	return
}
