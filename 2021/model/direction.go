package model

type Direction int

const (
	Undefined Direction = iota
	Forward
	Down
	Up
)

type Move struct {
	SubDirection  Direction
	NumberOfMoves int
}

func GetDirection(d string) Direction {
	switch d {
	case "forward":
		return Forward
	case "down":
		return Down
	case "up":
		return Up
	default:
		return Undefined
	}
}
