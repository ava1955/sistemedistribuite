package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		mesaj := <-ch
		fmt.Println(mesaj)
	}()
	ch <- "salut"
}
