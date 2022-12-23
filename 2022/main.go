package main

import (
	"advent_code_2022/main/problems"
	"flag"
	"fmt"
)

func main() {
	problem := flag.String("problem", "no problem was selected", "problem number")
	flag.Parse()
	fmt.Println("Runnung problem:", *problem)
	switch {
	case *problem == "1":
		printAnswer(problems.RunProgramOne())
	default:
		fmt.Println("Not problem was run")
	}
}

func printAnswer(result int) {
	fmt.Println("Answer: ", result)
}
