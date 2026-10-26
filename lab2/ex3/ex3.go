package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "primul mesaj"
		ch <- "al doilea mesaj"
	}()
	msg1 := <-ch
	msg2 := <-ch
	fmt.Println(msg1)
	fmt.Println(msg2)
}
