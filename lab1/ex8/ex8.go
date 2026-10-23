package main

import (
	"fmt"
)

func main() {
	unmap := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}
	unaltmap := make(map[int][]string)
	for nume, numar := range unmap {
		unaltmap[numar] = append(unaltmap[numar], nume)
	}
	for numar, nume := range unaltmap {
		fmt.Printf("Departamentul %d: [", numar)
		for i := range nume {
			if i == len(nume)-1 {
				fmt.Printf("%s", nume[i])
				break
			}
			fmt.Printf("%s, ", nume[i])
		}
		fmt.Println("]")
	}
}
