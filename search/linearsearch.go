package main

import "fmt"

func main() {
	var listOfIntegers = []int{3, 9, 4, 2, 4, 5, 6, 8, 2, 32, 54, 65, 76}
	searchNumber := 650
	fmt.Println(getIndex(searchNumber, listOfIntegers))
}

func getIndex(num int, listOfInteger []int) int {
	for index, intInList := range listOfInteger {
		if intInList == num {
			return index
		}
	}
	return -1
}
