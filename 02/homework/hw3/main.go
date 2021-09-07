package main

import "fmt"

func main() {
	var inputSlice = []int32{1, 2, 5, 2, 6, 2, 5}

	if len(inputSlice) == 0 {
		fmt.Print("Invalid input array!")
	} else {
		fmt.Print("Max length element: ", removeDuplicates(inputSlice))
	}
}

// to find maximum length elements
func removeDuplicates(inputSlice []int32) (outputSlice []int32) {
	if len(inputSlice) == 1 {
		outputSlice = inputSlice
		return
	}

	var sliceDictionary = map[int32]int32{}
	for _, i := range inputSlice {
		_, ok := sliceDictionary[i]
		if ok {
			sliceDictionary[i] = sliceDictionary[i] + 1
		} else {
			sliceDictionary[i] = 1
		}
	}

	for k := range sliceDictionary {
		outputSlice = append(outputSlice, k)
	}

	return
}
