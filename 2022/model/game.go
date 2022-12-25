package model

type Game struct {
	OpponentHand HandGesture
	MyHand       HandGesture
}

func (g Game) GiveMeScore() int {
	return g.MyHand.PlayGame(g.OpponentHand)
}
