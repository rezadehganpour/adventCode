package problems

import (
	"advent-code-2021/commonLib"
	"strconv"
)

func RunProgramSeven() int {
	drawings := commonLib.ReadDayFourDrawing("./inputFiles/ProblemSix.txt")
	players := commonLib.ReadDayFourBoard("./inputFiles/ProblemSix.txt")
	var winnerBoardNum int
	var lastNumberDraws int
	var bingo bool = false
	for _, draw := range drawings {
		for index := range players {
			number, exist := players[index].Numbers[draw]
			if exist {
				players[index].MatchNumbers[draw] = true
				players[index].CompletedRow[number.Row]++
				players[index].CompletedCol[number.Col]++
				if players[index].CompletedRow[number.Row] == 5 || players[index].CompletedCol[number.Col] == 5 {
					winnerBoardNum = index
					lastNumberDraws, _ = strconv.Atoi(draw)
					bingo = true
					break
				}
			}
		}
		if bingo {
			break
		}
	}
	var totalUnMatch int
	for row := range players[winnerBoardNum].Board {
		for col := range players[winnerBoardNum].Board[row] {
			_, exist := players[winnerBoardNum].MatchNumbers[players[winnerBoardNum].Board[row][col]]
			if !exist {
				newUnmatch, _ := strconv.Atoi(players[winnerBoardNum].Board[row][col])
				totalUnMatch += newUnmatch
			}
		}
	}
	return totalUnMatch * lastNumberDraws
}
