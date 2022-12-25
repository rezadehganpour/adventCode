package model

type Rock struct {
}

const rockWinPoint int = 1

func (r Rock) PlayGame(gesture HandGesture) int {
	var result int = rockWinPoint
	if gesture.GetType() == "R" {
		result += 3
	} else if gesture.GetType() == "S" {
		result += 6
	} else if gesture.GetType() == "P" {
		result += 0
	}
	return result
}

func (r Rock) GetType() string {
	return "R"
}

func (r Rock) GiveWinner() HandGesture {
	return Scissor{}
}
func (r Rock) GiveLooser() HandGesture {
	return Paper{}
}
func (r Rock) GiveDraw() HandGesture {
	return Rock{}
}
