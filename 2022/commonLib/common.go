package commonLib

import (
	"bufio"
	"os"
	"strconv"
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

func CheckForError(e error) {
	if e != nil {
		panic(e)
	}
}
