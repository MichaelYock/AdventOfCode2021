package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile := os.Args[1] // [0] is program name
	input, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input) // split defaults to ScanLines, which splits at new line and trims whitespace
	count := 0
	p := 0 // this needs to be the first line from scanner... dirty solution, minus 1 from final result

	for scanner.Scan() { // Moves through scanner
		c, err := strconv.Atoi(scanner.Text()) // scanner.Text returns string, convert to int with str.conv.Atoi
		if err != nil {
			log.Fatal(err)
		}
		if c > p {
			count++
		}

		p = c
	}
	input.Close()

	fmt.Println("Number of increases:", count-1)
}
