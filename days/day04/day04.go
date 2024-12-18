package day04

import (
	"fmt"
)

const XMAS = "XMAS"
const MAS = "MAS"

func Part01(mtx [][]string) {
	count := 0
	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx); j++ {
			if mtx[i][j] == "X" {
				count += getRight(i, j, mtx)
				count += getLeft(i, j, mtx)
				count += getUp(i, j, mtx)
				count += getDown(i, j, mtx)

				count += getRightUp(i, j, mtx)
				count += getRightDown(i, j, mtx)
				count += getLeftUp(i, j, mtx)
				count += getLefttDown(i, j, mtx)
			}
		}
	}
	fmt.Printf("Part 01: %d\n", count)

}

func Part02(mtx [][]string) {
	count := 0
	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx); j++ {
			if mtx[i][j] == "A" {
				count += getMas(i, j, mtx)
			}
		}
	}
	fmt.Printf("Part 02: %d\n", count)
}

func getRight(i, j int, mtx [][]string) int {
	maxLine := len(mtx[i])

	if j+3 >= maxLine {
		return 0
	}

	if isXmas(mtx[i][j], mtx[i][j+1], mtx[i][j+2], mtx[i][j+3]) {
		return 1
	}
	return 0
}

func getLeft(i, j int, mtx [][]string) int {
	if j-3 < 0 {
		return 0
	}

	if isXmas(mtx[i][j], mtx[i][j-1], mtx[i][j-2], mtx[i][j-3]) {
		return 1
	}
	return 0
}

func getDown(i, j int, mtx [][]string) int {
	maxColun := len(mtx)

	if i+3 >= maxColun {
		return 0
	}

	if isXmas(mtx[i][j], mtx[i+1][j], mtx[i+2][j], mtx[i+3][j]) {
		return 1
	}
	return 0
}

func getUp(i, j int, mtx [][]string) int {
	if i-3 < 0 {
		return 0
	}

	if isXmas(mtx[i][j], mtx[i-1][j], mtx[i-2][j], mtx[i-3][j]) {
		return 1
	}
	return 0
}

func getRightDown(i, j int, mtx [][]string) int {
	maxLine := len(mtx[i])
	maxColun := len(mtx)
	if i+3 >= maxLine || j+3 >= maxColun {
		return 0
	}

	if isXmas(mtx[i][j], mtx[i+1][j+1], mtx[i+2][j+2], mtx[i+3][j+3]) {
		return 1
	}
	return 0
}

func getRightUp(i, j int, mtx [][]string) int {
	maxLine := len(mtx[i])
	if i+3 >= maxLine || j-3 < 0 {
		return 0
	}

	if isXmas(mtx[i][j], mtx[i+1][j-1], mtx[i+2][j-2], mtx[i+3][j-3]) {
		return 1
	}
	return 0
}

func getLefttDown(i, j int, mtx [][]string) int {
	maxColun := len(mtx)
	if i-3 < 0 || j+3 >= maxColun {
		return 0
	}

	if isXmas(mtx[i][j], mtx[i-1][j+1], mtx[i-2][j+2], mtx[i-3][j+3]) {
		return 1
	}
	return 0
}

func getLeftUp(i, j int, mtx [][]string) int {
	if i-3 < 0 || j-3 < 0 {
		return 0
	}

	if isXmas(mtx[i][j], mtx[i-1][j-1], mtx[i-2][j-2], mtx[i-3][j-3]) {
		return 1
	}
	return 0
}

func getMas(i, j int, mtx [][]string) int {
	maxLine := len(mtx[i])
	maxColun := len(mtx)

	if i-1 < 0 || j-1 < 0 || i+1 >= maxLine || j+1 >= maxColun {
		return 0
	}

	isMasLeftDown := isMas(mtx[i-1][j-1], mtx[i][j], mtx[i+1][j+1]) || isMas(mtx[i+1][j+1], mtx[i][j], mtx[i-1][j-1])
	isMasRigthDown := isMas(mtx[i-1][j+1], mtx[i][j], mtx[i+1][j-1]) || isMas(mtx[i+1][j-1], mtx[i][j], mtx[i-1][j+1])

	if isMasLeftDown && isMasRigthDown {
		return 1
	}
	return 0
}

func isXmas(a, b, c, d string) bool {
	in := fmt.Sprintf("%s%s%s%s", a, b, c, d)
	return in == XMAS
}

func isMas(a, b, c string) bool {
	in := fmt.Sprintf("%s%s%s", a, b, c)
	return in == MAS
}
