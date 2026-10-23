package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("eroare")
		return
	}
	linie := scanner.Text()
	cuvinte := strings.Fields(linie)
	unmap := map[string]int{}
	for i := range cuvinte {
		unmap[cuvinte[i]] = unmap[cuvinte[i]] + 1
	}
	for cuvant, frec := range unmap {
		fmt.Print(cuvant, " ", frec, "\n")
	}
}
