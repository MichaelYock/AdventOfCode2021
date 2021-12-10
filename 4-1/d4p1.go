package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	file := os.Args[1]
	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	draw := strings.Split(scanner.Text(), ",")
	scanner.Scan()

	var board [][]string
	var winningestTurn int

	for scanner.Scan() {
		if scanner.Text() != "" {
			row := strings.Fields(scanner.Text())
			board = append(board, row)
		} else {
			winTurn := checkBoard(board)

			// do soething with winTurn
			if winTurn > winningestTurn {
				winningestTurn = winTurn
			}
		}
	}

	fmt.Println(draw)
	fmt.Println(winningestTurn)
}

func checkBoard(b [][]string) int {
	// find when board wins
	return rand.Intn(100)
}
