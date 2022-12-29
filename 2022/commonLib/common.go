package commonLib

import (
	"advent_code_2022/main/model"
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadDayOneFirstInput(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	var result []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newValue := scanner.Text()
		if newValue == "" {
			result = append(result, -1)
		} else {
			newCal, _ := strconv.Atoi(newValue)
			result = append(result, newCal)
		}

	}
	return result, scanner.Err()
}

func calculateGesture(gesture string) model.HandGesture {
	var newHandGesture model.HandGesture
	if gesture == "A" || gesture == "X" {
		newHandGesture = model.Rock{}
	} else if gesture == "B" || gesture == "Y" {
		newHandGesture = model.Paper{}
	} else if gesture == "C" || gesture == "Z" {
		newHandGesture = model.Scissor{}
	}
	return newHandGesture
}
func calculateEndGesture(oppoGesture model.HandGesture, result string) model.HandGesture {
	var myHand model.HandGesture
	myHand = calculateMyHand(result, oppoGesture)
	return myHand
}

func calculateMyHand(endResult string, oppoHand model.HandGesture) model.HandGesture {
	var result model.HandGesture
	if endResult == "X" {
		result = oppoHand.GiveWinner()
	} else if endResult == "Y" {
		result = oppoHand.GiveDraw()
	} else if endResult == "Z" {
		result = oppoHand.GiveLooser()
	}
	return result
}
func ReadDayTwoInput(filePath string) ([]model.Game, error) {
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	var result []model.Game
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newValue := scanner.Text()
		newValues := strings.Split(newValue, " ")
		var opponent model.HandGesture = calculateGesture(newValues[0])
		var myHand model.HandGesture = calculateEndGesture(opponent, newValues[1])
		newGame := model.Game{MyHand: myHand, OpponentHand: opponent}
		result = append(result, newGame)
	}
	return result, err
}

func ReadDayThreeInput(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newValue := scanner.Text()
		result = append(result, newValue)
	}
	return result, err
}

func ReadDayFourInput(filePath string) ([]model.ElfPair, error) {
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	var result []model.ElfPair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newValue := scanner.Text()
		twoPairs := strings.Split(newValue, ",")
		firstPairCord := strings.Split(twoPairs[0], "-")
		firstPairBegin, _ := strconv.Atoi(firstPairCord[0])
		firstPairEnd, _ := strconv.Atoi(firstPairCord[1])
		firstElf := model.Coordiantion{Begin: firstPairBegin, End: firstPairEnd}
		secondPairCord := strings.Split(twoPairs[1], "-")
		secondPairBegin, _ := strconv.Atoi(secondPairCord[0])
		secondPairEnd, _ := strconv.Atoi(secondPairCord[1])
		secondElf := model.Coordiantion{Begin: secondPairBegin, End: secondPairEnd}
		elfs := model.ElfPair{First: firstElf, Second: secondElf}
		result = append(result, elfs)
	}
	return result, err
}

func ReadDayFiveInput(filePath string) (model.Cargo, error) {
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var stacks = make([][]string, 10)

	for scanner.Scan() {
		var stackNum int
		newValue := scanner.Text()
		if !strings.Contains(newValue, "[") {
			break
		}
		for i := 0; i < len(newValue); i += 4 {
			if string(newValue[i]) == "[" {
				var value string = string(newValue[i+1])
				stack := stacks[stackNum]
				stack = append([]string{value}, stack...)
				stacks[stackNum] = stack
			}
			stackNum += 1
		}
	}
	re := regexp.MustCompile("[0-9]+")
	var instructions []model.CargoMove
	for scanner.Scan() {
		allMatches := re.FindAllString(scanner.Text(), -1)
		if len(allMatches) != 0 {
			quantity, _ := strconv.Atoi(allMatches[0])
			from, _ := strconv.Atoi(allMatches[1])
			to, _ := strconv.Atoi(allMatches[2])
			newItem := model.CargoMove{Qun: quantity, To: to, From: from}
			instructions = append(instructions, newItem)
		}
	}
	result := model.Cargo{Crates: stacks, Instructions: instructions}
	return result, err
}

func CheckForError(e error) {
	if e != nil {
		panic(e)
	}
}
