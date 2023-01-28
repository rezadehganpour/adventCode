package problems

import (
	"advent_code_2022/main/commonLib"
	"advent_code_2022/main/model"
	"fmt"
)

func RunProgramSeven() int {
	result := 70000001
	directories := commonLib.ReadDaySevenInput("./inputFiles/problem_7.txt")
	setSizeOfEachDir(directories[0])
	totalAvalSize := 70000000 - directories[0].TotalSize
	for _, dir := range directories {
		if totalAvalSize+dir.TotalSize >= 30000000 {
			if result > dir.TotalSize {
				result = dir.TotalSize
			}
			fmt.Printf("Dir Name: %s, total size: %d\n", dir.Name, dir.TotalSize)
		}
	}
	return result
}

func setSizeOfEachDir(root *model.Directory) {
	root.TotalSize = getSizeDirsRecursively(root)
}

func getSizeDirsRecursively(root *model.Directory) int {
	fmt.Printf("Dir in search %s\n", root.Name)
	var totalFileSize int
	var totalDirsSize int

	for _, file := range root.Files {
		fmt.Printf("Dir %s has this file %s with this size %d\n", root.Name, file.Name, file.Size)
		totalFileSize += file.Size
	}
	// var dirSize int
	for _, dir := range root.Directories {
		totalDirsSize += getSizeDirsRecursively(dir)
		fmt.Printf("Directory %s has this size %d\n", dir.Name, totalDirsSize)
	}
	root.TotalSize = totalDirsSize + totalFileSize
	return totalDirsSize + totalFileSize
}
