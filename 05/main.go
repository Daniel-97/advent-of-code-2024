package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func in_array(array []string, data string) bool {
	for _, elem := range array {
		if elem == data {
			return true
		}
	}

	return false
}

func part_1(file *os.File) int {
	scanner := bufio.NewScanner(file)
	cont := 0
	var page_ordering = make(map[string][]string)
	section_1 := true

	for scanner.Scan() {
		line := scanner.Text()

		if section_1 && line == "" {
			section_1 = false
			continue
		}

		if section_1 {
			// Section one, page ordering
			data := strings.Split(line, "|")
			first_page := data[0]
			second_page := data[1]

			if page_ordering[first_page] != nil {
				page_ordering[first_page] = append(page_ordering[first_page], second_page)
			} else {
				page_ordering[first_page] = []string{second_page}
			}
		} else {
			// Section 2, update
			pages := strings.Split(line, ",")
			is_valid := true
			middle_page := 0

			for i, page := range pages {

				if i == (len(pages) / 2) {
					middle_page, _ = strconv.Atoi(page)
				}

				page_order := page_ordering[page]
				for _, next_page := range pages[i+1:] {
					if !in_array(page_order, next_page) {
						is_valid = false
						break
					}
				}
			}

			if is_valid {
				cont += middle_page
			}

		}
	}
	return cont
}

func swap(array []string, index_1 int, index_2 int) {
	tmp := array[index_1]
	array[index_1] = array[index_2]
	array[index_2] = tmp
}

func part_2(file *os.File) int {
	scanner := bufio.NewScanner(file)
	cont := 0
	var page_ordering = make(map[string][]string)
	section_1 := true

	for scanner.Scan() {
		line := scanner.Text()

		if section_1 && line == "" {
			section_1 = false
			continue
		}

		if section_1 {
			// Section one, page ordering
			data := strings.Split(line, "|")
			first_page := data[0]
			second_page := data[1]

			if page_ordering[first_page] != nil {
				page_ordering[first_page] = append(page_ordering[first_page], second_page)
			} else {
				page_ordering[first_page] = []string{second_page}
			}
		} else {
			// Section 2, update
			pages := strings.Split(line, ",")

			is_valid := true
			middle_page := 0

			for i, page := range pages {

				page_order := page_ordering[page]
				for _, next_page := range pages[i+1:] {
					if !in_array(page_order, next_page) {
						is_valid = false
					}
				}
			}

			if !is_valid {
				// Try to corretct the update
				sort.Slice(pages, func(i int, j int) bool {
					page := pages[i]
					next_page := pages[j]
					page_order := page_ordering[page]
					return !in_array(page_order, next_page)
				})

				middle_page, _ = strconv.Atoi(pages[len(pages)/2])
				cont += middle_page
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
	res2 := part_2(file)

	fmt.Println("part1 sol:", res1)
	fmt.Println("part2 sol:", res2)

}
