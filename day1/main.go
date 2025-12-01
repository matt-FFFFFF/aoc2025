package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	dialLength = 100
)

func motion2int(motion string) int {
	direction := []rune(motion)[0]

	number, err := strconv.Atoi(string([]rune(motion)[1:]))
	if err != nil {
		panic("Cannot convert motion to int")
	}

	switch direction {
	case 'R':
		return number
	case 'L':
		return -number
	default:
		return 0
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: program motions.txt")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close() //nolint:errcheck

	scanner := bufio.NewScanner(file)
	position := 50

	var part1, part2 int

	for scanner.Scan() {
		motion := motion2int(scanner.Text())
		
		// Count how many times we cross 0 (including landing on it)
		if motion > 0 {
			// Moving right: count how many times we pass 0
			for i := 1; i <= motion; i++ {
				if (position + i) % dialLength == 0 {
					part2++
				}
			}
		} else if motion < 0 {
			// Moving left: count how many times we pass 0
			for i := 1; i <= -motion; i++ {
				if (position - i + dialLength) % dialLength == 0 {
					part2++
				}
			}
		}
		
		// Update position
		position = (position + motion % dialLength + dialLength) % dialLength
		
		// Part 1: count when dial ends at 0
		if position == 0 {
			part1++
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
