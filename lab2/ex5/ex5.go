package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "mesaj de la send"
	}()
	go func() {
		mesaj := <-ch1
		ch2 <- mesaj
	}()
	rezultat := <-ch2
	fmt.Println(rezultat)
}
