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

	var left_numbers []int
	var right_numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "   ")

		if len(data) != 2 {
			fmt.Println("Error: line does not contain two numbers")
			continue
		}

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

	fmt.Println("The total distance is ", distance)

}
