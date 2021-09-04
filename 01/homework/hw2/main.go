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
		err error
		a   float64
		b   float64
		c   float64
	)

	// input a, b, c from keyboard
	a, err = readNumberFromKeyboard("a: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	b, _ = readNumberFromKeyboard("b: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	c, _ = readNumberFromKeyboard("c: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	// solve quadractic equation
	x1, x2 := solveQuadraticEquation(a, b, c)
	fmt.Println("x1: ", x1)
	fmt.Println("x2: ", x2)
}

// to solve quadratic equation
func solveQuadraticEquation(a, b, c float64) (x1, x2 interface{}) {
	delta := calculateDelta(a, b, c)
	normalD := math.Sqrt(delta)
	absoluteD := math.Sqrt(math.Abs(delta))

	if normalD > 0 {
		x1 = (-b + normalD) / (2 * a)
		x2 = (-b - normalD) / (2 * a)
	} else if absoluteD == 0 {
		x1 = (-b) / (2 * a)
		x2 = (-b) / (2 * a)
	} else {
		x1 = complex(-b, -absoluteD)
		x2 = complex(-b, absoluteD)
	}
	return
}

//  to calculate delta for the quadratic equation
func calculateDelta(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

// read from keyboard and convert string to float
func readNumberFromKeyboard(msg string) (result float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")

	if result, err = strconv.ParseFloat(str, 64); err != nil {
		return 0.0, err
	}
	return result, nil
}
