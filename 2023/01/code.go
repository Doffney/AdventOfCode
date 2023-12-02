package main

import (
	"regexp"
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
	result := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if part2 {
		m := make(map[string]string)
		m["one"] = "1"
		m["two"] = "2"
		m["three"] = "3"
		m["four"] = "4"
		m["five"] = "5"
		m["six"] = "6"
		m["seven"] = "7"
		m["eight"] = "8"
		m["nine"] = "9"
		edgecasemap := make(map[string]string)
		edgecasemap["twone"] = "twoone"
		edgecasemap["oneight"] = "oneeight"
		edgecasemap["eightwo"] = "eighttwo"
		for key, value := range edgecasemap {
			input = strings.ReplaceAll(input, key, value)
		}
		lines = strings.Split(strings.TrimSpace(input), "\n")
		for _, line := range lines {
			re := regexp.MustCompile("(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|[0-9]")
			numbers := re.FindAllString(line, -1)
			for i, number := range numbers {
				if len(number) != 1 {
					numbers[i] = m[number]
				}
			}

			result += putEmTogether(numbers)
		}
		return result

	}

	for _, line := range lines {
		re := regexp.MustCompile("[0-9]")
		numbers := re.FindAllString(line, -1)
		result += putEmTogether(numbers)
	}
	return result

}

func putEmTogether(numbers []string) int {
	if len(numbers) == 1 {
		number, _ := strconv.Atoi(numbers[0])
		bignumber := number * 10
		return bignumber + number
	} else {
		bignumber, _ := strconv.Atoi(numbers[0])
		number, _ := strconv.Atoi(numbers[len(numbers)-1])
		return bignumber*10 + number
	}
}
