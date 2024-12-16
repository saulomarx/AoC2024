package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part01(input []string) {
	var sum int
	left := make([]int, len(input))
	right := make([]int, len(input))
	for idx, in := range input {

		l, r := getNumbersfromImput(in)
		left[idx] = l
		right[idx] = r
	}
	slices.Sort(left)
	slices.Sort(right)

	for idx, l := range left {
		r := right[idx]
		sum += getDiff(l, r)
	}

	fmt.Printf("Part 01: %d\n", sum)

}

func Part02(input []string) {
	left := make([]int, len(input))
	right := make([]int, len(input))

	commandsMap := make(map[int]int)
	for idx, in := range input {
		l, r := getNumbersfromImput(in)

		_, ok := commandsMap[l]
		if !ok {
			commandsMap[l] = 0
		}
		left[idx] = l
		right[idx] = r
	}
	slices.Sort(left)
	slices.Sort(right)

	for _, r := range right {
		_, ok := commandsMap[r]
		if ok {
			commandsMap[r] += 1
		}

	}
	sum := sumMap(commandsMap, left)
	fmt.Printf("Part 02: %d\n", sum)

}

func getNumbersfromImput(in string) (left, right int) {
	s := strings.Split(in, " ")
	left, _ = strconv.Atoi(s[0])
	right, _ = strconv.Atoi(s[1])
	return
}

func getDiff(a, b int) int {
	result := a - b
	if result < 0 {
		return result * (-1)
	}
	return result
}

func sumMap(m map[int]int, l []int) int {
	var sum = 0
	for _, value := range l {
		sum += value * m[value]
		// fmt.Printf("%d * %d = %d\n", value, m[value], m[value]*value)

	}
	return sum
}
