package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numere := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		ceva, err := strconv.Atoi(strings.TrimSpace(input))
		if err == nil {
			numere = append(numere, ceva)
		} else {
			fmt.Println("eroare")
		}
	}
	for i := range numere {
		if numere[i]%2 == 0 {
			fmt.Println(numere[i])
		}
	}

}
