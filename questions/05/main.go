package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type QVariant int64

const (
	Part1 QVariant = 1
	Part2 QVariant = 2
)

type Stack []string

// Push a single item onto the stack
func (s *Stack) Push(v string) Stack {
	*s = append(*s, v)
	return *s
}

// Push group onto stack "in order". So the whole
// stack is simply insterted on top
func (s *Stack) PushGroup(v Stack) Stack {
	*s = append(*s, v...)
	return *s
}

// Pops a single item from the stack
func (s *Stack) Pop() string {
	last := len(*s) - 1
	value := (*s)[last]
	*s = (*s)[:last]
	return value
}

// Pops a group of items, keeping their order.
// So if the total is 3, it will take the top 3 items off
// and maintain their order
func (s *Stack) PopGroup(total int) Stack {
	length := len(*s)
	values := (*s)[length-total:]
	*s = (*s)[:length-total]

	return values
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		panic("Stack is empty. cannot peek")
	}

	return (*s)[len(*s)-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

type Instruction struct {
	source      int
	destination int
	total       int
}

var instructionRegex, _ = regexp.Compile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)

func readLines(inputFile string) []string {
	dat, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

func parseToIntOrPanic(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

// Rows come in with the legend on the bottom
// Let's reverse the list, and start from the bottom.
func parseRows(rows []string) []Stack {

	length := len(rows)
	legend := strings.Trim(rows[length-1], " ")

	colNames := strings.Split(legend, " ")
	totalStacks, err := strconv.Atoi(colNames[len(colNames)-1])
	if err != nil {
		panic(err)
	}

	stacks := []Stack{}
	for i := 0; i < totalStacks; i += 1 {
		stacks = append(stacks, Stack{})
	}

	// Every column is always a width of 3, and then space delimited after
	colWidth := 3
	for rowIndex := length - 2; rowIndex >= 0; rowIndex -= 1 {
		row := rows[rowIndex]

		for colIndex, stackIndex := 0, 0; colIndex < len(row); colIndex, stackIndex = colIndex+colWidth, stackIndex+1 {
			value := row[colIndex : colIndex+colWidth]
			if value[0] == '[' {
				stacks[stackIndex] = append(stacks[stackIndex], string(value[1]))
			}

			// Assume space for column
			colIndex += 1
		}
	}

	return stacks
}

func parseInstructions(instructionsInput []string) []Instruction {
	instructions := []Instruction{}
	for _, input := range instructionsInput {
		// This returns the original string as match 0
		matches := instructionRegex.FindStringSubmatch(input)
		instructions = append(instructions, Instruction{
			source: parseToIntOrPanic(
				matches[2],
			) - 1, // 0 index the stacks. Instruction input assumes the first column is 1
			destination: parseToIntOrPanic(matches[3]) - 1,
			total:       parseToIntOrPanic(matches[1]),
		})
	}

	return instructions
}

// Returns all the stacks, and instructions
func parseInput(filename string) ([]Stack, []Instruction) {
	input := readLines(filename)

	// It'll be easier to start bottom to top to parse the row
	// We'll split the input into the rows and instructions
	rowsInput := []string{}
	instructionsInput := []string{}

	writeToRow := true
	for _, value := range input {
		if value == "" {
			writeToRow = false
		} else if writeToRow {
			rowsInput = append(rowsInput, value)
		} else {
			instructionsInput = append(instructionsInput, value)
		}
	}

	return parseRows(rowsInput), parseInstructions(instructionsInput)
}

func runInstructions(stacks []Stack, instructions []Instruction, variant QVariant) {
	for _, instruction := range instructions {
		sourceStack := stacks[instruction.source]
		destStack := stacks[instruction.destination]
		if variant == Part1 {
			for i := 0; i < instruction.total; i += 1 {
				if !sourceStack.IsEmpty() {
					destStack.Push(sourceStack.Pop())
				}
			}
		} else {
			if !sourceStack.IsEmpty() {
				totalToTake := min(instruction.total, len(sourceStack))
				destStack.PushGroup(sourceStack.PopGroup(totalToTake))
			}
		}

		stacks[instruction.source] = sourceStack
		stacks[instruction.destination] = destStack
	}

}

func parts(variant QVariant) {
	stacks, instructions := parseInput("./input.txt")

	runInstructions(stacks, instructions, variant)
	tops := ""
	for _, stack := range stacks {
		if !stack.IsEmpty() {
			tops += stack.Peek()
		}
	}

	fmt.Println("Part ", variant, " - ", tops)
}

func main() {
	parts(Part1)
	parts(Part2)
}
