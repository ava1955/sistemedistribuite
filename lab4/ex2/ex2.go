package main

import (
	"fmt"
	"sync"
)

func Node(id int, in chan int, out chan int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	max := id
	diametru := n - 1
	for k := 0; k < diametru; k++ {
		out <- max
		val := <-in
		if val > max {
			max = val
		}
	}

	if max == id {
		fmt.Println("Nodul", id, "e lider!")
	}
}

func main() {
	n := 3
	channels := make([]chan int, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan int, 1)
	}

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		nextIndex := (i + 1) % n
		go Node(i, channels[i], channels[nextIndex], n, &wg)
	}
	wg.Wait()
}
