package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	Left      = Pair{x: -1, y: 0}
	Right     = Pair{x: 1, y: 0}
	Up        = Pair{x: 0, y: -1}
	Down      = Pair{x: 0, y: 1}
	UpLeft    = Pair{x: -1, y: -1}
	UpRight   = Pair{x: 1, y: -1}
	DownLeft  = Pair{x: -1, y: 1}
	DownRight = Pair{x: 1, y: 1}
)

var directions = []Pair{
	Left, Right, Up, Down, UpLeft, UpRight, DownLeft, DownRight,
}

type Pair struct {
	x int
	y int
}

func (p Pair) equals(p2 Pair) bool {
	return p.x == p2.x && p.y == p2.y
}

func (p *Pair) add(p2 Pair) {
	p.x = p.x + p2.x
	p.y = p.y + p2.y
}

func (p Pair) toString() string {
	return fmt.Sprintf("x: %v, y: %v", p.x, p.y)
}

// Lazy normalize to avoid division
func (p *Pair) normalize() {
	if p.x > 0 {
		p.x = 1
	} else if p.x < 0 {
		p.x = -1
	}

	if p.y > 0 {
		p.y = 1
	} else if p.y < 0 {
		p.y = -1
	}
}

type Instruction struct {
	direction Pair
	distance  int
}

var inputToDirMap = map[string]Pair{
	"L": Left,
	"R": Right,
	"U": Up,
	"D": Down,
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
		instructions = append(instructions, Instruction{
			direction: inputToDirMap[tokens[0]],
			distance:  parseToIntOrPanic(tokens[1]),
		})
	}

	return instructions
}

// If pair1 is adjacent to pair2, including diagonal
func isAdjecent(pair1 Pair, pair2 Pair) bool {
	if pair1.equals(pair2) {
		return true
	}

	for _, dir := range directions {
		next := Pair{x: pair1.x + dir.x, y: pair1.y + dir.y}
		if next.equals(pair2) {
			return true
		}
	}

	return false
}

func runInstructions(instructions []Instruction, totalTails int) *map[Pair]bool {
	knots := make([]*Pair, totalTails+1)

	for i := range knots {
		knots[i] = &Pair{x: 0, y: 0}
	}

	tailVisited := &map[Pair]bool{
		*knots[0]: true,
	}

	for _, instr := range instructions {
		for i := 0; i < instr.distance; i += 1 {
			// printCurrentKnots(knots)
			for knotNum, knot := range knots {
				if knotNum == 0 {
					// head can move anywhere
					knot.add(instr.direction)
					continue
				}

				prevKnot := knots[knotNum-1]
				if !isAdjecent(*prevKnot, *knot) {
					// move in the direction of the knot in front
					dir := Pair{x: prevKnot.x - knot.x, y: prevKnot.y - knot.y}

					// Need to normalize though to a unit vector
					dir.normalize()
					knot.add(dir)

					// Last knot is the tail
					if knotNum == totalTails {
						(*tailVisited)[*knot] = true
					}
				}
			}
		}
	}

	return tailVisited
}

func countVisited(visited *map[Pair]bool) int {
	count := 0
	for _, v := range *visited {
		if v {
			count += 1
		}
	}

	return count
}

func printCurrentKnots(knots []*Pair) {
	// Set reasonable max mins
	minX := -15
	minY := -15
	maxX := 15
	maxY := 15

	knotMap := map[Pair]int{}
	for i, knot := range knots {
		knotMap[*knot] = i
	}

	for y := minY; y < maxY; y += 1 {
		for x := minX; x < maxX; x += 1 {
			value, exists := knotMap[Pair{x: x, y: y}]
			if exists {
				fmt.Print(value)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println("------")

	time.Sleep(1 * time.Second / 32)
}

func parts() {
	input := readLines("./input.txt")
	instructions := parseInstructions(input)

	visited := runInstructions(instructions, 1)
	count := countVisited(visited)

	fmt.Println("Part 1 - ", count)

	visited = runInstructions(instructions, 9)
	count = countVisited(visited)
	fmt.Println("Part 2 - ", count)
}

func main() {
	parts()
}
