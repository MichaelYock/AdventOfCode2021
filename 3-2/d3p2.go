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
	defer input.Close()
	scanner := bufio.NewScanner(input)
	var ogrData []string
	for scanner.Scan() {
		ogrData = append(ogrData, scanner.Text())
	}
	corData := append([]string{}, ogrData...)

	ogr, err := strconv.ParseInt(findCommon(ogrData, 0, true), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	cor, err := strconv.ParseInt(findCommon(corData, 0, false), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ogr, cor)
	fmt.Println(ogr * cor)
}

func findCommon(data []string, targetIndex int, mostLeast bool) string {
	if len(data) == 1 {
		return data[0]
	}

	var ones, zeros []string
	for _, val := range data {
		if rune(val[targetIndex]) == '0' {
			zeros = append(zeros, val)
		} else {
			ones = append(ones, val)
		}
	}
	if len(ones) >= len(zeros) == mostLeast {
		return findCommon(ones, targetIndex+1, mostLeast)
	}
	return findCommon(zeros, targetIndex+1, mostLeast)
}
