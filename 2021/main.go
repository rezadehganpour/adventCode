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
	case *problem == "3":
		printAnswer(problems.RunProgramThree())
	case *problem == "4":
		printAnswer(problems.RunProgramFour())
	case *problem == "5":
		printAnswer(problems.RunProgramFive())
	case *problem == "6":
		printAnswer(problems.RunProgramSix())
	case *problem == "7":
		printAnswer(problems.RunProgramSeven())
	case *problem == "8":
		printAnswer(problems.RunProgramEight())
	default:
		fmt.Println("Not problem was run")
	}
}

func printAnswer(result int) {
	fmt.Println("Answer: ", result)
}
