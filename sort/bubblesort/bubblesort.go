package main

import "fmt"

func main() {
	//a := []int{9, 3, 4, 1, 98, 43, 54, 87, 34}

	a := []int{1, 4, 6, 8, 7, 12, 34, 56, 7, 0, -10}
	bubbleSort(a)
}

func bubbleSort(array []int) {
	swapped := true
	for i := 0; i < len(array)-1; i++ {
		if swapped {
			swapped = false
			for j := 0; j < len(array)-1-i; j++ {
				if array[j] > array[j+1] {
					temp := array[j]
					array[j] = array[j+1]
					array[j+1] = temp
					swapped = true
				}
			}
		} else {
			break
		}
	}
	fmt.Println(array)

}
