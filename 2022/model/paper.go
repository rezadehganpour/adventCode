package model

type Paper struct {
}

const paperWinAmount int = 2

func (p Paper) PlayGame(gesture HandGesture) int {
	var result int = paperWinAmount
	if gesture.GetType() == "R" {
		result += 6
	} else if gesture.GetType() == "S" {
		result += 0
	} else if gesture.GetType() == "P" {
		result += 3
	}
	return result
}

func (p Paper) GetType() string {
	return "P"
}

func (p Paper) GiveWinner() HandGesture {
	return Rock{}
}
func (p Paper) GiveLooser() HandGesture {
	return Scissor{}
}
func (p Paper) GiveDraw() HandGesture {
	return Paper{}
}
