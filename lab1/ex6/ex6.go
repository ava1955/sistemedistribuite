package main

import (
	"fmt"
)

func main() {
	unmap := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}
	var nume []string
	var inaltimi []int

	for n, i := range unmap {
		nume = append(nume, n)
		inaltimi = append(inaltimi, i)
	}
	lungime := len(inaltimi)
	for i := 0; i < lungime-1; i++ {
		for j := 0; j < lungime-i-1; j++ {
			if inaltimi[j] < inaltimi[j+1] {
				auxInaltime := inaltimi[j]
				inaltimi[j] = inaltimi[j+1]
				inaltimi[j+1] = auxInaltime

				auxNume := nume[j]
				nume[j] = nume[j+1]
				nume[j+1] = auxNume
			}
		}
	}

	for i := 0; i < len(nume); i++ {
		fmt.Println(nume[i])
	}
}
