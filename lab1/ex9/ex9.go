package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())
	if !scanner.Scan() {
		return
	}
	s, _ := strconv.Atoi(scanner.Text())

	vizitate := make(map[int]int)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		v, _ := strconv.Atoi(scanner.Text())
		complement := s - v
		if index, gasit := vizitate[complement]; gasit {
			fmt.Println(i, " ", index)
		}
		vizitate[v] = i
	}
}
