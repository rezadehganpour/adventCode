package problems

import (
	"advent_code_2022/main/commonLib"
)

func RunProgramTwo() int {
	games, err := commonLib.ReadDayTwoInput("./inputFiles/problem_2.txt")
	commonLib.CheckForError(err)
	var result int = 0
	for _, game := range games {
		result += game.GiveMeScore()
	}
	return result
}
