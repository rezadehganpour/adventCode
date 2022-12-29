package problems

import (
	"advent_code_2022/main/commonLib"
	"advent_code_2022/main/model"
)

func RunProgramFive() string {
	cargo, err := commonLib.ReadDayFiveInput("./inputFiles/problem_5.txt")

	commonLib.CheckForError(err)
	var stacks [][]string = cargo.Crates
	var instructions []model.CargoMove = cargo.Instructions
	for _, instruction := range instructions {
		from := instruction.From - 1
		to := instruction.To - 1
		startOfArray := (len(stacks[from])) - instruction.Qun
		newCargo := stacks[from][startOfArray:]
		for i := 0; i < instruction.Qun; i++ {
			cargoDest := stacks[to]
			cargoDest = append(cargoDest, newCargo[i])
			cargoStart := stacks[from]
			cargoStart = cargoStart[:len(cargoStart)-1]
			stacks[from] = cargoStart
			stacks[to] = cargoDest
		}
	}
	var result string
	for _, stack := range stacks {
		length := len(stack)
		if length != 0 {
			value := stack[length-1]
			result += value
		}
	}
	return result
}
