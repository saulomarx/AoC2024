package main

import (
	"fmt"

	"github.com/saulomarx/AoC2024/days/day01"
	"github.com/saulomarx/AoC2024/days/day02"
	"github.com/saulomarx/AoC2024/days/day03"
	"github.com/saulomarx/AoC2024/utils"
)

func main() {
	workingDay := 3

	fmt.Printf("Day %d\n", workingDay)
	switch workingDay {
	case 1:
		runsDay01()
	case 2:
		runsDay02()
	case 3:
		runsDay03()
	default:
		fmt.Println("not today")
	}

}

func runsDay01() {
	samplePath := "./days/day01/imputs/example01.txt"
	p := "./days/day01/imputs/in01.txt"
	fmt.Println(samplePath)
	fmt.Println(p)
	input := utils.ReadLines(p)
	day01.Part01(input)
	day01.Part02(input)
}

func runsDay02() {
	samplePath := "./days/day03/imputs/example01.txt"
	p := "./days/day03/imputs/in01.txt"
	fmt.Println(samplePath)
	fmt.Println(p)
	input := utils.ReadLines(samplePath)
	day02.Part01(input)
	day02.Part02(input)
}

func runsDay03() {
	samplePath := "./days/day03/imputs/example01.txt"
	samplePath = "./days/day03/imputs/example02.txt"
	p := "./days/day03/imputs/in02.txt"
	fmt.Println(samplePath)
	fmt.Println(p)
	input := utils.ReadLines(p)
	day03.Part01(input)
	day03.Part02(input)
}
