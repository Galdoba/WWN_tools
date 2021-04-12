package asset

import "strings"

const (
	HighMage     = 1
	Elementalist = 2
	Necromancer  = 3
	Healer       = 4
	Vowed        = 5
)

type Tradition interface {
	Name() string
}

func NewTradition(names []string) Tradition {
	trd := Asset{}
	tra1 := names[0]
	tra2 := ""
	if len(names) == 2 {
		tra2 = names[1]
	}
	tradname := tra1 + "/" + tra2
	tradname = strings.TrimSuffix(tradname, "/")
	if tra1 == tra2 {
		tradname = "Full " + strings.TrimPrefix(tra1, "Partial ")
	}
	trd.AssetName = tradname

	return &trd
}

func selectTradition(mtp int) []string {
	trdName := []string{}
	return trdName
}

// func mageChoice() string {
// 	mageOptions := []string{"High Mage", "Elementalist", "Necromancer", "No, I want to merge traditions!"}
// 	chosen := chooseOption("As a Full Mage your character can master one tradition or merge two of them to be more unique:", mageOptions)
// 	if chosen == 3 {
// 		return mergeTraditions()
// 	}
// 	return mageOptions[chosen]
// }

/*
M 11 = Full High Mage
M 12 = Partial High Mage/Partial Elementalist
M 13 = Partial High Mage/Partial Necromancer
M 14 = Partial High Mage/Healer
M 15 = Partial High Mage/Vowed
M 22 = Full Elementalist
M 23 = Partial Elementalist/Partial Necromancer
M 24 = Partial Elementalist/Healer
M 25 = Partial Elementalist/Vowed
M 33 = Full Necromancer
M 34 = Partial Necromancer/Healer
M 35 = Partial Necromancer/Vowed
A 10 = Partial High Mage
A 20 = Partial Elementalist
A 30 = Partial Necromancer
A 40 = Healer
A 50 = Vowed


*/
