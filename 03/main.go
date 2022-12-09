package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
	// "github.com/juliangruber/go-intersect"
)

var INPUT_FILE = "./input.txt"

func main() {
	file, err := os.Open(INPUT_FILE)
	poe(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var totalValue int
	groupTotal := 3
	group := make([][]rune, groupTotal)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()

		lineNumMod := lineNum % 3
		// fmt.Printf("lineNumMod is %d\n", lineNumMod)
		// fmt.Printf("adding to group lineNumMod: %d\n", lineNumMod)
		group[lineNumMod] = []rune(line)

		if lineNumMod == groupTotal-1 {
			// check for intersection
			is := intersect(group[0], group[1], group[2])
			fmt.Printf("group is: %s, %s, %s\n", string(group[0]), string(group[1]), string(group[2]))
			value := getValue(is[0])
			// fmt.Printf("value is %d\n", value)
			totalValue += value
		}
		lineNum++
	}

	fmt.Printf("total value is: %d\n", totalValue)
}

func getValue(r rune) int {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	if unicode.IsUpper(r) {
		return int(r) - 64 + 26

	}
	return int(r) - 96
}

func intersect(a, b, c []rune) []rune {
	var res []rune
	for _, i := range a {
		for _, j := range b {
			for _, k := range c {
				if i == j && j == k && !elemExists(i, res) {
					res = append(res, i)
				}
			}
		}
	}
	return res
}

func elemExists(r rune, a []rune) bool {
	for _, i := range a {
		if r == i {
			return true
		}
	}
	return false
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
