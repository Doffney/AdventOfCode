package main

import (
	"slices"
	"strconv"
	"strings"
	"unicode"

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
	sum := 0
	if part2 {

		for currline, line := range lines {
			for index, sign := range line {
				if sign == '*' {
					sum += gearChecker3000(lines, index, currline)
				}
			}
		}
		return sum
	}
	for currline, line := range lines {
		number := ""
		indexbegin := 0
		indexend := 0
		endofline := false
		for i, char := range line {
			if unicode.IsNumber(char) {
				number = number + string(char)
				if indexbegin == 0 {
					indexbegin = i
				}
				if i == 139 {
					indexend = i
					endofline = true
					if checkForSymbol(lines, indexbegin, indexend, currline) {
						intnumber, _ := strconv.Atoi(number)
						sum += intnumber
					}
				}
			} else {
				if number != "" && !endofline {
					indexend = i
					if checkForSymbol(lines, indexbegin, indexend, currline) {
						intnumber, _ := strconv.Atoi(number)
						sum += intnumber
					}
				}
				number = ""
				indexbegin = 0
				indexend = 0
			}

		}
	}
	return sum
}

func checkForSymbol(lines []string, nobegin int, noend int, currline int) bool {
	if currline != 0 {
		prevline := lines[currline-1]
		rangetocheck := ""
		if noend == 139 { //endoflinecheck
			rangetocheck = prevline[nobegin-1 : noend]
		} else {
			rangetocheck = prevline[nobegin-1 : noend+1]
		}
		for _, char := range rangetocheck {
			if char != '.' && !unicode.IsNumber(char) {
				return true
			}
		}
	}
	actualline := lines[currline]
	currlinerangetocheck := ""
	if noend == 139 { //endoflinecheck
		currlinerangetocheck = actualline[nobegin-1 : noend]
	} else {
		currlinerangetocheck = actualline[nobegin-1 : noend+1]
	}
	for _, char := range currlinerangetocheck {
		if char != '.' && !unicode.IsNumber(char) {
			return true
		}
	}
	if currline != 139 {
		forwline := lines[currline+1]
		rangetocheck := ""
		if noend == 139 { //endoflinecheck
			rangetocheck = forwline[nobegin-1 : noend]
		} else {
			rangetocheck = forwline[nobegin-1 : noend+1]
		}
		for _, char := range rangetocheck {
			if char != '.' && !unicode.IsNumber(char) {
				return true
			}
		}
	}
	return false
}

func numbercheck666(line string, posfirstnumber int) string {
	number := string(line[posfirstnumber])
	i := 1
	for i <= 2 {
		if unicode.IsNumber(rune(line[posfirstnumber-i])) {
			number = string(line[posfirstnumber-i]) + number
		} else {
			break
		}
		i++
	}
	n := 1
	for n <= 2 {
		if unicode.IsNumber(rune(line[posfirstnumber+n])) {
			number = number + string(line[posfirstnumber+n])
		} else {
			break
		}
		n++
	}
	// fmt.Println(number)
	return number
}

func gearChecker3000(lines []string, gearposition int, currline int) int {
	numbercounter := 0
	numbers := []string{"", "", ""}
	prevline := lines[currline-1]
	topofgear := prevline[gearposition-1 : gearposition+2]
	for numberpos, letter := range topofgear {
		if unicode.IsNumber(letter) {
			foundnumber := numbercheck666(prevline, gearposition-1+numberpos)
			if !slices.Contains(numbers, foundnumber) {
				numbers[numbercounter] = foundnumber
				numbercounter++
				if numbercounter == 3 {
					return 0
				}
			}
		}
	}
	nexttogear := lines[currline][gearposition-1 : gearposition+2]
	for index, char := range nexttogear {
		if unicode.IsNumber(char) {
			foundnumber := numbercheck666(lines[currline], gearposition-1+index)
			if !slices.Contains(numbers, foundnumber) {
				numbers[numbercounter] = foundnumber
				numbercounter++
				if numbercounter == 3 {
					return 0
				}
			}
		}
	}

	futline := lines[currline+1]
	bottomofgear := futline[gearposition-1 : gearposition+2]
	for numberpos, letter := range bottomofgear {
		if unicode.IsNumber(letter) {
			foundnumber := numbercheck666(futline, gearposition-1+numberpos)
			if !slices.Contains(numbers, foundnumber) {
				numbers[numbercounter] = foundnumber
				numbercounter++
				if numbercounter == 3 {
					return 0
				}
			}
		}
	}
	number1, _ := strconv.Atoi(numbers[0])
	// fmt.Print(numbers[0] + "*")
	number2, _ := strconv.Atoi(numbers[1])
	// fmt.Print(numbers[1] + ":")
	//fmt.Println(number1 * number2)
	// fmt.Println()
	return number1 * number2
}
