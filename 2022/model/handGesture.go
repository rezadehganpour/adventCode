package model

type HandGesture interface {
	PlayGame(HandGesture) int
	GetType() string
	GiveWinner() HandGesture
	GiveLooser() HandGesture
	GiveDraw() HandGesture
}
