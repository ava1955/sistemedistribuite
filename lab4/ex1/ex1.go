package main

import (
	"fmt"
	"sync"
)

func Node(id int, in chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	out <- id
	for val := range in {
		if val == -1 {
			out <- -1
			return
		}
		if val > id {
			out <- val
		} else if val == id {
			fmt.Println("Nodul", id, "devine LIDER!")
			out <- -1
			return
		}
	}
}

func main() {
	n := 4
	channels := make([]chan int, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan int, 1)
	}

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		nextIndex := (i + 1) % n
		go Node(i, channels[i], channels[nextIndex], &wg)
	}
	wg.Wait()
}
