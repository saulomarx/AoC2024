package main

import (
	"fmt"

	"github.com/saulomarx/AoC2024/days/day01"
	"github.com/saulomarx/AoC2024/days/day02"
	"github.com/saulomarx/AoC2024/days/day03"
	"github.com/saulomarx/AoC2024/days/day04"
	"github.com/saulomarx/AoC2024/days/day05"

	"github.com/saulomarx/AoC2024/utils"
)

func main() {
	workingDay := 1

	fmt.Printf("Day %d\n", workingDay)
	switch workingDay {
	case 1:
		runsDay01()
	case 2:
		runsDay02()
	case 3:
		runsDay03()
	case 4:
		runsDay04()
	case 5:
		runsDay05()
	default:
		fmt.Println("not today")
	}

}

func runsDay01() {
	day := "day01"
	rootPath := fmt.Sprintf("./days/%s/inputs/", day)
	samplePath := fmt.Sprintf("%s%s", rootPath, "example01.txt")
	p := fmt.Sprintf("%s%s", rootPath, "in01.txt")
	fmt.Println(samplePath)
	fmt.Println(p)
	input := utils.ReadLines(p)
	day01.Part01(input)
	day01.Part02(input)
}

func runsDay02() {
	day := "day02"
	rootPath := fmt.Sprintf("./days/%s/inputs/", day)
	samplePath := fmt.Sprintf("%s%s", rootPath, "example01.txt")
	p := fmt.Sprintf("%s%s", rootPath, "in01.txt")
	fmt.Println(samplePath)
	fmt.Println(p)
	input := utils.ReadLines(samplePath)
	day02.Part01(input)
	day02.Part02(input)
}

func runsDay03() {
	day := "day03"
	rootPath := fmt.Sprintf("./days/%s/inputs/", day)
	samplePath := fmt.Sprintf("%s%s", rootPath, "example01.txt")
	p := fmt.Sprintf("%s%s", rootPath, "in01.txt")
	fmt.Println(samplePath)
	fmt.Println(p)
	input := utils.ReadLines(p)
	day03.Part01(input)
	day03.Part02(input)
}

func runsDay04() {
	day := "day04"
	rootPath := fmt.Sprintf("./days/%s/inputs/", day)
	samplePath := fmt.Sprintf("%s%s", rootPath, "example01.txt")
	p := fmt.Sprintf("%s%s", rootPath, "in01.txt")
	fmt.Println(samplePath)
	fmt.Println(p)

	input := utils.ReadLines(p)
	mtx := utils.ReadStrMatrix(input)
	day04.Part01(mtx)
	day04.Part02(mtx)
}

func runsDay05() {
	day := "day05"
	rootPath := fmt.Sprintf("./days/%s/inputs/", day)
	samplePath := fmt.Sprintf("%s%s", rootPath, "example01.txt")
	p := fmt.Sprintf("%s%s", rootPath, "in01.txt")
	fmt.Println(samplePath)
	fmt.Println(p)

	input := utils.ReadLines(p)
	mtx := utils.ReadStrMatrix(input)
	day05.Part01(mtx)
	day05.Part02(mtx)
}
