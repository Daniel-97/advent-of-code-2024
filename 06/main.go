package main

import (
	"bufio"
	"fmt"
	"os"
)

func find_start_position(matrix [][]rune) (i int, j int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '^' {
				return i, j
			}
		}
	}

	return -1, -1
}

func part_1(file *os.File) int {
	scanner := bufio.NewScanner(file)
	visited := 0

	var matrix [][]rune

	for scanner.Scan() {
		line := scanner.Text()

		row := []rune(line)
		matrix = append(matrix, row)
	}

	i, j := find_start_position(matrix)
	direction := '^'
	for ok := true; ok; ok = (i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix)) {

		if matrix[i][j] != 'X' {
			matrix[i][j] = 'X'
			visited++
		}

		i, j, direction = walk(matrix, i, j, direction)
	}

	return visited

}

func walk(matrix [][]rune, i int, j int, direction rune) (new_i int, new_j int, new_direction rune) {
	if direction == '^' {
		if i-1 >= 0 && matrix[i-1][j] == '#' {
			direction = '>'
			j++
		} else {
			i--
		}
	} else if direction == 'v' {
		if i+1 < len(matrix) && matrix[i+1][j] == '#' {
			direction = '<'
			j--
		} else {
			i++
		}

	} else if direction == '>' {
		if j+1 < len(matrix[i]) && matrix[i][j+1] == '#' {
			direction = 'v'
			i++
		} else {
			j++
		}
	} else if direction == '<' {
		if j-1 >= 0 && matrix[i][j-1] == '#' {
			direction = '^'
			i--
		} else {
			j--
		}
	}
	return i, j, direction
}

func check_loop(matrix [][]rune, i int, j int) {
	direction := '^'
	for ok := true; ok; ok = (i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix)) {

		i, j, direction = walk(matrix, i, j, direction)
	}
}

func part_2(file *os.File) int {
	scanner := bufio.NewScanner(file)
	var matrix [][]rune

	for scanner.Scan() {
		line := scanner.Text()

		row := []rune(line)
		matrix = append(matrix, row)
	}

	i, j := find_start_position(matrix)

	return 0
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
	res2 := part_2(file)

	fmt.Println("part1 sol:", res1)
	fmt.Println("part2 sol:", res2)

}
