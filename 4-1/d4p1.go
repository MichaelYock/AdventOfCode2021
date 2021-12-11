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
	boards, draw := setUpBoards()

	winTurn, board := checkBoards(boards, draw)

	fmt.Println(winTurn, board)
}

func setUpBoards() ([][][]int, []int) {
	file := os.Args[1]
	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	scanner.Scan() // step first line of input
	drawStrings := strings.Split(scanner.Text(), ",")
	var draw []int
	for _, item := range drawStrings {
		number, _ := strconv.ParseInt(item, 10, 0)
		draw = append(draw, int(number))
	}

	scanner.Scan() // step next line

	var board [][]int
	var boards [][][]int

	for scanner.Scan() {
		if scanner.Text() != "" {
			row := strings.Fields(scanner.Text())
			var rowInts []int
			for _, item := range row {
				number, _ := strconv.ParseInt(item, 10, 0)

				rowInts = append(rowInts, int(number))
			}

			board = append(board, rowInts)
		} else {
			boards = append(boards, board)
			board = nil
		}
	}
	return boards, draw
}

func checkBoards(boards [][][]int, draw []int) (int, int) {
	var drawn []int
	for i, board := range boards {
		for _, num := range draw {
			drawn = append(drawn, num)
			if checkBoardWins(board, drawn) {
				return len(drawn), i
			}
			break
		}
	}
	return 0, 0
}

func checkBoardWins(board [][]int, drawn []int) bool {
	// rows
	for r := 0; r < 5; r++ {
		marked := 0
		for c := 0; c < 5; c++ {
			for _, val := range drawn {
				if board[r][c] == val {
					marked++
				}
			}
		}

		if marked > 4 {
			return true
		}
	}

	// cols
	for c := 0; c < 5; c++ {
		marked := 0
		for r := 0; r < 5; r++ {
			for _, val := range drawn {
				if board[r][c] == val {
					marked++
				}
			}
		}
		if marked > 4 {
			return true
		}
	}
	return false
}
