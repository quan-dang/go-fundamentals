package main

import (
	"fmt"
	"time"
)

func main() {
	// get current time
	dt := time.Now()
	fmt.Printf("   %s %d  \n", dt.Month(), dt.Year())

	dates := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	for _, date := range dates {
		fmt.Printf("%2s ", date)
	}
	fmt.Print("\n")

	for i := 0; i < 5; i++ {
		dates = append(dates, []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}...)
	}

	// get first day and last day of this month
	firstDay := time.Date(dt.Year(), dt.Month(), 1, 0, 0, 0, 0, dt.Location())
	lastDay := firstDay.AddDate(0, 1, -1)

	flag := true
	count := 1
	for _, date := range dates {
		if (date != firstDay.Weekday().String()[:2] && flag) || (count > int(lastDay.Day())) {
			fmt.Print("   ")
		} else if date == firstDay.Weekday().String()[:2] && flag {
			fmt.Printf("%2d ", count)
			count = count + 1
			flag = false
		} else if date == lastDay.Weekday().String()[:2] && count == int(lastDay.Day()) {
			fmt.Printf("%2d ", count)
			count = count + 1
		} else {
			fmt.Printf("%2d ", count)
			count = count + 1
		}

		if date == "Sa" {
			fmt.Print("\n")
		}
	}
}
