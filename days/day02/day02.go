package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func Part01(input []string) {
	inMtx := make([][]int, len(input))
	for idx, in := range input {
		line := getNumbersfromImput(in)
		inMtx[idx] = line

	}

	sum := 0
	for _, line := range inMtx {
		if validateSafety(line) {
			sum++
		}

	}

	fmt.Printf("Part 01: %d\n", sum)

}

func Part02(input []string) {
	inMtx := make([][]int, len(input))
	for idx, in := range input {
		line := getNumbersfromImput(in)
		inMtx[idx] = line

	}

	sum := 0
	for _, line := range inMtx {
		if validateSafety(line) {
			sum++
			continue
		}

		valiDampener := false
		for i := 0; i < len(line); i++ {
			newLine := make([]int, len(line))
			copy(newLine, line)
			newLine = append(newLine[:i], newLine[i+1:]...)

			if validateSafety(newLine) {
				valiDampener = true
				break
			}

		}

		if valiDampener {
			sum++
		}

	}

	fmt.Printf("Part 02: %d\n", sum)

}

func getNumbersfromImput(in string) (line []int) {
	s := strings.Split(in, " ")
	for _, v := range s {
		c, _ := strconv.Atoi(v)
		line = append(line, c)
	}

	return
}

func validateSafety(line []int) bool {
	return (isIncrease(line) || isDecrease(line)) && validDiff(line)
}

func validateSafetyDampener(line []int) bool {
	fmt.Println(line)
	return (isIncrease(line) || isDecrease(line)) && validDiff(line)
}

func isIncrease(line []int) bool {
	for idx := 1; idx < len(line); idx++ {
		before := line[idx-1]
		after := line[idx]

		if before > after {
			return false
		}
	}

	return true
}

func isDecrease(line []int) bool {
	for idx := 1; idx < len(line); idx++ {
		before := line[idx-1]
		after := line[idx]
		if before < after {
			return false
		}
	}

	return true
}

func validDiff(line []int) bool {
	for idx := 1; idx < len(line); idx++ {
		diff := line[idx-1] - line[idx]
		if diff < 0 {
			diff = diff * (-1)
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
