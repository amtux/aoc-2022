package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var INPUT_FILE = "./input.txt"

func main() {
	file, err := os.Open(INPUT_FILE)
	poe(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// we only need max calories
	// track total cals in
	var thisCal int
	var totalCals []int

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(scanner.Text())
		if line == "" {
			totalCals = append(totalCals, thisCal)
			thisCal = 0
		} else {
			thisCal += stoi(line)
			// fmt.Printf("this cal value: %d \n", thisCal)
		}
	}
	// pick up last this
	totalCals = append(totalCals, thisCal)
	sort.Sort(sort.Reverse(sort.IntSlice(totalCals)))

	poe(scanner.Err())
	fmt.Printf("max cals: %d\n", totalCals)
	fmt.Printf("top 3 total: %d\n", totalCals[0]+totalCals[1]+totalCals[2])
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
