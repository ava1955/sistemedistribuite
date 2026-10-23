package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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
	numere := make([]float32, 0, len(cuvinte))
	for _, cuvant := range cuvinte {
		val, err := strconv.ParseFloat(cuvant, 32)
		if err != nil {
			fmt.Println("eroare")
			return
		}
		numere = append(numere, float32(val))
	}
	if !scanner.Scan() {
		fmt.Println("eroare")
		return
	}
	input := scanner.Text()
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err == nil {
		if index >= 0 && index < len(numere) {
			numere = slices.Delete(numere, index, index+1)
		} else {
			fmt.Println("eroare")
			return
		}
		for i := range numere {
			fmt.Print(numere[i], " ")
		}
		fmt.Println()
	} else {
		fmt.Println("eroare")
		return
	}
}
