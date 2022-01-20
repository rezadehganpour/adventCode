package problems

import (
	"advent-code-2021/commonLib"
)

func RunProgramTwo() int {
	dat, err := commonLib.ReadInput("./inputFiles/ProblemTwo.txt")
	commonLib.CheckForError(err)
	result := 0
	currVal := dat[0] + dat[1] + dat[2]
	for index := 1; index < len(dat)-2; index++ {
		newValue := dat[index] + dat[index+1] + dat[index+2]
		if newValue > currVal {
			result++
		}
		currVal = newValue
	}
	return result
}
