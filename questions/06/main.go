package main

import (
	"fmt"
	"os"
	"strings"
)

type QVariant int64

const (
	Part1 QVariant = 1
	Part2 QVariant = 2
)

func readLines(inputFile string) []string {
	dat, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

// If all runes in the map are unique
func isUnique(occurrences map[rune]int) bool {
	for _, v := range occurrences {
		if v > 1 {
			return false
		}
	}

	return true
}

// Returns the index at which the substring of length "totalUniqueChars" are all unique
func findStarterMarker(input string, totalUniqueChars int) int {
	// Two pointer. One at the start, and one at the end.
	// Load current packet into map of letter occurrences
	// if all values in the map are 1, return second pointer position
	// Otherwise, remove one occurrence of the starter pointer in the map by 1 (or delete if 0),
	// and increment all pointers by 1,
	runes := []rune(input)

	// If less than 4 letters, abort
	if len(runes) < totalUniqueChars {
		return -1
	}

	occurrences := map[rune]int{}
	for i := 0; i < totalUniqueChars-1; i += 1 {
		occurrences[runes[i]] += 1
	}

	for start, end := 0, totalUniqueChars-1; end < len(runes); start, end = start+1, end+1 {
		occurrences[runes[end]] += 1

		if isUnique(occurrences) {
			return end
		}

		startRune := runes[start]
		if occurrences[startRune] == 1 {
			// Delete it if it's 0 to speed up the `isUnique` check
			delete(occurrences, startRune)
		} else {
			occurrences[startRune] -= 1
		}
	}

	return -1
}

func parts(variant QVariant) {
	// There's only one line in this input
	input := readLines("./input.txt")[0]

	totalUniqueChars := 4
	if variant == Part2 {
		totalUniqueChars = 14
	}

	// Output the index for the question as if it was in a list where the first index is 1
	fmt.Println("Part ", variant, " - ", findStarterMarker(input, totalUniqueChars)+1)
}

func main() {
	parts(Part1)
	parts(Part2)
}
