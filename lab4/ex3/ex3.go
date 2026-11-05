package main

import (
	"fmt"
	"sync"
)

func Node(id int, in chan float64, out chan float64, initialValue float64, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	x := initialValue
	di := 1
	rounds := n * 5

	for t := 0; t < rounds; t++ {
		out <- x
		neighborVal := <-in
		x = (1.0 / float64(di+1)) * (x + neighborVal)
	}

	var votFinal int
	if x > 0.5 {
		votFinal = 1
	} else {
		votFinal = 0
	}

	fmt.Println("Nodul", id, "are valoarea", x, ", vot majoritar:", votFinal)
}

func main() {
	n := 4
	initialValues := []float64{1.0, 1.0, 0.0, 1.0}

	channels := make([]chan float64, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan float64, 1)
	}

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		nextIndex := (i + 1) % n
		go Node(i, channels[i], channels[nextIndex], initialValues[i], n, &wg)
	}

	wg.Wait()
}
