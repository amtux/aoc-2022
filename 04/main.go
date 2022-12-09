package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "github.com/juliangruber/go-intersect"
)

var INPUT_FILE = "./input.txt"

func main() {
	/*
	  2-4,6-8
	  2-3,4-5
	  5-7,7-9
	  2-8,3-7
	  6-6,4-6
	  2-6,4-8
	*/
	file, err := os.Open(INPUT_FILE)
	poe(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var totalValue int
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")
		var p []int

		for _, i := range pairs {
			sec := strings.Split(i, "-")
			for _, j := range sec {
				p = append(p, stoi(j))
			}
		}
		isFullyContained := checkFullyContains(p)

		fmt.Printf("yo p is %v. fully: %t\n", p, isFullyContained)
		if isFullyContained {
			totalValue++
		}

	}

	fmt.Printf("total value is: %d\n", totalValue)
}

func checkFullyContains(p []int) bool {
	l1 := p[0]
	l2 := p[1]
	r1 := p[2]
	r2 := p[3]

	if l1 <= r1 {
		//l1 is on the left most side or equal
		if l2 >= r1 {
			// l2 is overlapping r1 or equal
			return true
		}
	}
	if r1 <= l1 {
		//r1 is on the left most side or equal
		if r2 >= l1 {
			// r2 is overlapping l1 or equal to it
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
