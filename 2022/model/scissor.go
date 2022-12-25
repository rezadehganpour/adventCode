package model

type Scissor struct {
}

const scissorWinPoint int = 3

func (s Scissor) PlayGame(gesture HandGesture) int {
	var result int = scissorWinPoint
	if gesture.GetType() == "R" {
		result += 0
	} else if gesture.GetType() == "S" {
		result += 3
	} else if gesture.GetType() == "P" {
		result += 6
	}
	return result
}

func (s Scissor) GetType() string {
	return "S"
}

func (s Scissor) GiveWinner() HandGesture {
	return Paper{}
}
func (s Scissor) GiveLooser() HandGesture {
	return Rock{}
}
func (s Scissor) GiveDraw() HandGesture {
	return Scissor{}
}
