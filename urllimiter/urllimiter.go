package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	list := []string{"raj", "sujit", "vinod", "sanjay", "sdutta", "azeem", "jogi", "nishal", "123"}
	i := 0
	chunk := 2
	j := chunk
	for i < len(list) {
		var wg sync.WaitGroup
		for i < j && i < len(list) {
			wg.Add(1)
			go callRoutine(list[i], &wg)
			i++
		}
		time.Sleep(2 * time.Second)
		j = j + chunk
		wg.Wait()
	}
}

func callRoutine(s string, wg *sync.WaitGroup) {
	fmt.Println("hello ", s)
	wg.Done()
}
