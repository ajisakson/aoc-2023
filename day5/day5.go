package main

import (
	"bufio"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	sum := 0

	f, err := os.Open("test")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		println(text)
	}

	println(sum)

}

func part2() {
}
