package asset

type Attribute interface {
	Score() int
	Modifer() int
	SetScore(ns int)
}

func NewAttribute(name string) Attribute {
	atr := Asset{}
	atr.Group = TypeAttribute
	atr.AssetName = name
	atr.BonusScore = make(map[string]int)
	atr.BonusModifier = make(map[string]int)
	return &atr
}

func (a *Asset) Score() int {
	base := a.Rank
	if a.Rank == 0 {
		return 0
	}
	for _, val := range a.BonusScore {
		base += val
	}
	return base
}

func (a *Asset) Modifer() int {
	base := 0
	switch a.Rank {
	default:
		base = -999
	case 3:
		base = -2
	case 4, 5, 6, 7:
		base = -1
	case 8, 9, 10, 11, 12, 13:
		base = 0
	case 14, 15, 16, 17:
		base = 1
	case 18:
		base = 2
	}
	for _, v := range a.BonusModifier {
		base += v
	}
	return base
}

func (a *Asset) SetScore(ns int) {
	a.Rank = ns
}
