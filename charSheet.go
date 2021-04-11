package character

import "strconv"

/*
----------------------------------------+---------------------------------------
Name: __________________________________| STR :
Player: ________________________________| DEX :
Race/Species: __________________________| CON :
Goal:         __________________________| INT :
Description:  __________________________| WIS :
________________________________________| CHA :
-----------------+









*/
func (chr *Character) Sheet() string {
	clearTerm()
	//fmt.Println("----------CLEAR SCREEN----------")
	str := ""
	if len(chr.Attribute) > 0 {
		str += "ATTRIBUTES:\n"
	}
	for _, att := range []string{STR, DEX, CON, INT, WIS, CHA} {
		if _, ok := chr.Attribute[att]; !ok {
			continue
		}
		score := chr.Attribute[att].Score()
		mod := chr.Attribute[att].Modifer()
		scoreStr := strconv.Itoa(score)
		for len(scoreStr) < 2 {
			scoreStr = " " + scoreStr
		}
		modStr := " ("
		if mod > -1 {
			modStr += " "
		}
		for len(att) < 13 {
			att += " "
		}
		modStr += strconv.Itoa(mod) + " )"
		if mod == -999 {
			modStr = ""
		}
		str += "   " + att + ": " + scoreStr + modStr + "\n"
	}
	if chr.Background != nil {
		str += "---------------------------\n"
		str += "BACKGROUND      : " + chr.Background.NameBKG() + "\n"
	}
	if len(chr.Skill) > 0 {
		str += "---------------------------\n"
		str += "SKILLS:\n"
		for _, sk := range allSkills() {
			if val, ok := chr.Skill[sk]; ok {
				skname := sk
				for len(skname) < 10 {
					skname += " "
				}
				str += "   " + skname + "	: " + val.LevelStr() + "\n"
			}
		}

	}
	if chr.Class != nil {
		str += "---------------------------\n"
		str += "CLASS           : " + chr.Class.NameClass() + "\n"
		for _, val := range chr.Class.AvailableAbbilities() {
			str += "   " + val + "\n"
		}
	}
	if chr.Foci != nil {
		str += "---------------------------\n"
		str += "FOCI:\n"
		for _, fc := range allFociList() {
			if val, ok := chr.Foci[fc]; ok {
				str += "   " + val.NameFoci() + "\n"
			}
		}
	}
	str += "---------------------------\n\n"
	return str
}
