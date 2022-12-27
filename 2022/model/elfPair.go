package model

type ElfPair struct {
	First  Coordiantion
	Second Coordiantion
}

func (elfPair ElfPair) CheckForFullConflict() bool {
	if elfPair.First.Begin >= elfPair.Second.Begin && elfPair.Second.End >= elfPair.First.Begin {
		return true
	} else if elfPair.Second.Begin >= elfPair.First.Begin && elfPair.First.End >= elfPair.Second.Begin {
		return true
	}
	return false
}
