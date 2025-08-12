package main

import (
	"flag"
	"fmt"
	"time"
)

func parseTime(target string) time.Time {
	parsedDate, err := time.Parse("2006/01/02", target)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsedDate)
	return parsedDate
}

func nightsUntil(target time.Time) int {
	return int(time.Until(target).Hours() / 24)
}

func Sleeps() {
	fmt.Println("Sleeps")
	bday := flag.String("bday", "", "Your Birthday in YYYY/MM/DD format. ex: 1999/01/01")

	flag.Parse()
	target := parseTime(*bday)

	fmt.Printf("nights until your birthday ar %d and today is %d", nightsUntil(target), target.Compare(time.Now()))
}
