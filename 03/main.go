package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part_1(file *os.File) int {
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])

			sum = sum + x*y
		}
	}

	return sum

}

func part_2(file *os.File) int {
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		re_do := regexp.MustCompile(`do\(\)`)
		re_do_not := regexp.MustCompile(`don't\(\)`)

		matches := re.FindAllStringSubmatchIndex(line, -1)
		matches_do := re_do.FindAllStringSubmatchIndex(line, -1)
		matches_do_not := re_do_not.FindAllStringSubmatchIndex(line, -1)

		for _, match := range matches {
			x, _ := strconv.Atoi(line[match[2]:match[3]])
			y, _ := strconv.Atoi(line[match[4]:match[5]])

			sum = sum + x*y
		}
	}

	return sum

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
