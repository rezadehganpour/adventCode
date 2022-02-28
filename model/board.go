package model

type Player struct {
	Board        [5][5]string
	CompletedRow [5]int
	CompletedCol [5]int
	Numbers      map[string]DayFourCoordination
	MatchNumbers map[string]bool
}

type DayFourCoordination struct {
	Row int
	Col int
}
