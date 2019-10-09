package main

import (
	"fmt"
	"time"
)

func main() {

	var tim time.Time
	fmt.Println(tim)
	fmt.Println(tim.IsZero())
	tim = time.Now()
	fmt.Println(tim.Zone())
	fmt.Println(5*60*60 + 30*60)
	s := tim.Format(time.RFC3339)
	fmt.Println(tim)
	x := tim.UTC()
	fmt.Println(s)
	fmt.Println(x.Format(time.RFC3339))
}
