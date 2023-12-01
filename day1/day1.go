package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var NumMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

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

		// add the two digit number to a sum
		sum += intNumber
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	println(sum)
}

func part2() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		count++
		text := scanner.Text()

		firstInt := -1
		lastInt := -1
		// parse the string for all instances of integers or words in the numMap
		tokens := parseLinesIntoTokens(text)
		for _, str := range tokens {
			if firstInt == -1 {
				if isStringSingleDigitInteger(str) {
					firstInt, err = strconv.Atoi(str)
					if err != nil {
						panic(err)
					}
				} else {
					firstInt = parseStringForFirstInt(str)
				}
			}

			if isStringSingleDigitInteger(str) {
				lastInt, err = strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
			} else {
				parsedLastInt := parseStringForLastInt(str)
				if parsedLastInt != -1 {
					lastInt = parsedLastInt
				}
			}
		}

		strNum := strconv.Itoa(firstInt) + strconv.Itoa(lastInt)
		strNumInt, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}

		sum += strNumInt
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	println(sum)
}

func main() {
	// part1()
	part2()
}

func parseStringForFirstInt(text string) int {
	for i := 0; i < len(text); i++ {
		word := fmt.Sprintf("%c", text[i])
		for j := i + 1; j < len(text); j++ {
			word = word + fmt.Sprintf("%c", text[j])
			for str, strint := range NumMap {
				if word == str {
					return strint
				}
			}
		}
	}

	return -1
}

func parseStringForLastInt(text string) int {
	for i := len(text) - 1; i > 0; i-- {
		word := fmt.Sprintf("%c", text[i])
		for j := i - 1; j >= 0; j-- {
			word = fmt.Sprintf("%c", text[j]) + word
			for str, strint := range NumMap {
				if word == str {

					return strint
				}
			}
		}
	}

	return -1
}

func parseLinesIntoTokens(text string) []string {
	tokens := []string{}
	word := ""
	for i := 0; i < len(text); i++ {
		char := text[i]
		if isByteDigit(char) {
			if word != "" {
				tokens = append(tokens, word)
				word = ""
			}
			tokens = append(tokens, string(char))
		} else {
			word += string(char)
		}
	}
	if word != "" {
		tokens = append(tokens, word)
	}

	return tokens
}

func isByteDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isStringSingleDigitInteger(text string) bool {
	if len(text) != 1 {
		return false
	}
	_, err := strconv.Atoi(text)
	return err == nil
}
