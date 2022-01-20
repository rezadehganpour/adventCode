package commonLib

import (
	"bufio"
	"os"
	"strconv"
)

func CheckForError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput(filePath string) ([]int, error) {
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
