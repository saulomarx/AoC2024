package main

import (
	"fmt"

	"github.com/saulomarx/AoC2024/days/day01"
	"github.com/saulomarx/AoC2024/utils"
)

func main() {
	workingDay := 1

	fmt.Printf("Day %d\n", workingDay)
	switch workingDay {
	case 1:
		runsDay01()
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
