package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var INPUT_FILE = "./input.txt"

func main() {
	/*
	   opponent me
	   A Y - A rock, Y paper
	   B X - B paper, X rock  = lose
	   C Z - C scissors, Z scissors
	*/
	file, err := os.Open(INPUT_FILE)
	poe(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var totalScore int
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		left := s[0]
		right := s[1]
		fmt.Printf(
			"left: '%s', right: '%s'\n",
			left,
			right,
		)
		moveScore := calcMoveScore(left, right)
		fmt.Printf("move score: %d\n", moveScore)
		totalScore += moveScore
	}
	fmt.Printf("total score: %d\n", totalScore)

}

func calcMoveScore(left, right string) int {
	// if left is rock
	// rightRune := []rune(right)
	// rr := rightRune[0]

	// var leftIndex int
	// switch left {
	// case "A":
	// 	// rock
	// 	leftIndex = 0
	// case "B":
	// 	// paper
	// 	leftIndex = 1
	// case "C":
	// 	// scissors
	// 	leftIndex = 2
	// }

	var result int
	var score int
	switch right {
	case "X":
		// lose
		switch left {
		case "A":
			score = 3
		case "B":
			score = 1
		case "C":
			score = 2
		}
		result = 0
	case "Y":
		// draw
		switch left {
		case "A":
			score = 1
		case "B":
			score = 2
		case "C":
			score = 3
		}
		result = 3
	case "Z":
		// win
		result = 6
		switch left {
		case "A":
			score = 2
		case "B":
			score = 3
		case "C":
			score = 1
		}
	}
	return score + result
}

// poe is panic on error
func poe(err error) {
	if err != nil {
		panic(err)
	}
}

func stoi(s string) int {
	i, err := strconv.Atoi(s)
	poe(err)
	return i

}
