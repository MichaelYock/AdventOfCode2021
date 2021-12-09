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
	var ones [][]byte
	var zeros [][]byte

	for _, val := range workingData {
		if val[targetIndex] == 48 {
			zeros = append(zeros, val)
		} else {
			ones = append(ones, val)
		}
	}
	fmt.Println("ones:", len(ones), "zeros:", len(zeros))
	if largest {
		if len(zeros) == len(ones) {
			return ones
		}
		if len(zeros) > len(ones) {
			return zeros
		} else {
			return ones
		}
	} else {
		if len(zeros) == len(ones) {
			return zeros
		}
		if len(ones) > len(zeros) && len(zeros) != 0 {
			return zeros
		} else if len(ones) != 0 {
			return ones
		} else {
			return zeros
		}
	}
}

// 1912335
