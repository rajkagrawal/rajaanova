package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(getDate(5))
}

func getDate(minusDays int) int64 {
	timeAfter := time.Now().AddDate(0, 0, -minusDays)
	fmt.Println(timeAfter)
	newTie := time.Date(timeAfter.Year(), timeAfter.Month(), timeAfter.Day(), 0, 0, 0, 0, timeAfter.Location())
	fmt.Println(newTie)
	return newTie.Unix()
}
