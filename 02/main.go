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

func is_safe(level []string) bool {
	is_safe := true
	is_increasing := true
	is_decresing := true

	for i := 0; i < len(level)-1; i++ {
		val, _ := strconv.Atoi(level[i])
		next_val, _ := strconv.Atoi(level[i+1])

		inc, dec, safe := compare_val(val, next_val)
		is_increasing = is_increasing && inc
		is_decresing = is_decresing && dec
		is_safe = safe
		if !is_safe || !((is_increasing && inc) || (is_decresing && dec)) {
			return false
		}
	}

	return true
}

func remove_index(array []string, index int) []string {
	copied := make([]string, len(array))
	copy(copied, array)
	return append(copied[:index], copied[index+1:]...)
}

func part_1(file *os.File) int {
	scanner := bufio.NewScanner(file)
	safe_reports := 0

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")

		if is_safe(data) {
			safe_reports = safe_reports + 1
		}
	}

	return safe_reports

}

func part_2(file *os.File) int {
	scanner := bufio.NewScanner(file)
	safe_reports := 0

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")

		if is_safe(data) {
			safe_reports = safe_reports + 1
		} else {
			for i := 0; i < len(data); i++ {
				new_data := remove_index(data, i)
				if is_safe(new_data) {
					safe_reports = safe_reports + 1
					break
				}
			}
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
	res2 := part_2(file)

	fmt.Println("part1 sol:", res1)
	fmt.Println("part2 sol:", res2)

}
