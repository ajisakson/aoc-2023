package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		id, rest := getGameId(text)
		arrGames := strings.Split(rest, "; ")
		valid := true
		for _, str := range arrGames {
			pulls := strings.Split(str, ", ")
			for _, pull := range pulls {
				colors := strings.Split(pull, " ")
				amount, _ := strconv.Atoi(colors[0])
				switch colors[1] {
				case "red":
					if amount > 12 {
						valid = false
					}
				case "green":
					if amount > 13 {
						valid = false
					}
				case "blue":
					if amount > 14 {
						valid = false
					}
				}
			}
		}

		if valid {
			sum += id
		}
	}

	println(sum)
}

func getGameId(text string) (int, string) {
	str := strings.Split(text, ": ")
	id := strings.Split(str[0], "Game ")
	returnId, _ := strconv.Atoi(id[1])
	return returnId, str[1]
}
