package main

import (
	"fmt"
	"os"
	"strings"
)

type Move int64

const (
	Rock Move = iota
	Paper
	Scissors
)

type Outcome int64

const (
	Win Outcome = iota
	Lose
	Draw
)

type Play struct {
	opponent Move
	mine     Move
}

type MoveOutcome struct {
	opponent Move
	outcome  Outcome
}

var inputToMove = map[string]Move{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var inputToOutcome = map[string]Outcome{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

var moveToScore = map[Move]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var outcomeToScore = map[Outcome]int{
	Lose: 0,
	Draw: 3,
	Win:  6,
}

var moveBeats = map[Move]Move{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var moveLosesTo = map[Move]Move{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

func readLines(inputFile string) []string {
	dat, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

// Did a win against b
func didWin(a Move, b Move) bool {
	return b == moveBeats[a]
}

func playGameP1(play Play) int {
	score := moveToScore[play.mine]
	if play.opponent == play.mine {
		return score + outcomeToScore[Draw]
	}

	if didWin(play.mine, play.opponent) {
		return score + outcomeToScore[Win]
	} else {
		return score + outcomeToScore[Lose]
	}
}

func playGameP2(moveOutcome MoveOutcome) int {
	score := outcomeToScore[moveOutcome.outcome]
	if moveOutcome.outcome == Draw {
		return score + moveToScore[moveOutcome.opponent]
	}

	if moveOutcome.outcome == Win {
		return score + moveToScore[moveLosesTo[moveOutcome.opponent]]
	} else {
		return score + moveToScore[moveBeats[moveOutcome.opponent]]
	}
}

func part1() {
	input := readLines("./input.txt")

	var gameMoves []Play
	for _, value := range input {
		splitValues := strings.Split(value, " ")
		gameMoves = append(
			gameMoves,
			Play{opponent: inputToMove[splitValues[0]], mine: inputToMove[splitValues[1]]},
		)
	}

	score := 0
	for _, play := range gameMoves {
		score += playGameP1(play)
	}

	fmt.Println("Part 1 - Score ", score)
}

func part2() {
	input := readLines("./input.txt")

	var gameMoves []MoveOutcome
	for _, value := range input {
		splitValues := strings.Split(value, " ")
		gameMoves = append(
			gameMoves,
			MoveOutcome{
				opponent: inputToMove[splitValues[0]],
				outcome:  inputToOutcome[splitValues[1]],
			},
		)
	}

	score := 0
	for _, play := range gameMoves {
		score += playGameP2(play)
	}

	fmt.Println("Part 2 - Score ", score)
}

func main() {
	part1()
	part2()
}
