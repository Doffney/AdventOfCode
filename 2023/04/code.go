package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	lines := strings.Split(strings.TrimSpace(input), "\n")
	re := regexp.MustCompile("[0-9]+")
	sum := 0
	if part2 {
		count := 0
		wongames := make([]string, 0)
		wongames = append(wongames, lines...) //initialgame
		for len(wongames) > count {
			line := wongames[count]
			gameresult := 0
			cardNo, rawgameinput, _ := strings.Cut(line, ":")
			gamesides := strings.Split(rawgameinput, "|")
			winningnumbers := re.FindAllString(gamesides[0], -1)
			scratchnumbers := re.FindAllString(gamesides[1], -1)
			for _, number := range scratchnumbers {
				if slices.Contains(winningnumbers, number) {
					gameresult++
				}
			}
			if gameresult != 0 {
				GameNo := re.FindAllString(cardNo, 1)
				i := 0
				blub, _ := strconv.Atoi(GameNo[0])
				for i < gameresult {
					wongames = append(wongames, lines[blub+i])
					i++
				}
			}
			count++
		}
		fmt.Println(wongames[len(wongames)-1])
		return count
	}

	for _, line := range lines {
		gameresult := 0
		_, rawgameinput, _ := strings.Cut(line, ":")
		gamesides := strings.Split(rawgameinput, "|")
		winningnumbers := re.FindAllString(gamesides[0], -1)
		scratchnumbers := re.FindAllString(gamesides[1], -1)
		for _, number := range scratchnumbers {
			if slices.Contains(winningnumbers, number) {
				if gameresult == 0 {
					gameresult++
				} else {
					gameresult = gameresult * 2
				}
			}
		}
		sum += gameresult
	}
	return sum
}
