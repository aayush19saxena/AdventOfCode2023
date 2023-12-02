package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f2("input_puzzle2.txt")
}

type cubeSet struct {
	id    int
	red   int
	green int
	blue  int
}

func readLines(filePath string) []string {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func strip(s string, tokens []string) string {
	stripped := s
	for _, t := range tokens {
		stripped = strings.ReplaceAll(stripped, t, "")
	}
	return stripped
}

func lineToCubes(line string) cubeSet {
	var game cubeSet

	s := strip(line, []string{":", ",", ";"})
	split := strings.Split(s, " ")

	game.id, _ = strconv.Atoi(split[1])
	for i := 2; i < len(split)-1; i += 2 {
		val, _ := strconv.Atoi(split[i])
		color := split[i+1]

		switch color {
		case "blue":
			game.blue = max(game.blue, val)
		case "green":
			game.green = max(game.green, val)
		case "red":
			game.red = max(game.red, val)
		}
	}
	return game
}

func isGameWithin(game cubeSet, within cubeSet) bool {
	if game.blue <= within.blue && game.green <= within.green && game.red <= within.red {
		return true
	}
	return false
}

func f2(file string) (int, int) {

	within := cubeSet{0, 12, 13, 14}
	p1, p2 := 0, 0
	data := readLines(file)
	for _, line := range data {
		cubeSet := lineToCubes(line)
		if ok := isGameWithin(cubeSet, within); ok {
			p1 += cubeSet.id
		}
		p2 += cubeSet.red * cubeSet.blue * cubeSet.green
	}

	fmt.Printf("Part1: %d, Part2: %d\n", p1, p2)
	return p1, p2
}
