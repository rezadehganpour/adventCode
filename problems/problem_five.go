package problems

import (
	"advent-code-2021/commonLib"
	"strconv"
	"strings"
)

func RunProgramFive() int {
	matrix, _ := commonLib.ReadDayThreeInput("./inputFiles/ProblemFive.txt")
	var row int = 0
	var col int = 0
	freqNumsGama := ""
	for col < len(matrix[0]) {
		numOne := 0
		numZero := 0
		row = 0
		for row < len(matrix) {
			if matrix[row][col] == "1" {
				numOne++
			} else {
				numZero++
			}
			row += 1
		}
		if numOne > numZero {
			freqNumsGama += "1"
		} else {
			freqNumsGama += "0"
		}
		col += 1
	}
	gamma, err := strconv.ParseInt(freqNumsGama, 2, 64)
	freqNumsEplison := ""
	commonLib.CheckForError(err)
	for _, bin := range strings.Split(freqNumsGama, "") {
		if bin == "1" {
			freqNumsEplison += "0"
		} else {
			freqNumsEplison += "1"
		}
	}
	epsilon, err := strconv.ParseInt(freqNumsEplison, 2, 64)
	return int(gamma) * int(epsilon)
}
