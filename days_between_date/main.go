package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	date1d := IndDate(1991, 12, 21)
	date2d := IndDate(2022, 5, 23)

	days, err := daysBetweenDateDirect(date1d, date2d)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(days)

	// ^^ solution 1 ======================================
	// ======================================
	// vv solution 2 ====================================

	date1 := IndoDate{
		Day:   21,
		Month: 12,
		Year:  1991,
	}

	date2 := IndoDate{
		Day:   23,
		Month: 05,
		Year:  2022,
	}

	days, err = daysBetweenDate(date1, date2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(days)

}

// Solution 1==============================
func IndDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func daysBetweenDateDirect(from, to time.Time) (int, error) {

	// validate from date is before to date
	if !from.Before(to) {
		if from.Equal(to) {
			return 0, nil
		}
		return 0, errors.New("first date must be older than second date")
	}
	days := to.Sub(from).Hours() / 24

	return int(days), nil
}

// 2 Solution (for decimal int with zero train)==========================
type IndoDate struct {
	Day   int
	Month int
	Year  int
}

func (i IndoDate) ToDate() (time.Time, error) {
	format := "2006-01-02" // YYYY-MM-DD
	return time.Parse(format, fmt.Sprintf("%04d-%02d-%02d", i.Year, i.Month, i.Day))
}

func daysBetweenDate(from, to IndoDate) (int, error) {
	// parse input to realdate golang
	fromDate, err := from.ToDate()
	if err != nil {
		return 0, err
	}
	toDate, err := to.ToDate()
	if err != nil {
		return 0, err
	}

	// validate from date is before to date
	if !fromDate.Before(toDate) {
		if fromDate.Equal(toDate) {
			return 0, nil
		}
		return 0, errors.New("first date must be older than second date")
	}
	days := toDate.Sub(fromDate).Hours() / 24

	return int(days), nil
}
