package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
