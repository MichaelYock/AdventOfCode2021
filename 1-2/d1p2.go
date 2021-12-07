package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile := os.Args[1]
	input, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input)
	count := 0

	window := []int{}

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		window = append(window, n)

		if len(window) > 3 {
			if sum(window)-window[0] > sum(window)-window[3] {
				count++
			}
			window = window[1:]
		}

	}
	input.Close()

	fmt.Println("Number of increases:", count)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
