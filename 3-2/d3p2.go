package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file := os.Args[1]
	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	var data [][]byte
	var ogr []byte
	var cor []byte

	for scanner.Scan() {
		data = append(data, scanner.Bytes())
	}
	wd := data

	for i, _ := range data[0] {
		if len(wd) != 0 {
			wd = mostCommon(wd, i, true)
			ogr = append(ogr, wd[0][i])
		}

	}

	wd = data
	for i, _ := range data[0] {
		if len(wd) != 0 {
			wd = mostCommon(wd, i, false)
			cor = append(cor, wd[0][i])
		}
	}

	ogrInt, err := strconv.ParseInt(string(ogr), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	corInt, err := strconv.ParseInt(string(cor), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ogr, cor)
	fmt.Println(ogrInt, corInt)
	fmt.Println(ogrInt * corInt)
}

func mostCommon(workingData [][]byte, targetIndex int, largest bool) [][]byte {
	ones := 0
	zeros := 0
	var targetNum byte
	var result [][]byte

	for _, val := range workingData {
		if val[targetIndex] == 48 {
			zeros++
		} else {
			ones++
		}
	}

	if largest {
		if ones >= zeros {
			targetNum = 49
		} else {
			targetNum = 48
		}
	} else {
		if zeros > ones {
			targetNum = 49
		} else {
			targetNum = 48
		}
		if zeros < 1 {
			targetNum = 49
		}
		if ones < 1 {
			targetNum = 48
		}
	}

	for _, val := range workingData {
		if val[targetIndex] == targetNum {
			result = append(result, val)
		}
	}
	return result
}

// 1912335
