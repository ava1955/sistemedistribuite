package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		fmt.Println("Procesare in desfasurare...")
		time.Sleep(500 * time.Millisecond)
		done <- true
	}()
	<-done
	fmt.Println("Functia main stie ca gorutina s-a terminat")
}
