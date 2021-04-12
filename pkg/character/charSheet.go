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
		str += "BACKGROUND      : " + chr.Background.Name() + "\n"
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
		str += "CLASS           : " + chr.Class.Name() + "\n"
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
	if chr.Tradition != nil {
		str += "---------------------------\n"
		str += "MAGIC TRADITION:\n"
		str += "   " + chr.Tradition.Name() + "\n"
	}
	str += "---------------------------\n\n"
	return str
}

func (chr *Character) DrawAttributeBlock() string {
	/*
			01234567890123456789012345678901234567890123456789012345678901234567890123456789 -- 79
			+-----------------------------------------------------------------------------+
			| NAME:                                 | LEVEL:                              |
			+---ATTRIBUTES----------------------------------------------------------------+
			| STR:  3 | DEX: 14 | CON:  9 | INT:  6 | WIS: 11 | CHA:  8 |   W O R L D S   |
			|   - 2   |   + 1   |    0    |   - 1   |   + 1   |    0    | WITHOUT  NUMBER |
		    +---INFO--------------------------SAVES----------BONUS------------------------+
			| RACE      : Human           | Physical: XX |  Base Attack : XX |            |
			| CLASS     : Adventurer (EM) | Evasion : XX | Melee Attack : XX | Effort  XX |
			| BACKGROUND: Performer       |  Mental : XX | Ranged Attack: XX | HP     XXX |
			|                             |  Luck   : XX |  Initiative  : XX |            |
			+---SKILLS---------------------------------------FOCI-------------------------+
			| Administer 0 | Lead       0 | Sail       0 | Specialist (Administer)  Lvl 1 |
			| Connect    0 | Magic      0 | Shoot      0 | Specialist (Administer)  Lvl 1 |
			| Convince   0 | Notice     0 | Sneak      0 | Specialist (Administer)  Lvl 1 |
			| Craft      0 | Perform    0 | Stab       0 | Specialist (Administer)  Lvl 1 |
			| Exert      0 | Pray       0 | Survive    0 | Specialist (Administer)  Lvl 1 |
			| Heal       0 | Punch      0 | Trade      0 | Specialist (Administer)  Lvl 1 |
			| Know       0 | Ride       0 | Work       0 | Specialist (Administer)  Lvl 1 |
			+---MAGIC---------------------------------------------------------------------+
			|   TRADITION: Partial Vothite Thought Noble/Partial Sarulite Blood Priest    |
			| ARTS (unlimited 0-22)                 | SPELLS (unlimited)                  |
			| ARTS (unlimited 0-22)                 | SPELLS (unlimited)                  |
			| ARTS (unlimited 0-22)                 | SPELLS (unlimited)                  |
			+-----------------------------------------------------------------------------+




	*/
	return ""
}
