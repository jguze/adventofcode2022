package main

import (
	"fmt"
	"os"
	"strings"
)

type Bundle struct {
	full      string
	start     map[rune]int
	end       map[rune]int
	all       map[rune]int
	bundleNum int
}

func readLines(inputFile string) []string {
	dat, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

func countItems(s string) map[rune]int {
	count := map[rune]int{}

	for _, c := range s {
		count[c] += 1
	}

	return count
}

// Finds the first common char in the bundle
func findFirstCommon(bundle Bundle) rune {
	for k, _ := range bundle.start {
		if bundle.end[k] > 0 {
			return k
		}
	}

	panic(
		fmt.Sprint(
			"No common char found in bundle ",
			bundle.bundleNum,
			len(bundle.start),
			len(bundle.end),
		),
	)
}

func findCommonInBundles(bundles []Bundle) rune {
	set := map[rune]int{}

	for _, bundle := range bundles {
		for k, v := range bundle.all {
			if v >= 0 {
				set[k] += 1
			}
		}
	}

	for k, v := range set {
		if v == len(bundles) {
			return k
		}
	}

	panic(
		fmt.Sprint(
			"No common char found in bundles",
			bundles[0].bundleNum,
		),
	)
}

// Assume input is even
func createBundle(input string, bundleNum int) Bundle {
	length := len(input)
	halfway := (length / 2) - 1

	return Bundle{
		full:      input,
		start:     countItems(input[0 : halfway+1]),
		end:       countItems(input[halfway+1:]),
		all:       countItems(input),
		bundleNum: bundleNum,
	}
}

// a - z is 1 - 26, and A - Z i 27 - 52
func convertItemToValue(item rune) rune {
	lowercaseBaseValue := rune('a') - 1
	uppercaseBaseValue := rune('A') - 1 - 26 // Uppercase values are 26 above

	if item >= lowercaseBaseValue {
		return item - lowercaseBaseValue
	}

	return item - uppercaseBaseValue
}

func parts() {
	input := readLines("./input.txt")

	var bundles []Bundle

	for i, value := range input {
		bundles = append(bundles, createBundle(value, i))
	}

	score := 0
	for _, bundle := range bundles {
		common := findFirstCommon(bundle)
		// fmt.Println(
		// 	"Common - ",
		// 	common,
		// 	", ",
		// 	string(common),
		// 	", value: ",
		// 	convertItemToValue(common),
		// )
		score += int(convertItemToValue(common))
	}

	fmt.Println("Part 1 - ", score)

	groups := 3
	score = 0
	for i := 0; i < len(bundles); i += groups {
		common := findCommonInBundles(bundles[i : i+groups])
		score += int(convertItemToValue(common))
	}

	fmt.Println("Part 2 - ", score)
}

func main() {
	parts()
}
