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

func ReadDayFourDrawing(filePath string) []string {
	file, err := os.Open(filePath)
	CheckForError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	result := strings.Split(scanner.Text(), ",")
	return result
}

func ReadDayFourBoard(filePath string) []model.Player {
	file, err := os.Open(filePath)
	CheckForError(err)
	scanner := bufio.NewScanner(file)
	// Skip the initial drawing input
	scanner.Scan()
	scanner.Text()
	scanner.Scan()
	scanner.Text()
	var result []model.Player
	var boardNum int = 0
	var gameRowNum int = 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			gameRowNum = 0
			boardNum++
		} else {
			if gameRowNum == 0 {
				//first row in the board
				var board [5][5]string
				var newBoardCompletedRow [5]int
				var newBoardCompletedCol [5]int
				var newBoardMatchNumber map[string]bool = make(map[string]bool)
				var allNumbers map[string]model.DayFourCoordination = make(map[string]model.DayFourCoordination)
				newPlayer := model.Player{
					Board:        board,
					CompletedRow: newBoardCompletedRow,
					CompletedCol: newBoardCompletedCol,
					Numbers:      allNumbers,
					MatchNumbers: newBoardMatchNumber,
				}
				result = append(result, newPlayer)
			}
			var inputLine []string = strings.Split(scanner.Text(), " ")
			var bingoLine []string
			for _, value := range inputLine {
				if strings.TrimSpace(value) != "" {
					bingoLine = append(bingoLine, value)
				}
			}
			for index, value := range bingoLine {
				result[boardNum].Board[gameRowNum][index] = value
				newCord := model.DayFourCoordination{
					Row: gameRowNum,
					Col: index,
				}
				result[boardNum].Numbers[value] = newCord
			}
			gameRowNum++
		}
	}
	return result
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
