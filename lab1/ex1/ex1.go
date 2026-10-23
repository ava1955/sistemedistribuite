package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prim(x int) bool {
	d := 2
	for d <= x/2 {
		if x%d == 0 {
			return false
		}
		d++
	}
	return true
}
func main() {
	var ceva int
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		var err error
		ceva, err = strconv.Atoi(strings.TrimSpace(input))
		if err == nil {
			if prim(ceva) {
				fmt.Println("prim")
			} else {
				fmt.Println("nu e prim")
			}
		} else {
			fmt.Println("eroare")
		}
	}
}
