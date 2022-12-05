package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

type WorkerGroup struct {
	first    Range
	second   Range
	groupNum int
}

func readLines(inputFile string) []string {
	dat, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

// Creates a range from a number like 1-4
func createRange(str string) Range {
	strSplit := strings.Split(str, "-")

	start, err := strconv.Atoi(strSplit[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(strSplit[1])
	if err != nil {
		panic(err)
	}

	return Range{
		start: start,
		end:   end,
	}
}

func createGroup(input string, groupNum int) WorkerGroup {
	ranges := strings.Split(input, ",")
	fmt.Println(ranges)

	return WorkerGroup{
		first:    createRange(ranges[0]),
		second:   createRange(ranges[1]),
		groupNum: groupNum,
	}
}

// 1. Partial overlap where first.start is less than second.end
// 2. Partial overlap where second.start is less than first.end
// 3. Full overlap where first is contained by second
// 4. Full overlap where second is contained by first
func findOverlap(group WorkerGroup, fullOverlapOnly bool) *Range {
	if !fullOverlapOnly && group.second.start <= group.first.start &&
		group.first.start <= group.second.end {
		return &Range{
			start: group.first.start,
			end:   group.second.end,
		}
	} else if !fullOverlapOnly && group.first.start <= group.second.start && group.second.start <= group.first.end {
		return &Range{
			start: group.second.start,
			end:   group.first.end,
		}
	} else if group.first.start >= group.second.start && group.first.end <= group.second.end {
		return &group.first
	} else if group.second.start >= group.first.start && group.second.end <= group.first.end {
		return &group.second
	}

	return nil
}

func findCompleteOverlap(group WorkerGroup) *Range {
	return findOverlap(group, true)
}

func findPartialOverlap(group WorkerGroup) *Range {
	return findOverlap(group, false)
}

func parts() {
	input := readLines("./input.txt")

	overlapFull := 0
	overlapPartial := 0
	for i, value := range input {
		newGroup := createGroup(value, i)
		if findCompleteOverlap(newGroup) != nil {
			overlapFull += 1
		}

		if findPartialOverlap(newGroup) != nil {
			overlapPartial += 1
		}
	}

	fmt.Println("Part 1 - ", overlapFull)
	fmt.Println("Part 2 - ", overlapPartial)
}

func main() {
	parts()
}
