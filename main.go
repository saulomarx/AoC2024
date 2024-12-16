package main

import (
	"fmt"

	"github.com/saulomarx/AoC2024/utils"
	"github.com/saulomarx/AocC023/days/day01"
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
	p := "./days/day01/imputs/in01.txt"
	input := utils.ReadLines(p)
	day01.Part01(input)
	day01.Part02(input)
}
