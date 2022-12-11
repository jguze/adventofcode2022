package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command int64

const (
	Noop Command = iota
	Addx
)

type Instruction struct {
	command Command
	value   int
}

var inputToCommandMap = map[string]Command{
	"noop": Noop,
	"addx": Addx,
}

func parseToIntOrPanic(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func readLines(inputFile string) []string {
	dat, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

func parseInstructions(input []string) []Instruction {
	instructions := []Instruction{}
	for _, line := range input {
		tokens := strings.Split(line, " ")
		instruction := Instruction{
			command: inputToCommandMap[tokens[0]],
		}
		if instruction.command == Addx {
			instruction.value = parseToIntOrPanic(tokens[1])
		}
		instructions = append(instructions, instruction)
	}

	return instructions
}

func runCommands(
	instructions []Instruction,
	cyclesToTrack map[int]bool,
	drawWhileRunning bool,
) map[int]int {
	register := 1
	maxInstrTime := 2

	cycleToValue := map[int]int{}

	cycleOperationBuffer := make([]int, maxInstrTime)
	for i := 0; i < maxInstrTime; i += 1 {
		cycleOperationBuffer[i] = 0
	}

	newlineCycleEvery := 40

	cycle := 0
	for _, instruction := range instructions {
		instructionLength := 0
		// Process instruction
		if instruction.command == Addx {
			instructionLength = 2
			cycleOperationBuffer[(cycle+1)%maxInstrTime] += instruction.value
		} else if instruction.command == Noop {
			instructionLength = 1
		}

		for i := 0; i < instructionLength; i += 1 {
			// The question is 1 indexed instead of 0, and we calculate
			// during the cycle, not at the end
			_, exists := cyclesToTrack[cycle+1]
			if exists {
				cycleToValue[cycle+1] = register
			}

			// Draw pixel
			if drawWhileRunning {
				if cycle%newlineCycleEvery == 0 {
					fmt.Println()
				}

				pixelPos := cycle % newlineCycleEvery
				if pixelPos-1 <= register && pixelPos+1 >= register {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}

			register += cycleOperationBuffer[cycle%maxInstrTime]
			cycleOperationBuffer[cycle%maxInstrTime] = 0

			cycle += 1
		}
	}

	return cycleToValue
}

func calcSignalStrength(cycleToValue map[int]int) int {
	signalStrength := 0
	for cycle, value := range cycleToValue {
		signalStrength = signalStrength + (cycle * value)
	}

	return signalStrength
}

func parts() {
	input := readLines("./input.txt")
	instructions := parseInstructions(input)

	cyclesToTrack := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}

	cycleToValue := runCommands(instructions, cyclesToTrack, false)
	signalStrength := calcSignalStrength(cycleToValue)

	fmt.Println("Part 1 - ", signalStrength)

	fmt.Println("Part 2")
	runCommands(instructions, cyclesToTrack, true)
}

func main() {
	parts()
}
