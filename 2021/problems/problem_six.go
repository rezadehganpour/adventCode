package problems

import (
	"advent-code-2021/commonLib"
	"strconv"
)

func RunProgramSix() int {
	commonMap := make(map[int][]int)
	matrix, _ := commonLib.ReadDayThreeInput("./inputFiles/ProblemFive.txt")
	var index int = 0
	var solutionIsFound bool = false
	var prevCommonVal int = -1
	var resultIndex int = -1
	for solutionIsFound != true {
		if len(commonMap[0]) == 1 || len(commonMap[1]) == 1 {
			if len(commonMap[0]) == 1 {
				resultIndex = commonMap[0][0]
			} else {
				resultIndex = commonMap[1][0]
			}
			solutionIsFound = true
		} else {
			var row int = 0
			if index == 0 {
				for row < len(matrix) {
					var tempVal string = matrix[row][index]
					if tempVal == "1" {
						commonMap[1] = append(commonMap[1], row)
					} else {
						commonMap[0] = append(commonMap[0], row)
					}
					row++
				}
			} else {
				var currentList = commonMap[prevCommonVal]
				newTemp := make([]int, 0)
				commonMap[prevCommonVal] = newTemp
				for _, val := range currentList {
					var tempVal string = matrix[val][index]
					if tempVal == "1" {
						commonMap[1] = append(commonMap[1], val)
					} else {
						commonMap[0] = append(commonMap[0], val)
					}
				}
			}
			newTemp := make([]int, 0)
			if len(commonMap[0]) > len(commonMap[1]) {
				commonMap[1] = newTemp
				prevCommonVal = 0
			} else if len(commonMap[0]) < len(commonMap[1]) {
				commonMap[0] = newTemp
				prevCommonVal = 1
			} else {
				commonMap[0] = newTemp
				prevCommonVal = 1
			}
		}
		index++
	}
	var oxygen string = ""
	for _, val := range matrix[resultIndex] {
		oxygen += val
	}
	solutionIsFound = false
	index = 0
	commonMap = make(map[int][]int)
	for solutionIsFound != true {
		if len(commonMap[0]) == 1 || len(commonMap[1]) == 1 {
			if len(commonMap[0]) == 1 {
				resultIndex = commonMap[0][0]
			} else {
				resultIndex = commonMap[1][0]
			}
			solutionIsFound = true
		} else {
			var row int = 0
			if index == 0 {
				for row < len(matrix) {
					var tempVal string = matrix[row][index]
					if tempVal == "1" {
						commonMap[1] = append(commonMap[1], row)
					} else {
						commonMap[0] = append(commonMap[0], row)
					}
					row++
				}
			} else {
				var currentList = commonMap[prevCommonVal]
				newTemp := make([]int, 0)
				commonMap[prevCommonVal] = newTemp
				for _, val := range currentList {
					var tempVal string = matrix[val][index]
					if tempVal == "1" {
						commonMap[1] = append(commonMap[1], val)
					} else {
						commonMap[0] = append(commonMap[0], val)
					}
				}
			}
			newTemp := make([]int, 0)
			if len(commonMap[0]) > len(commonMap[1]) {
				commonMap[0] = newTemp
				prevCommonVal = 1
			} else if len(commonMap[0]) < len(commonMap[1]) {
				commonMap[1] = newTemp
				prevCommonVal = 0
			} else {
				commonMap[1] = newTemp
				prevCommonVal = 0
			}
		}
		index++
	}
	var co2 string = ""
	for _, val := range matrix[resultIndex] {
		co2 += val
	}
	oxygenRes, err := strconv.ParseInt(oxygen, 2, 64)
	commonLib.CheckForError(err)
	co2Res, err := strconv.ParseInt(co2, 2, 64)
	commonLib.CheckForError(err)
	return int(oxygenRes) * int(co2Res)
}
