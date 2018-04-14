package main

import (
	"fmt"
	"time"
)

func main() {
	var listOfIntegers = []int{3, 9, 4, 2, 5, 6, 8, 32, 65, 76}
	searchNumber := 33
	fmt.Println(isPresent(searchNumber, listOfIntegers))
}

//Even if the listOfIntegers is nil we return -1 index
func isPresent(num int, listOfInteger []int) bool {
	sort(listOfInteger)
	fmt.Println(listOfInteger)
	return binarySearch(listOfInteger, num, 0, len(listOfInteger))

}

func binarySearch(listOfInteger []int, num, min, max int) bool {
	mid := min + (max-min)/2
	fmt.Println("mid :", mid, " max : ", max, " min : ", min)
	time.Sleep(2 * time.Second)
	if listOfInteger[mid] == num {
		return true
	}
	if max <= min {
		return false
	}
	if listOfInteger[mid] > num {
		return binarySearch(listOfInteger, num, min, mid-1)
	}
	return binarySearch(listOfInteger, num, mid, max)
}

//Below is selections sort here
func sort(listOfInteger []int) {
	lengthOfInteger := len(listOfInteger)
	j := lengthOfInteger
	for k := 0; k < lengthOfInteger-1; k++ {
		for i := 0; i < j-1; i++ {
			if listOfInteger[i] > listOfInteger[i+1] {
				z := listOfInteger[i]
				listOfInteger[i] = listOfInteger[i+1]
				listOfInteger[i+1] = z
			}
		}
		j = j - 1
	}
}
