package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("eroare")
		return
	}
	linie := scanner.Text()
	cuvinte := strings.Fields(linie)
	lungime := len(cuvinte)
	for i := 0; i < lungime/2; i++ {
		indexaux := lungime - 1 - i
		aux := cuvinte[indexaux]
		cuvinte[indexaux] = cuvinte[i]
		cuvinte[i] = aux
		fmt.Print(cuvinte[i], " ")
	}
	for i := lungime / 2; i < lungime; i++ {
		fmt.Print(cuvinte[i], " ")
	}
	fmt.Println()
}
