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

	var data [][]byte
	var ogr int64
	var cor int64
	for scanner.Scan() {
		n := scanner.Bytes()
		data = append(data, n)
	}
	input.Close()

	ogr = calcOgr(data, "high", 1)
	cor = calcCor(data, "low", 0)

	fmt.Print(ogr, cor)
}

func calcOgr(workingData [][]byte, HL string, preference int) int64 {
	var result []byte
	for i := 0; i < len(workingData[0]); i++ {
		workingData = mostCommon(workingData, i, HL, preference)
		result = append(result, workingData[0][i])
		fmt.Println(workingData)
	}
	output, err := strconv.ParseInt(string(result), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func calcCor(workingData [][]byte, HL string, preference int) int64 {
	var result []byte
	for i := 0; i < len(workingData[0]); i++ {
		workingData = mostCommon(workingData, i, HL, preference)
		result = append(result, workingData[0][i])
	}
	output, err := strconv.ParseInt(string(result), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func mostCommon(list [][]byte, index int, highLow string, prefer int) [][]byte {
	var ones [][]byte
	var zeros [][]byte

	for i := 0; i < len(list); i++ {
		if list[i][index] == 48 {
			zeros = append(zeros, list[i])
		} else {
			ones = append(ones, list[i])
		}
	}

	if prefer == 0 {
		if highLow == "high" {
			if len(ones) > len(zeros) {
				return ones
			} else {
				return zeros
			}
		} else {
			if len(ones) > len(zeros) {
				return zeros
			} else {
				return ones
			}
		}
	} else {
		if highLow == "high" {
			if len(ones) >= len(zeros) {
				return ones
			} else {
				return zeros
			}
		} else {
			if len(ones) >= len(zeros) {
				return zeros
			} else {
				return ones
			}
		}
	}
}
