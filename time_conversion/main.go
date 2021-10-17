package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	timeInput := "07:05:45PM"
	timeOutput := timeConversion(timeInput)

	fmt.Println(timeOutput)
}

func timeConversion(s string) string {
	// 07:05:45PM
	// 19:05:45
	// pecah jadi 3 bagian
	sSlice := strings.Split(s, ":")
	if len(sSlice) != 3 {
		return "time not valid"
	}
	hour, _ := strconv.Atoi(sSlice[0])
	minute, _ := strconv.Atoi(sSlice[1])
	second, _ := strconv.Atoi(sSlice[2][:2])
	amPm := sSlice[2][2:]

	if amPm == "PM" {
		if hour != 12 {
			hour = hour + 12
		}
	} else {
		// is AM
		// if 12 AM == 00
		if hour == 12 {
			hour = 0
		}
	}

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}
