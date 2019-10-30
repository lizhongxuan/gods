package main

import (
	"fmt"
	"time"
)

func main() {
	FromTime := "2019-10-23"
	ToTime := "2019-10-25"
	days := getEveryDay(FromTime, ToTime)
	fmt.Println(days)
}

func getEveryDay(FromTime, ToTime string) []string {
	FromTime = FromTime + " 00:00:00"
	timeTemplate := "2006-01-02 15:04:05"
	from_time, err := time.ParseInLocation(timeTemplate, FromTime, time.Local)
	if err != nil {
		fmt.Println("from_time error:", err)
		return nil
	}

	var to_time time.Time
	if ToTime != "" {
		ToTime = ToTime + " 23:59:59"
		to_time, err = time.ParseInLocation(timeTemplate, ToTime, time.Local)
		if err != nil {
			fmt.Println("to_time error:", err)
			return nil
		}
	} else {
		to_time = time.Now()
	}

	if to_time.Before(from_time) {
		fmt.Println("error: ToTime before FromTime")
		return nil
	}

	var days []string
	for to_time.After(from_time) {
		days = append(days, from_time.Format("2006-01-02"))
		from_time = from_time.AddDate(0, 0, 1)
	}
	return days
}
