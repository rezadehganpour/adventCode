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
	case *problem == "2":
		printAnswer(problems.RunProgramTwo())
	case *problem == "3":
		printAnswer(problems.RunProgramThree())
	case *problem == "4":
		printAnswer(problems.RunProgramFour())
	case *problem == "5":
		printAnswerInString(problems.RunProgramFive())
	case *problem == "6":
		printAnswer(problems.RunProgramSix())
	default:
		fmt.Println("Not problem was run")
	}
}

func printAnswerInString(result string) {
	fmt.Println("Answer: ", result)
}

func printAnswer(result int) {
	fmt.Println("Answer: ", result)
}
