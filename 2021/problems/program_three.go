package problems

import (
	"advent-code-2021/commonLib"
	"advent-code-2021/model"
)

func RunProgramThree() int {
	dat, err := commonLib.ReadDayTwoInput("./inputFiles/ProblemThree.txt")
	commonLib.CheckForError(err)
	var x int = 0
	var y int = 0
	var result int = 0
	for _, val := range dat {
		if val.SubDirection == model.Down {
			y += val.NumberOfMoves
		} else if val.SubDirection == model.Up {
			y -= val.NumberOfMoves
		} else {
			x += val.NumberOfMoves
		}
	}
	result = x * y
	return result
}
