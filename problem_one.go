package main

import (
	"fmt"

	"advent-code-2021/commonLib"
)

func main() {
	dat, err := commonLib.ReadInput("./inputFiles/ProblemOne.txt")
	commonLib.CheckForError(err)
	result := 0
	currentVal := dat[0]
	for index := 1; index < len(dat); index++ {
		if dat[index] > currentVal {
			result++
		}
		currentVal = dat[index]
	}
	fmt.Println(result)
}
