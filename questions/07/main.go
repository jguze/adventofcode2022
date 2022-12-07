package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type QVariant int64

const (
	Part1 QVariant = 1
	Part2 QVariant = 2
)

type CommandType int64

const (
	CD CommandType = iota
	LS
)

type TerminalState int64

const (
	Command TerminalState = iota
	Output
)

var stringToCommand = map[string]CommandType{
	"cd": CD,
	"ls": LS,
}

type Node struct {
	parent   *Node
	children map[string]*Node
	size     int
	isDir    bool
	name     string
}

func (node *Node) computeSize() int {
	// If a file, just return the size
	if !node.isDir {
		return node.size
	}

	size := 0
	for _, child := range node.children {
		size += child.computeSize()
	}

	return size
}

func (node *Node) toString() string {
	return fmt.Sprintf("%v: size - %v, isDir - %v", node.name, node.computeSize(), node.isDir)
}

func NewNode(parent *Node, size int, isDir bool, name string) *Node {
	return &Node{
		parent:   parent,
		children: map[string]*Node{},
		size:     size,
		isDir:    isDir,
		name:     name,
	}
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

func isCommand(line string) bool {
	return string(line[0]) == "$"
}

func handleCommand(line string, currentNode *Node, root *Node) *Node {
	tokens := strings.Split(line, " ")
	// command is always at index 1

	command, exists := stringToCommand[tokens[1]]
	if !exists {
		panic(fmt.Sprintf("token is not a command: ", tokens[1]))
	}

	if command == CD {
		dir := tokens[2]
		if dir == "/" {
			return root
		} else if dir == ".." {
			return currentNode.parent
		} else {
			// Look for existing dir, or create one
			_, dirExists := currentNode.children[dir]
			if dirExists {
				return currentNode.children[dir]
			} else {
				newNode := NewNode(currentNode, 0, true, dir)
				currentNode.children[dir] = newNode
				return newNode
			}
		}
	} else if command == LS {
		// Because the only thing not a command is output from ls, we can just no-op
		return currentNode
	} else {
		panic(fmt.Sprintf("No handler set up for command: %v", command))
	}
}

func handleOutput(line string, currentNode *Node) {
	tokens := strings.Split(line, " ")
	if tokens[0] == "dir" {
		dirName := tokens[1]
		_, exists := currentNode.children[dirName]
		if !exists {
			currentNode.children[dirName] = NewNode(currentNode, 0, true, dirName)
		}
	} else {
		// Must be file size
		size := parseToIntOrPanic(tokens[0])
		filename := tokens[1]

		fileNode, exists := currentNode.children[filename]
		if !exists {
			currentNode.children[filename] = NewNode(currentNode, size, false, filename)
		} else {
			fileNode.size = size
		}
	}
}

func parseInput(input []string) *Node {
	// Assume the first like is $ cd / and throw it out
	root := NewNode(nil, 0, true, "/")
	input = input[1:]

	currentNode := root
	for _, line := range input {
		if isCommand(line) {
			currentNode = handleCommand(line, currentNode, root)
		} else {
			handleOutput(line, currentNode)
		}
	}

	return root
}

func listDirs(root *Node) []*Node {
	nodesToVisit := []*Node{
		root,
	}

	directories := []*Node{root}

	for len(nodesToVisit) > 0 {
		currentNode := nodesToVisit[0]
		nodesToVisit = nodesToVisit[1:]
		for _, node := range currentNode.children {
			if node.isDir {
				directories = append(directories, node)
				nodesToVisit = append(nodesToVisit, node)
			}
		}
	}

	return directories
}

func part1() {
	input := readLines("./input.txt")

	root := parseInput(input)
	directories := listDirs(root)

	totalSize := 0
	for _, dir := range directories {
		dirSize := dir.computeSize()
		if dirSize <= 100000 {
			totalSize += dirSize
		}
	}

	fmt.Println("Part 1 - ", totalSize)
}

func part2() {
	input := readLines("./input.txt")

	root := parseInput(input)
	directories := listDirs(root)

	sort.SliceStable(directories, func(i, j int) bool {
		return directories[i].computeSize() < directories[j].computeSize()
	})

	totalSpace := 70000000
	requiredSpace := 30000000

	currentSpace := totalSpace - root.computeSize()

	for _, dir := range directories {
		dirSize := dir.computeSize()
		if currentSpace+dir.computeSize() >= requiredSpace {
			fmt.Println("Part 2 - ", dirSize)
			return
		}
	}
}

func main() {
	part1()
	part2()
}
