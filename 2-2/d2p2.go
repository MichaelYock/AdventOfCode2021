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
	aim := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		n := scanner.Text()
		command := strings.Split(n, " ")
		direction := command[0]
		distance, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "up":
			aim -= distance
		case "down":
			aim += distance
		case "forward":
			x += distance
			y += (aim * distance)
		}
	}
	input.Close()
	fmt.Println(x * y)
}
