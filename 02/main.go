package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func compare_val(val1 int, val2 int) (is_inc bool, is_dec bool, is_safe bool) {
	return val1 < val2, val1 > val2, math.Abs(float64(val1-val2)) > 0 && math.Abs(float64(val1-val2)) <= 3

}

func part_1(file *os.File) int {
	scanner := bufio.NewScanner(file)
	safe_reports := 0

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")

		is_safe := true
		is_increasing := true
		is_decresing := true

		for i := 0; i < len(data)-1; i++ {
			val, _ := strconv.Atoi(data[i])
			next_val, _ := strconv.Atoi(data[i+1])

			inc, dec, safe := compare_val(val, next_val)
			is_increasing = is_increasing && inc
			is_decresing = is_decresing && dec
			is_safe = safe
			if !is_safe {
				break
			}
		}

		if is_safe && (is_decresing || is_increasing) {
			safe_reports = safe_reports + 1
		}
	}

	return safe_reports

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
	// score := part_2(file)

	fmt.Println("part1 sol:", res1)
	// fmt.Println("part2 sol:", score)

}
