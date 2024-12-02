package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func find_min_and_reset(numbers *[]int) int {
	smallest := math.MaxInt32
	index := -1

	for i, num := range *numbers {
		if num != -1 && num < smallest {
			smallest = num
			index = i
		}
	}

	(*numbers)[index] = -1
	return smallest
}

func part_1(file *os.File) int {
	var left_numbers []int
	var right_numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "   ")

		num, _ := strconv.Atoi(data[0])
		left_numbers = append(left_numbers, num)

		num, _ = strconv.Atoi(data[1])
		right_numbers = append(right_numbers, num)

	}

	distance := 0
	for i := 0; i < len(left_numbers); i++ {
		rel_diff := find_min_and_reset(&left_numbers) - find_min_and_reset(&right_numbers)
		distance += int(math.Abs(float64(rel_diff)))
	}

	return distance
}

func part_2(file *os.File) int {
	scanner := bufio.NewScanner(file)
	var num_map = make(map[int]int)
	var right_nums []int

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "   ")

		left_num, _ := strconv.Atoi(data[0])
		right_num, _ := strconv.Atoi(data[1])

		num_map[left_num] = 0
		right_nums = append(right_nums, right_num)
	}

	for _, num := range right_nums {
		if _, ok := num_map[num]; ok {
			num_map[num] = num_map[num] + 1
		}
	}

	var score = 0
	for key, val := range num_map {
		score = score + (key * val)
	}

	return score
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

	distance := part_1(file)
	file.Seek(0, 0)
	score := part_2(file)

	fmt.Println("part1 sol:", distance)
	fmt.Println("part2 sol:", score)

}
