package problems

import (
	"advent_code_2022/main/commonLib"
	"sort"
)

func RunProgramOne() int {
	calories, error := commonLib.ReadDayOneFirstInput("./inputFiles/problem_1_1.txt")
	commonLib.CheckForError(error)
	var currCalorie int = 0
	var maxCal int = 0
	var totalCal []int
	for _, calorie := range calories {
		if calorie == -1 {
			totalCal = append(totalCal, currCalorie)
			currCalorie = 0
		} else {
			currCalorie += calorie
		}
	}
	sort.Ints(totalCal)
	var calLength int = len(totalCal)
	maxCal = totalCal[calLength-1] + totalCal[calLength-2] + totalCal[calLength-3]
	return maxCal
}
