package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

const rcubemax = 12
const gcubemax = 13
const bcubemax = 14

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if part2 {
		for _, game := range lines {
			sum += number(game)
		}
		return sum
	}
	for i, game := range lines {
		valido := true
		gamesplit := strings.Split(game, ":")
		re := regexp.MustCompile("[,;]")
		cubes := re.Split(strings.TrimSpace(gamesplit[1]), -1)
		for _, cube := range cubes {
			if !validate(cube) {
				valido = false
			}
		}
		if valido {
			sum += i + 1
		}
	}
	return sum
}

func number(game string) int {
	rmax := 0
	gmax := 0
	bmax := 0
	gamesplit := strings.Split(game, ":")
	re := regexp.MustCompile("[,;]")
	cubes := re.Split(strings.TrimSpace(gamesplit[1]), -1)
	for _, cube := range cubes {
		numberandcube := strings.Split(strings.TrimSpace(cube), " ")
		numberofcubes, _ := strconv.Atoi(numberandcube[0])
		switch numberandcube[1] {
		case "red":
			if numberofcubes > rmax {
				rmax = numberofcubes
			}
		case "green":
			if numberofcubes > gmax {
				gmax = numberofcubes
			}
		case "blue":
			if numberofcubes > bmax {
				bmax = numberofcubes
			}
		}
	}
	// fmt.Println(rmax * gmax * bmax)
	return rmax * gmax * bmax
}

func validate(cube string) bool {
	numberandcube := strings.Split(strings.TrimSpace(cube), " ")
	numberofcubes, _ := strconv.Atoi(numberandcube[0])
	switch numberandcube[1] {
	case "red":
		return numberofcubes <= rcubemax
	case "green":
		return numberofcubes <= gcubemax
	case "blue":
		return numberofcubes <= bcubemax
	}
	return false
}
