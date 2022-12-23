package problems

import (
	"advent-code-2021/commonLib"
	"advent-code-2021/model"
)

func RunProgramFour() int {
	dat, err := commonLib.ReadDayTwoInput("./inputFiles/ProblemFour.txt")
	commonLib.CheckForError(err)
	var x int = 0
	var y int = 0
	var aim int = 0
	var result int = 0
	for _, val := range dat {
		if val.SubDirection == model.Down {
			aim += val.NumberOfMoves
		} else if val.SubDirection == model.Up {
			aim -= val.NumberOfMoves
		} else {
			x += val.NumberOfMoves
			y += val.NumberOfMoves * aim
		}
	}
	result = x * y
	return result
}
