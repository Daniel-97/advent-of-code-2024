package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part_1(file *os.File) int {

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
