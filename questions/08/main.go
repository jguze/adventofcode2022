package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	Left  = Pair{x: -1, y: 0}
	Right = Pair{x: 1, y: 0}
	Up    = Pair{x: 0, y: -1}
	Down  = Pair{x: 0, y: 1}
)

var directions = []Pair{
	Left, Right, Up, Down,
}

type Cell struct {
	value       int
	visible     bool
	scenicScore map[Pair]int

	Pair
}

type Pair struct {
	x int
	y int
}

func (c Cell) toString() string {
	return fmt.Sprint(c.value)
}

func (c Cell) maxScenic() int {
	product := 1
	for _, dir := range directions {
		product *= c.scenicScore[dir]
	}

	return product
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

func loadGrid(input []string) [][]*Cell {
	height := len(input)
	width := len(input[0])

	grid := make([][]*Cell, height)
	for i := range grid {
		grid[i] = make([]*Cell, width)
	}

	for y, row := range input {
		for x, col := range row {
			grid[y][x] = &Cell{
				value:       parseToIntOrPanic(string(col)),
				visible:     false,
				scenicScore: map[Pair]int{},
				Pair:        Pair{x, y},
			}
		}
	}

	return grid
}

func printGrid(grid [][]*Cell) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col.toString())
		}
		fmt.Println()
	}
}

func printGridVis(grid [][]*Cell) {
	for _, row := range grid {
		for _, col := range row {
			if col.visible {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
	}
}

func isOutOfBounds(pair Pair, grid [][]*Cell) bool {
	// Assume square
	max := len(grid) - 1
	return pair.x < 0 || pair.y < 0 || pair.x > max || pair.y > max
}

// Walk in the direction until a tree higher than us exists.
// Return both if it can see the edge, and the total trees it has seen until
// then, including the tree blocking it.
func canSeeEdgeFromHeight(
	height int,
	currentCell *Cell,
	dir Pair,
	grid [][]*Cell,
	currentTotalTrees int,
) (bool, int) {
	nextX := currentCell.x + dir.x
	nextY := currentCell.y + dir.y

	if isOutOfBounds(Pair{x: nextX, y: nextY}, grid) {
		return true, currentTotalTrees
	}

	nextCell := grid[nextY][nextX]
	if nextCell.value >= height {
		return false, currentTotalTrees + 1
	}

	result, total := canSeeEdgeFromHeight(height, nextCell, dir, grid, currentTotalTrees+1)
	return result, total
}

func calcVisibilityAndScenicForCell(grid [][]*Cell, cell *Cell) {
	for _, dir := range directions {
		result, total := canSeeEdgeFromHeight(cell.value, cell, dir, grid, 0)
		if result {
			cell.visible = result
		}

		cell.scenicScore[dir] = total
	}
}

func countVisibleCells(grid [][]*Cell) int {
	total := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell.visible {
				total += 1
			}
		}
	}

	return total
}

func findMaxScenic(grid [][]*Cell) int {
	max := -1
	for _, row := range grid {
		for _, cell := range row {
			maxScenic := cell.maxScenic()
			if maxScenic > max {
				max = maxScenic
			}
		}
	}

	return max
}

func buildGridScores(grid [][]*Cell) {
	for _, row := range grid {
		for _, cell := range row {
			calcVisibilityAndScenicForCell(grid, cell)
		}
	}
}

func parts() {
	input := readLines("./input.txt")
	grid := loadGrid(input)
	buildGridScores(grid)

	total := countVisibleCells(grid)
	fmt.Println("Part 1 - ", total)

	maxScenic := findMaxScenic(grid)
	fmt.Println("Part 2 - ", maxScenic)
}

func main() {
	parts()
}
