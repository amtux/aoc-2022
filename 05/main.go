package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var INPUT_FILE = "./input.txt"

func main() {
	// temp stuff for loading
	totalStacks := 9
	var stackPos []int
	for i := 0; i < totalStacks; i++ {
		pos := i*4 + 1
		stackPos = append(stackPos, pos)
	}
	tempList := make([][]rune, totalStacks)
	for i := range tempList {
		tempList[i] = make([]rune, 0)
	}

	movesList := make([][]int, 0)

	// actual compute stack
	stacks := make([]stack, totalStacks)
	for i := range stacks {
		stacks[i] = make(stack, 0)
	}

	file, err := os.Open(INPUT_FILE)
	poe(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// var totalValue int
	var isStackLoaded bool
	for scanner.Scan() {
		line := scanner.Text()
		r := []rune(line)
		if !isStackLoaded {
			// fmt.Printf("pos are: %v,%v,%v|\n",
			// 	string(r[1]),
			// 	string(r[5]),
			// 	string(r[9]),
			// )
			if r[1] == 49 {
				// is r[1] == '1' then stop stacking
				isStackLoaded = true
				fmt.Println("stack is loaded")
			} else {
				// load up the stack
				for i, v := range stackPos {
					if r[v] != 32 { //dont stack space
						tempList[i] = append(tempList[i], r[v])
					}
				}
			}
		}
		//load moves
		if isStackLoaded && len(r) > 0 && r[0] == 109 {
			// execute if first letter is 'm'
			re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
			matches := re.FindAllString(line, -1)
			move := []int{
				stoi(matches[0]),
				stoi(matches[1]),
				stoi(matches[2]),
			}
			// move := []int{(int(r[5]) - '0'), (int(r[12]) - '0'), (int(r[17]) - '0')}
			movesList = append(movesList, move)
		}
	}
	// post processing
	for k := range tempList {
		// load it into the stack
		tempList[k] = reverseSlice(tempList[k])
		for _, v := range tempList[k] {
			// fmt.Printf("about to push %v into stack\n", string(v))
			stacks[k].Push(v)
		}
	}

	for _, move := range movesList {
		fmt.Printf("move %d times from %d to %d\n", move[0], move[1], move[2])
		from := move[1] - 1
		to := move[2] - 1
		var ts stack
		for i := 0; i < move[0]; i++ {
			poppedRune := stacks[from].Pop()
			// stacks[to].Push(poppedRune)
			ts.Push(poppedRune)
		}
		ts = reverseSlice(ts)
		for _, v := range ts {
			stacks[to].Push(v)

		}
	}
	for i := range stacks {
		fmt.Printf("FINAL: stack #%d is: %+v\n", i, string(stacks[i]))
	}
}

func reverseSlice(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
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

type stack []rune

// func (s stack) Push(v rune) stack {
// 	return append(s, v)
// }
//
// func (s stack) Pop() (stack, rune) {
// 	l := len(s)
// 	return s[:l-1], s[l-1]
// }
func (s *stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *stack) Pop() rune {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return ret
}
