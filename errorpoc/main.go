package main

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	s := getError()
	fmt.Println(s)
	logrus.Debug("hellor")
	errors.New("hello error new error")
}

type rajErr struct {
}

func (a rajErr) Error() string {
	fmt.Println("some good error")
	return ""
}

func getError() error {
	return rajErr{}
}
