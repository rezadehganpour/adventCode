package main

import (
	"advent-code-2021/problems"
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
	case *problem == "2":
		printAnswer(problems.RunProgramTwo())
	default:
		fmt.Println("Not problem was run")
	}
}

func printAnswer(result int) {
	fmt.Println("Answer: ", result)
}
