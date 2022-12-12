package main

import (
	"fmt"
	"strconv"
)

type Board struct {
	stacks map[int]*Queue[string]
}

func (b *Board) Add(stack int, item string) {
	b.stacks[stack].Push(item)
}

func splitBy(input []string, sep string) ([]string, []string) {
	for i, line := range input {
		if line == sep {
			return input[:i], input[i+1:]
		}
	}
	panic("No separator found")
}

func tryParse(a rune, out *int) bool {
	res, err := strconv.Atoi(string(a))
	if err != nil {
		return false
	} else {
		*out = res
		return true
	}
}

func day5() {
	input := readFile("inputs/day5.txt")

	a, _ := splitBy(input, "")

	for row := len(a) - 1; row >= 0; row-- {
		for col, c := range a[row] {
			res := 0
			if tryParse(c, &res) {
				fmt.Printf("number at %d: %d \n", col, res)
			}
		}
	}
	// printList(a)
}
