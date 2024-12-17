package day03

import (
	"fmt"
	"strconv"
	"unicode"
)

//verifcar o do e dont

func Part01(input []string) {
	sum := 0
	for _, in := range input {
		for i := 0; i < len(in); i++ {
			if in[i] == 'm' {
				s, idx := resolveMult(in, i)
				sum += s
				i = idx
			}
		}
	}

	fmt.Printf("Part 01: %d\n", sum)

}

func Part02(input []string) {
	sum := 0
	isEnable := true

	for _, in := range input {
		for i := 0; i < len(in); i++ {
			if in[i] == 'd' {
				isDo, idx := findDo(in, i)
				if isDo {
					isEnable = true
					i = i + idx
				}

				isDont, idx := findDont(in, i)
				if isDont {
					isEnable = false
					i = i + idx
				}
			}
			if in[i] == 'm' {
				s, idx := resolveMult(in, i)
				if isEnable {
					sum += s
				}
				i = idx
			}
		}
	}

	fmt.Printf("Part 02: %d\n", sum)

}

func resolveMult(s string, idx int) (int, int) {
	if s[idx] != 'm' {
		return 0, idx
	}
	idx++
	if s[idx] != 'u' {
		return 0, idx
	}
	idx++
	if s[idx] != 'l' {
		return 0, idx
	}
	idx++
	if s[idx] != '(' {
		return 0, idx
	}
	idx++

	firstNumber := make([]rune, 0)
	for unicode.IsDigit(rune(s[idx])) {
		firstNumber = append(firstNumber, rune(s[idx]))
		idx++
	}

	if s[idx] != ',' {
		return 0, idx
	}
	idx++

	secondNumber := make([]rune, 0)
	for unicode.IsDigit(rune(s[idx])) {
		secondNumber = append(secondNumber, rune(s[idx]))
		idx++
	}

	if s[idx] != ')' {
		return 0, idx
	}

	a, _ := strconv.Atoi(string(firstNumber))
	b, _ := strconv.Atoi(string(secondNumber))

	return a * b, idx

}

func findDont(s string, idx int) (bool, int) {
	isDont := s[idx:idx+7] == "don't()"
	if isDont {
		return isDont, 7
	}
	return false, 0

}

func findDo(s string, idx int) (bool, int) {
	isDO := s[idx:idx+4] == "do()"
	if isDO {
		return isDO, 4
	}
	return false, 0
}
