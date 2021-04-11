package asset

import "strconv"

type Skill interface {
	Level() int
	LevelStr() string
	SetLevel(nl int)
}

func (a *Asset) Level() int {
	base := a.Rank
	if base > 4 {
		base = 4
	}
	if base < 0 {
		base = 0
	}
	return base
}

func (a *Asset) LevelStr() string {
	return "Level-" + strconv.Itoa(a.Level())
}

func (a *Asset) SetLevel(nl int) {
	a.Rank = nl
}

func NewSkill(name string) Skill {
	skl := Asset{}
	skl.Group = TypeSkill
	skl.Name = name
	return &skl
}
