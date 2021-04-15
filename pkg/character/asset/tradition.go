package asset

import (
	"strings"

	"github.com/Galdoba/utils"
)

const (
	HighMage     = 1
	Elementalist = 2
	Necromancer  = 3
	Healer       = 4
	Vowed        = 5
)

type Tradition interface {
	Name() string
	AvailableSpells() []string
	AvailableArts() []string
	DualMage() bool
	SATable() string
	AddSpell(string)
	AddArt(string)
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

func (a *Asset) AvailableSpells() []string {
	return a.AssosiatedList1
}

func (a *Asset) AvailableArts() []string {
	return a.AssosiatedList2
}

func (a *Asset) AddSpell(newSpell string) {
	a.AssosiatedList1 = utils.AppendUniqueStr(a.AssosiatedList1, newSpell)
}

func (a *Asset) AddArt(newArt string) {
	a.AssosiatedList2 = append(a.AssosiatedList2, newArt)
}

func (a *Asset) DualMage() bool {
	dual := strings.Split(a.AssetName, "Partial")
	if len(dual) == 3 {
		return true
	}
	return false
}

func (a *Asset) SATable() string {
	if a.AssetName == "Full High Mage" {
		return "FHM"
	}
	if a.AssetName == "Full Elementalist" {
		return "FEM"
	}
	if a.AssetName == "Full Necromancer" {
		return "FNM"
	}
	if a.DualMage() {
		return "DMT"
	}
	if strings.Contains(a.AssetName, "High Mage") {
		return "PHM"
	}
	if strings.Contains(a.AssetName, "Elementalist") {
		return "PEM"
	}
	if strings.Contains(a.AssetName, "Necromancer") {
		return "PNM"
	}
	return "NT"
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
HMF  M 11 = Full High Mage
DPM  M 12 = Partial High Mage/Partial Elementalist
DPM  M 13 = Partial High Mage/Partial Necromancer
HMP  M 14 = Partial High Mage/Healer
HMP  M 15 = Partial High Mage/Vowed
EMF  M 22 = Full Elementalist
DPM  M 23 = Partial Elementalist/Partial Necromancer
EMP  M 24 = Partial Elementalist/Healer
EMP  M 25 = Partial Elementalist/Vowed
NMF  M 33 = Full Necromancer
NMP  M 34 = Partial Necromancer/Healer
NMP  M 35 = Partial Necromancer/Vowed
HMP  A 10 = Partial High Mage
EMP  A 20 = Partial Elementalist
NMP  A 30 = Partial Necromancer
---  A 40 = Healer
---  A 50 = Vowed

HMF  M 11 = Full High Mage
EMF  M 22 = Full Elementalist
NMF  M 33 = Full Necromancer
DPM  M 12 = Partial High Mage/Partial Elementalist
DPM  M 13 = Partial High Mage/Partial Necromancer
DPM  M 23 = Partial Elementalist/Partial Necromancer
HMP  M 14 = Partial High Mage/Healer
HMP  M 15 = Partial High Mage/Vowed
HMP  A 10 = Partial High Mage
EMP  A 20 = Partial Elementalist
EMP  M 24 = Partial Elementalist/Healer
EMP  M 25 = Partial Elementalist/Vowed
NMP  M 34 = Partial Necromancer/Healer
NMP  M 35 = Partial Necromancer/Vowed
NMP  A 30 = Partial Necromancer
---  A 40 = Healer
---  A 50 = Vowed


*/
