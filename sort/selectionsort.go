package main

import "fmt"

func main() {
	a := []int{9, 3, 4, 1, 98, 43, 54, 87, 34}
	fmt.Println("hello")
	sort(a)
}

func sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		max := 0
		for j := 0; j < len(array)-1-i; j++ {
			if array[max] < array[j+1] {
				max = j + 1
			}
		}
		temp := array[max]
		array[max] = array[len(array)-i-1]
		array[len(array)-i-1] = temp
	}
	fmt.Println(array)
}
