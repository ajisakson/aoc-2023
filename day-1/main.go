package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func part1() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		var firstInt rune = 'i'
		var lastInt rune = 'j'
		// get the first integer
		i := 0
		for firstInt == 'i' {
			for i < len(text) {
				char := rune(text[i])
				if unicode.IsDigit(char) {
					firstInt = char
					break
				}
				i++
			}

		}

		j := len(text) - 1
		for lastInt == 'j' {
			for j >= 0 {
				char := rune(text[j])
				if unicode.IsDigit(char) {
					lastInt = char
					break
				}
				j--
			}

		}

		// put them together to form a two digit number
		intString := string(firstInt) + string(lastInt)
		intNumber, err := strconv.Atoi(intString)
		if err != nil {
			panic(err)
		}

		println(intNumber)

		// add the two digit number to a sum
		sum += intNumber
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// print the sum
	println(sum)
}

func part2() {

}

func main() {
	part1()
	part2()
}
