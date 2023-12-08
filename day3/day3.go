package main

import (
	"bufio"
	"os"
	"unicode"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("./test")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	line1 := []rune{}
	line2 := []rune{}
	line3 := []rune{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()

		// as I scan through lines, I want to check if there are symbols adjacent to numbers
		// as apposed to the reverse, so if 35 occupies 2:3 and 2:4 I need to check to see
		// if there is a non-. symbol at 1:2, 1:3, 1:4, 1:5, 2:2, 2:5, 3:2, 3:3, 3:4, or 3:5
		// So I need to store the line above and below each number into memory, split each
		// of the spaces into a rune, determine which runes are part of entire numbers
	}
	println("Part 1:")
}

func part2() {

}

func isRuneAdjacent(positions []int, line1 []rune, line2 []rune, line3 []rune) bool {
	for _, pos := range positions {
		return isSymbol(line1[pos-1]) ||
			isSymbol(line1[pos]) ||
			isSymbol(line1[pos+1]) ||
			isSymbol(line2[pos-1]) ||
			isSymbol(line2[pos]) ||
			isSymbol(line2[pos+1]) ||
			isSymbol(line3[pos-1]) ||
			isSymbol(line3[pos]) ||
			isSymbol(line3[pos+1])
	}
	return false
}

func isSymbol(char rune) bool {
	if char != '.' && !unicode.IsNumber(char) {
		return true
	}
	return false
}
