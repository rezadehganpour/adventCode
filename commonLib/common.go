package commonLib

import (
	"advent-code-2021/model"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CheckForError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadDayOneInput(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	var result []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newValue, _ := strconv.Atoi(scanner.Text())
		result = append(result, newValue)
	}
	return result, scanner.Err()
}

func ReadDayTwoInput(filePath string) ([]model.Move, error) {
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	var result []model.Move
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		submarineMove := strings.Split(scanner.Text(), " ")
		newDirection := model.GetDirection(submarineMove[0])
		newMoveNum, _ := strconv.Atoi(submarineMove[1])
		var newMove = model.Move{SubDirection: newDirection, NumberOfMoves: newMoveNum}
		result = append(result, newMove)
	}
	return result, scanner.Err()
}

func ReadDayThreeInput(filePath string) ([][12]string, error) {
	var numberOfLines = getNumberOfLinesCount(filePath)
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	var matrix = make([][12]string, numberOfLines)
	scanner := bufio.NewScanner(file)
	var row int = 0
	for scanner.Scan() {
		characters := strings.Split(scanner.Text(), "")
		for col, character := range characters {
			matrix[row][col] = character
		}
		row++
	}
	return matrix, scanner.Err()
}

func getNumberOfLinesCount(filePath string) int {
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	var result int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result++
	}
	return result
}
