package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Person struct {
	ten       string
	heSoLuong float32
	troCap    float32
}

func main() {
	var danhSach = []Person{
		{"Ti", 1.5, 3000000.0},
		{"Leu", 1.5, 3000000.0},
		{"Van", 1.7, 1000000.0},
		{"Quan", 3.0, 3000000.0},
		{"Le", 1.1, 2000000.0},
		{"Tung", 1.5, 3000000.0},
		{"Hai", 3.0, 3000000.0},
		{"Duy", 1.7, 1000000.0},
		{"To", 1.5, 3000000.0},
	}

	var (
		err     error
		feature int64
	)

	fmt.Println("Vui long chon chuc nang tuong ung")
	fmt.Println("1. Sap xep nhan vien tang dan theo bang chu cai")
	fmt.Println("2. Sap xep nhan vien muc luong giam dan")
	fmt.Println("3. Danh sach nhan vien co muc luong lon thu 2")
	// input a, b, c from keyboard
	feature, err = readNumberFromKeyboard("Nhap chuc nang: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	switch feature {
	case 1:
		fmt.Println("Ket qua: ", sortByLetter(danhSach))
	case 2:
		fmt.Println("Ket qua: ", sortBySalary(danhSach))
	case 3:
		fmt.Println("Ket qua: ", max2Salary(danhSach))
	default:
		fmt.Println("Vui long nhap chu nang 1, 2, hoac 3.")
	}
}

// sort danhSach by letters
func sortByLetter(danhSach []Person) []Person {
	sort.SliceStable(danhSach, func(i int, j int) bool {
		return danhSach[i].ten < danhSach[j].ten
	})
	return danhSach
}

// calculate salary
func calculateSalary(danhSach Person) float32 {
	return danhSach.heSoLuong*1500000 + danhSach.troCap
}

// sort danhSach by salary
func sortBySalary(danhSach []Person) []Person {
	sort.SliceStable(danhSach, func(i int, j int) bool {
		return calculateSalary(danhSach[i]) > calculateSalary(danhSach[j])
	})
	return danhSach
}

// find person with the second biggest salary
func max2Salary(danhSach []Person) (ketQua []Person) {
	var second float32 = -1.0
	first := calculateSalary(danhSach[0])

	// find second biggest salary
	for i := 1; i < len(danhSach); i++ {
		currentSalary := calculateSalary(danhSach[i])
		if currentSalary > first {
			second = first
			first = currentSalary
		} else if currentSalary > second && currentSalary != first {
			second = currentSalary
		}
	}

	// to find person with equal salary to second
	for i := 0; i < len(danhSach); i++ {
		if calculateSalary(danhSach[i]) == second {
			ketQua = append(ketQua, danhSach[i])
		}
	}

	return
}

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

// read from keyboard and convert string to float
func readNumberFromKeyboard(msg string) (result int64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")

	if result, err = strconv.ParseInt(str, 10, 64); err != nil {
		return 0.0, err
	}
	return result, nil
}
