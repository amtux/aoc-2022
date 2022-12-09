package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var INPUT_FILE = "./test.txt"

func main() {

	file, err := os.Open(INPUT_FILE)
	poe(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r := []rune(line)
		fmt.Print(string(r))
		println()
	}
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
