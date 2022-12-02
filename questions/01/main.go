package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(inputFile string) []string {
	dat, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

func parts() {
	input := readLines("./input.txt")

	var elfCalories []int

	currentElf := 0
	elfCalories = append(elfCalories, 0)
	for _, value := range input {
		if len(value) == 0 {
			currentElf += 1
			elfCalories = append(elfCalories, 0)
		} else {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			elfCalories[currentElf] += intValue
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))

	fmt.Println("Part 1 - Max: ", elfCalories[0])
	fmt.Println("Part 2 - ", elfCalories[0]+elfCalories[1]+elfCalories[2])
}

func main() {
	parts()
}
