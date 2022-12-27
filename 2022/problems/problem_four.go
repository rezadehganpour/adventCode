package problems

import (
	"advent_code_2022/main/commonLib"
)

func RunProgramFour() int {
	elfsPair, err := commonLib.ReadDayFourInput("./inputFiles/problem_4.txt")
	commonLib.CheckForError(err)
	var result int

	for _, elfPair := range elfsPair {
		if elfPair.CheckForFullConflict() {
			result += 1
		}
	}
	return result
}
