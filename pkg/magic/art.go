package magic

type Art interface {
	Name() string
	Tradition() string
	Description() string
}

func ArtListByTradition(tradition string) []string {
	suggest := []string{}
	allArt := allArts()
	l := len(allArt)
	for i := 0; i <= l; i++ {
		if v, ok := allArt[i]; ok {
			if v.Tradition() == tradition {
				suggest = append(suggest, v.Name())
			}
		}

	}
	return suggest
}

func allArts() map[int]Art {
	artList := make(map[int]Art)
	artList[1] = &construct{1, TraditionHighMagic, 0, "Arcane Lexicon", ""}
	artList[2] = &construct{1, TraditionHighMagic, 0, "Counter Magic", ""}
	artList[3] = &construct{1, TraditionHighMagic, 0, "Empowered Sorcery", ""}
	artList[4] = &construct{1, TraditionHighMagic, 0, "Hang Sorcery", ""}
	artList[5] = &construct{1, TraditionHighMagic, 0, "Inexorable Effect", ""}
	artList[6] = &construct{1, TraditionHighMagic, 0, "Iron Resolution", ""}
	artList[7] = &construct{1, TraditionHighMagic, 0, "Preparatory Countermagic", ""}
	artList[8] = &construct{1, TraditionHighMagic, 0, "Psychic Conversion", ""}
	artList[9] = &construct{1, TraditionHighMagic, 0, "Restrained Casting", ""}
	artList[10] = &construct{1, TraditionHighMagic, 0, "Retain Sorcery", ""}
	artList[11] = &construct{1, TraditionHighMagic, 0, "Sense Magic", ""}
	artList[12] = &construct{1, TraditionHighMagic, 0, "Suppress Magic", ""}
	artList[13] = &construct{1, TraditionHighMagic, 0, "Swift Casting", ""}
	artList[14] = &construct{1, TraditionHighMagic, 0, "Ward Allies", ""}
	artList[15] = &construct{1, TraditionHighMagic, 0, "Wizard’s Grandeur", ""}

	artList[16] = &construct{1, TraditionElementalist, 0, "Beckoned Deluge", ""}
	artList[17] = &construct{1, TraditionElementalist, 0, "Earthsight", ""}
	artList[18] = &construct{1, TraditionElementalist, 0, "Elemental Blast", ""}
	artList[19] = &construct{1, TraditionElementalist, 0, "Flamesight", ""}
	artList[20] = &construct{1, TraditionElementalist, 0, "Pavis of Elements", ""}
	artList[21] = &construct{1, TraditionElementalist, 0, "Petrifying Stare", ""}
	artList[22] = &construct{1, TraditionElementalist, 0, "Rune of Destruction", ""}
	artList[23] = &construct{1, TraditionElementalist, 0, "Steps of Air", ""}
	artList[24] = &construct{1, TraditionElementalist, 0, "Stunning Shock", ""}
	artList[25] = &construct{1, TraditionElementalist, 0, "Thermal Shield", ""}

	artList[26] = &construct{1, TraditionHealer, 0, "Empowered Healer", ""}
	artList[27] = &construct{1, TraditionHealer, 0, "Facile Healer", ""}
	artList[28] = &construct{1, TraditionHealer, 0, "Far Healer", ""}
	artList[29] = &construct{1, TraditionHealer, 0, "Final Repose", ""}
	artList[30] = &construct{1, TraditionHealer, 0, "Healer’s Eye", ""}
	artList[31] = &construct{1, TraditionHealer, 0, "Limb Restoration", ""}
	artList[32] = &construct{1, TraditionHealer, 0, "Purge Ailment", ""}
	artList[33] = &construct{1, TraditionHealer, 0, "Refined Restoration", ""}
	artList[34] = &construct{1, TraditionHealer, 0, "Revive the Fallen", ""}
	artList[35] = &construct{1, TraditionHealer, 0, "Swift Healer", ""}
	artList[36] = &construct{1, TraditionHealer, 0, "The Healer’s Knife", ""}
	artList[37] = &construct{1, TraditionHealer, 0, "Tireless Vigor", ""}
	artList[38] = &construct{1, TraditionHealer, 0, "Vital Furnace", ""}

	artList[39] = &construct{1, TraditionNecromancer, 0, "Bonetalker", ""}
	artList[40] = &construct{1, TraditionNecromancer, 0, "Cold Flesh", ""}
	artList[41] = &construct{1, TraditionNecromancer, 0, "Consume Life Energy", ""}
	artList[42] = &construct{1, TraditionNecromancer, 0, "False Death", ""}
	artList[43] = &construct{1, TraditionNecromancer, 0, "Gravesight", ""}
	artList[44] = &construct{1, TraditionNecromancer, 0, "Keeper of the Gate", ""}
	artList[45] = &construct{1, TraditionNecromancer, 0, "Life Bridge", ""}
	artList[46] = &construct{1, TraditionNecromancer, 0, "Master of Bones", ""}
	artList[47] = &construct{1, TraditionNecromancer, 0, "Red Harvest", ""}
	artList[48] = &construct{1, TraditionNecromancer, 0, "Unaging", ""}
	artList[49] = &construct{1, TraditionNecromancer, 0, "Uncanny Ichor", ""}
	artList[50] = &construct{1, TraditionNecromancer, 0, "Unliving Persistence", ""}

	artList[51] = &construct{1, TraditionVowed, 0, "Brutal Counter", ""}
	artList[52] = &construct{1, TraditionVowed, 0, "Faultless Awareness", ""}
	artList[53] = &construct{1, TraditionVowed, 0, "Hurling Throw", ""}
	artList[54] = &construct{1, TraditionVowed, 0, "The Inward Eye", ""}
	artList[55] = &construct{1, TraditionVowed, 0, "Leap of the Heavens", ""}
	artList[56] = &construct{1, TraditionVowed, 0, "Master’s Vigor", ""}
	artList[57] = &construct{1, TraditionVowed, 0, "Mob Justice", ""}
	artList[58] = &construct{1, TraditionVowed, 0, "Nimble Ascent", ""}
	artList[59] = &construct{1, TraditionVowed, 0, "Purified Body", ""}
	artList[60] = &construct{1, TraditionVowed, 0, "Revivifying Breath", ""}
	artList[61] = &construct{1, TraditionVowed, 0, "Shattering Strike", ""}
	artList[62] = &construct{1, TraditionVowed, 0, "Style Weaponry", ""}
	artList[63] = &construct{1, TraditionVowed, 0, "Unobtrusive Step", ""}

	return artList
}
