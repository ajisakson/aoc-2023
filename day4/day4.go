package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {
	sum := 0

	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cardVal := 0
		text := scanner.Text()
		winningNumbers := getWinningNumbers(text)
		myNumbers := getMyNumbers(text)

		for _, winNumber := range winningNumbers {
			for j, myNumber := range myNumbers {
				if winNumber == myNumber {
					fmt.Println(winNumber)
					if cardVal == 0 {
						cardVal = 1
					} else {
						cardVal *= 2
					}

					myNumbers = append(myNumbers[:j], myNumbers[j+1:]...)

					break
				}
			}
		}
		println(cardVal)
		sum += cardVal
	}

	fmt.Println(sum)
}

func part2() {
	sum := 0

	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	countMap := make(map[int]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cardVal := 0
		text := scanner.Text()
		cardNum := getCardNum(text)
		winningNumbers := getWinningNumbers(text)
		myNumbers := getMyNumbers(text)
		for _, winNumber := range winningNumbers {
			for j, myNumber := range myNumbers {
				if winNumber == myNumber {
					cardVal++

					myNumbers = append(myNumbers[:j], myNumbers[j+1:]...)

					break
				}
			}
		}

		copies := countMap[cardNum]
		if copies == 0 {
			countMap[cardNum] = 1
			copies = 1
		}

		println(cardNum, cardVal, copies)

		for i := 1; i <= cardVal; i++ {
			if _, exists := countMap[cardNum+i]; exists {
				countMap[cardNum+i] += copies
			} else {
				countMap[cardNum+i] = copies + 1
			}
			val := countMap[cardNum+i]
			println(cardNum+i, val)
		}

		// Find out how many times I have won on my current card
		// Find out how many copies of my current card I have won
		// for each time I have won, add the number of copies of my card to the next card in the sequence
	}

	for _, val := range countMap {
		sum += val
	}

	fmt.Println(sum)
}

func getCardNum(text string) int {
	firstSplit := strings.Split(text, ": ")
	secondSplit := strings.Fields(firstSplit[0])
	cardNum, _ := strconv.Atoi(secondSplit[1])

	return cardNum
}

func getWinningNumbers(text string) []int {
	var winningNumbers []int

	firstSplit := strings.Split(text, ": ")
	secondSplit := strings.Split(firstSplit[1], " |")
	winningString := secondSplit[0]
	winningSplit := strings.Fields(winningString)
	for _, number := range winningSplit {
		integer, _ := strconv.Atoi(number)
		winningNumbers = append(winningNumbers, integer)
	}

	return winningNumbers
}

func getMyNumbers(text string) []int {
	var myNumbers []int

	firstSplit := strings.Split(text, ": ")
	secondSplit := strings.Split(firstSplit[1], " |")
	myString := secondSplit[1]
	mySplit := strings.Fields(myString)
	for _, number := range mySplit {
		integer, _ := strconv.Atoi(number)
		myNumbers = append(myNumbers, integer)
	}

	return myNumbers
}
