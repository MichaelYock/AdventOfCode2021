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

	winTurn, board, drawn := checkBoards(boards, draw)
	finalScore := calcScore(board, drawn)

	fmt.Println("Turn:", winTurn, ", Board:", board, " Final Draw:", draw[winTurn-1], " Score:", finalScore*draw[winTurn-1])
}

func setUpBoards() ([][][]int, []int) {
	file := os.Args[1]
	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

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

func checkBoards(boards [][][]int, draw []int) (int, [][]int, []int) {
	var drawn []int

	for _, num := range draw {
		drawn = append(drawn, num)
		for _, board := range boards {
			if checkBoardWins(board, drawn) {
				return len(drawn), board, drawn
			}
		}
	}
	return 0, nil, nil
}

func checkBoardWins(board [][]int, drawn []int) bool {
	// rows
	for r := 0; r < 5; r++ {
		marked := 0
		for c := 0; c < 5; c++ {
			if contains(drawn, board[r][c]) {
				marked++
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
			if contains(drawn, board[r][c]) {
				marked++
			}
		}
		if marked > 4 {
			return true
		}
	}
	return false
}

func calcScore(board [][]int, drawn []int) int {
	score := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !contains(drawn, board[r][c]) {
				score += board[r][c]
			}
		}
	}
	return score
}

func contains(arr []int, b int) bool {
	for _, val := range arr {
		if b == val {
			return true
		}
	}
	return false
}
