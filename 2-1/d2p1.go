package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := os.Args[1]
	input, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	x := 0
	y := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		n := scanner.Text()
		command := strings.Split(n, " ")

		distance, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		switch command[0] {
		case "up":
			y -= distance
		case "down":
			y += distance
		case "forward":
			x += distance
		}

	}
	input.Close()
	fmt.Println("Co-ordinates:", x*y)
}
