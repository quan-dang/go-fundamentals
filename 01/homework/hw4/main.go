package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		err        error
		n          int
		primeSlice []int
	)

	// input N from keyboard
	n, err = readNumberFromKeyboard("Nhap so tu nhien N: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	// exit program if N < 0
	if n < 0 || n > 100000 {
		fmt.Println("N khong hop le!")
		return
	}

	// loop over all natural number smaller than N+1
	for i := 0; i < n+1; i++ {
		if isPrime(i) {
			primeSlice = append(primeSlice, i)
		}
	}

	fmt.Println("Mang cac so nguyen to nho hon hoac bang N la: ", primeSlice)
}

// to check if a prime number
func isPrime(inputNumber int) (result bool) {
	prime_flag := 1

	if inputNumber > 1 {
		for i := 2; i < int(math.Sqrt(float64(inputNumber)))+1; i++ {
			if inputNumber%i == 0 {
				prime_flag = 0
				break
			}
		}
		if prime_flag == 0 {
			result = false
		} else {
			result = true
		}
	} else {
		result = false
	}
	return
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
