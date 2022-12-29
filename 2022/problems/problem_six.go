package problems

import (
	"advent_code_2022/main/commonLib"
)

func RunProgramSix() int {
	stream := commonLib.ReadDaySixInput("./inputFiles/problem_6.txt")
	var result int
	var index int

	for index < len(stream)-3 {
		var checker = make(map[string]bool)
		isMatch := true
		tailer := index
		for i := 0; i <= 13; i++ {
			if !checker[string(stream[tailer])] {
				checker[string(stream[tailer])] = true
			} else {
				isMatch = false
				break
			}
			tailer += 1
		}
		if isMatch {
			result = index + 14
			break
		}
		index += 1
	}
	return result
}
