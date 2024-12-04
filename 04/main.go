package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(array []rune) bool {
	return string(array) == "XMAS" || string(array) == "SAMX"
}

func part_1(file *os.File) int {
	scanner := bufio.NewScanner(file)
	cont := 0
	var matrix [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			// Check right
			if j+4 <= len(matrix[i]) && check(matrix[i][j:j+4]) {
				cont++
			}

			// Check down
			if i+4 <= len(matrix) {
				column := []rune{matrix[i][j], matrix[i+1][j], matrix[i+2][j], matrix[i+3][j]}
				if check(column) {
					cont++
				}
			}

			// Check diagonally down
			if i+4 <= len(matrix) && j+4 <= len(matrix[i]) {
				diagonal := []rune{matrix[i][j], matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3]}
				if check(diagonal) {
					cont++
				}
			}

			// Check diagonally up
			if i-3 >= 0 && j+4 <= len(matrix[i]) {
				diagonal := []rune{matrix[i][j], matrix[i-1][j+1], matrix[i-2][j+2], matrix[i-3][j+3]}
				if check(diagonal) {
					cont++
				}
			}
		}
	}
	return cont
}

func main() {

	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Please provide an input file")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}

	defer file.Close()

	res1 := part_1(file)
	file.Seek(0, 0)
	// res2 := part_2(file)

	fmt.Println("part1 sol:", res1)
	// fmt.Println("part2 sol:", res2)

}
