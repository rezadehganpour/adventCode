package commonLib

import (
	"advent_code_2022/main/model"
	"bufio"
	"os"
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

func CheckForError(e error) {
	if e != nil {
		panic(e)
	}
}
