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
	var bits [12]int
	var gamma string
	var epsilon string

	for scanner.Scan() {
		n := scanner.Text()
		for i := 0; i < len(n); i++ {
			if n[i] == 49 {
				bits[i]++
			} else {
				bits[i]--
			}
		}
	}

	for i := 0; i < len(bits); i++ {
		if bits[i] > 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(epsilonInt * gammaInt)
}
