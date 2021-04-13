package character

import (
	"sort"
	"strconv"
)

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
	// spells := magic.Spells("Any", 2)
	// for i := range spells {
	// 	fmt.Println(spells[i].Tradition, spells[i].Level, spells[i].Name)
	// }
	// panic(5)
	clearTerm()
	return chr.StatBlock()

}

func (chr *Character) StatBlock() string {
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
	lSep := "| "
	rSep := " |\n"
	mSep := " | "
	name := trimToLen(chr.Name, 31)
	level := trimToLen(strconv.Itoa(chr.Level), 28)
	strAtr := chr.onSheetAtr(STR)
	dexAtr := chr.onSheetAtr(DEX)
	conAtr := chr.onSheetAtr(CON)
	intAtr := chr.onSheetAtr(INT)
	wisAtr := chr.onSheetAtr(WIS)
	chaAtr := chr.onSheetAtr(CHA)
	strMod := chr.onSheetMod(STR)
	dexMod := chr.onSheetMod(DEX)
	conMod := chr.onSheetMod(CON)
	intMod := chr.onSheetMod(INT)
	wisMod := chr.onSheetMod(WIS)
	chaMod := chr.onSheetMod(CHA)
	/////////////////////////DRAW
	str := ""
	str += "+-----------------------------------------------------------------------------+\n"
	str += lSep + "NAME: " + name + mSep + "LEVEL: " + level + rSep
	str += "+---ATTRIBUTES----------------------------------------------------------------+\n"
	str += lSep + "STR: " + strAtr + mSep + "DEX: " + dexAtr + mSep + "CON: " + conAtr + mSep + "INT: " + intAtr + mSep + "WIS: " + wisAtr + mSep + "CHA: " + chaAtr + rSep
	str += lSep + strMod + mSep + dexMod + mSep + conMod + mSep + intMod + mSep + wisMod + mSep + chaMod + rSep
	str += "+---INFO--------------------------SAVES----------BONUS------------------------+\n"
	str += lSep + "RACE      : " + trimToLen(chr.Race, 15) + mSep + "Physical: " + "XX" + mSep + " Base Attack : " + "XX" + mSep + "          " + rSep
	str += lSep + "BACKGROUND: " + trimToLen(chr.onSheetBackground(), 15) + mSep + "Evasion : " + "XX" + mSep + "Melee Attack : " + "XX" + mSep + " Effort XX" + rSep
	str += lSep + "CLASS     : " + trimToLen(chr.onSheetClass(), 15) + mSep + " Mental : " + "XX" + mSep + "Ranged Attack: " + "XX" + mSep + " HP     XX" + rSep
	str += lSep + "                           " + mSep + " Luck   : " + "XX" + mSep + " Initiative  : " + "XX" + mSep + "          " + rSep
	str += "+---SKILLS--------------------+--------------+---FOCI-------------------------+\n"
	str += lSep + chr.onSheetSkill(Administer) + mSep + chr.onSheetSkill(Lead) + mSep + chr.onSheetSkill(Sail) + mSep + chr.onSheetFoci(0) + rSep
	str += lSep + chr.onSheetSkill(Connect) + mSep + chr.onSheetSkill(Magic) + mSep + chr.onSheetSkill(Shoot) + mSep + chr.onSheetFoci(1) + rSep
	str += lSep + chr.onSheetSkill(Convince) + mSep + chr.onSheetSkill(Notice) + mSep + chr.onSheetSkill(Sneak) + mSep + chr.onSheetFoci(2) + rSep
	str += lSep + chr.onSheetSkill(Craft) + mSep + chr.onSheetSkill(Perform) + mSep + chr.onSheetSkill(Stab) + mSep + chr.onSheetFoci(3) + rSep
	str += lSep + chr.onSheetSkill(Exert) + mSep + chr.onSheetSkill(Pray) + mSep + chr.onSheetSkill(Survive) + mSep + chr.onSheetFoci(4) + rSep
	str += lSep + chr.onSheetSkill(Heal) + mSep + chr.onSheetSkill(Punch) + mSep + chr.onSheetSkill(Trade) + mSep + chr.onSheetFoci(5) + rSep
	str += lSep + chr.onSheetSkill(Know) + mSep + chr.onSheetSkill(Ride) + mSep + chr.onSheetSkill(Work) + mSep + chr.onSheetFoci(6) + rSep
	str += "+---MAGIC---------------------------------------------------------------------+\n"
	str += lSep + chr.onSheetMagic() + rSep
	table := chr.artSpellTable()
	for i := 0; i < len(table); i++ {
		str += lSep + table[i] + rSep
	}
	str += "+-----------------------------------------------------------------------------+\n"
	//| Administer 0 | Lead       0 | Sail       0 | Specialist (Administer)  Lvl 1 |
	//| Connect    0 | Magic      0 | Shoot      0 | Specialist (Administer)  Lvl 1 |
	//| Convince   0 | Notice     0 | Sneak      0 | Specialist (Administer)  Lvl 1 |
	//| Craft      0 | Perform    0 | Stab       0 | Specialist (Administer)  Lvl 1 |
	//| Exert      0 | Pray       0 | Survive    0 | Specialist (Administer)  Lvl 1 |
	//| Heal       0 | Punch      0 | Trade      0 | Specialist (Administer)  Lvl 1 |
	//| Know       0 | Ride       0 | Work       0 | Specialist (Administer)  Lvl 1 |
	return str
}

func trimToLen(str string, ln int, right ...bool) string {
	pushRight := false
	if len(right) > 0 {
		pushRight = right[0]
	}
	for len(str) < ln {
		switch pushRight {
		case false:
			str += " "
		case true:
			str = " " + str
		}
	}
	for len(str) > ln+1 {
		str = string([]byte(str)[0 : len(str)-1])
	}
	return str
}

func (chr *Character) onSheetAtr(data string) string {
	atr := ""
	switch data {
	case STR, DEX, CON, INT, WIS, CHA:
		if val, ok := chr.Attribute[data]; ok {
			atr = strconv.Itoa(val.Score())
			if atr == "0" {
				atr = ""
			}
		}
	}
	atr = trimToLen(atr, 5, true)
	return atr
}

func (chr *Character) onSheetMod(data string) string {
	atr := " "
	switch data {
	case STR, DEX, CON, INT, WIS, CHA:
		if val, ok := chr.Attribute[data]; ok {
			mod := val.Modifer()
			switch mod {
			case -999:
				atr += "         "
			case 0:
				atr += "    0    "
			default:
				if mod < 0 {
					atr += "   - " + strconv.Itoa(mod*-1) + "   "
				}
				if mod > 0 {
					atr += "   + " + strconv.Itoa(mod) + "   "
				}
			}
		}
	}
	//atr = trimToLen(atr, 2, true)
	return atr
}

func (chr *Character) onSheetClass() string {
	cls := ""
	if chr.Class != nil {
		cls = chr.Class.Name()
	}
	return cls
}

func (chr *Character) onSheetBackground() string {
	bgr := ""
	if chr.Background != nil {
		bgr = chr.Background.Name()
	}
	return bgr
}

func (chr *Character) onSheetSkill(skl string) string {
	//skl = trimToLen(skl, 10)
	for k, val := range chr.Skill {
		if k == skl {
			skl = trimToLen(skl, 10)
			skl = skl + " " + strconv.Itoa(val.Level()) // + val.LevelStr()
			skl = trimToLen(skl, 12)
			return skl
		}
	}
	skl = trimToLen(skl, 12)
	return skl
}

func (chr *Character) onSheetFoci(n int) string {
	fociList := []string{}
	for k, v := range chr.Foci {
		fociList = append(fociList, trimToLen(k, 25)+"Lvl "+strconv.Itoa(v.Level()))
	}
	sort.Sort(sort.StringSlice(fociList))
	for len(fociList) < 7 {
		fociList = append(fociList, trimToLen("", 30))
	}
	switch n {
	case 0, 1, 2, 3, 4, 5, 6:
		return fociList[n]
	}
	return "fociList"
}

func (chr *Character) onSheetMagic() string {
	if chr.Tradition != nil {
		return "TRADITION: " + trimToLen(chr.Tradition.Name(), 64)
	}
	//artList := []string{}
	return " Character can not cast spells                                             "
}

func (chr *Character) artSpellTable() []string {
	if chr.Tradition == nil {
		return []string{}
	}
	artList := []string{}
	artList = append(artList, "Art 1")
	artList = append(artList, "Art 2")
	artList = append(artList, "Art 3")
	artList = append(artList, "Art 4")
	artList = append(artList, "Art 5")
	artList = append(artList, "Art 6")
	spellList := []string{}
	spellList = append(spellList, "Spell 1")
	spellList = append(spellList, "Spell 2")
	spellList = append(spellList, "Spell 3")
	max := len(artList)
	if len(spellList) > max {
		max = len(spellList)
	}
	for len(artList) < max {
		artList = append(artList, " ")
	}
	for len(spellList) < max {
		spellList = append(spellList, " ")
	}
	table := []string{"---ARTS------------------------------+---SPELLS----------------------------"}
	for i := 0; i < max; i++ {
		table = append(table, trimToLen(artList[i], 36)+" | "+trimToLen(spellList[i], 36))
	}
	return table
}
