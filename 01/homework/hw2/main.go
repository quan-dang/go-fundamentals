package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		err          error
		x            int
		randomNumber int
	)

	// generate the random number to guess
	randomNumber = rand.Intn(101)

	for x != -1 {
		// input x from keyboard
		x, err = readNumberFromKeyboard("Nhap x: ")
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		if x == randomNumber {
			fmt.Println("Ban da doan dung!")
			break
		} else if x < randomNumber {
			fmt.Println("So ban doan nho hon X!")
		} else {
			fmt.Println("So ban doan lon hon X!")
		}
	}
}

// read from keyboard and convert string to int
func readNumberFromKeyboard(msg string) (result int, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")

	if result, err = strconv.Atoi(str); err != nil {
		return 0.0, err
	}
	return result, nil
}
