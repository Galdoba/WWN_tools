package magic

const (
	TraditionHighMagic    = "High Magic"
	TraditionElementalist = "Elementalist"
	TraditionNecromancer  = "Necromancer"
)

type Spell struct {
	Tradition   string
	Level       int
	Name        string
	Description string
}

func allSpells() map[int]Spell {
	spMap := make(map[int]Spell)
	spMap[0] = Spell{TraditionHighMagic, 1, "Apprehending the Arcane Form", ""}
	spMap[1] = Spell{TraditionHighMagic, 1, "Cognitive Supersession of the Inferior Orders", ""}
	spMap[2] = Spell{TraditionHighMagic, 1, "The Coruscating Coffin", ""}
	spMap[3] = Spell{TraditionHighMagic, 1, "Damnation of the Sense", ""}
	spMap[4] = Spell{TraditionHighMagic, 1, "Decree of Ligneous Dissolution", ""}
	spMap[5] = Spell{TraditionHighMagic, 1, "The Excellent Transpicuous Transformation", ""}
	spMap[6] = Spell{TraditionHighMagic, 1, "Imperceptible Cerebral Divulgence", ""}
	spMap[7] = Spell{TraditionHighMagic, 1, "Ineluctable Shackles of Volition", ""}
	spMap[8] = Spell{TraditionHighMagic, 1, "The Long Amber Moment", ""}
	spMap[9] = Spell{TraditionHighMagic, 1, "Phantasmal Mimesis", ""}
	spMap[10] = Spell{TraditionHighMagic, 1, "Velocitous Imbuement", ""}
	spMap[11] = Spell{TraditionHighMagic, 1, "Wardpact Invocation", ""}
	spMap[12] = Spell{TraditionHighMagic, 1, "The Wind of the Final Repose", ""}
	spMap[13] = Spell{TraditionHighMagic, 2, "Calculation of the Evoked Servitor", ""}
	spMap[14] = Spell{TraditionHighMagic, 2, "Casting Forth the Inner Eye", ""}
	spMap[15] = Spell{TraditionHighMagic, 2, "Conjunction of the Inexorable Step", ""}
	spMap[16] = Spell{TraditionHighMagic, 2, "Decree of Lithic Dissolution", ""}
	spMap[17] = Spell{TraditionHighMagic, 2, "Extirpate Arcana", ""}
	spMap[18] = Spell{TraditionHighMagic, 2, "The Inexorable Imputation", ""}
	spMap[19] = Spell{TraditionHighMagic, 2, "Jade Palanquin of the Faceless God", ""}
	spMap[20] = Spell{TraditionHighMagic, 2, "Mantle of Disjecting Dissection", ""}
	spMap[21] = Spell{TraditionHighMagic, 2, "Prudentially Transient Abnegation of Life", ""}
	spMap[22] = Spell{TraditionHighMagic, 2, "Resounding Temporal Echo", ""}
	spMap[23] = Spell{TraditionHighMagic, 2, "The Verdant Vallation", ""}
	spMap[24] = Spell{TraditionHighMagic, 2, "Visitation of the Clement Clime", ""}
	spMap[25] = Spell{TraditionHighMagic, 3, "Adopt the Simulacular Visage", ""}
	spMap[26] = Spell{TraditionHighMagic, 3, "Conjunct the Vital Viscera", ""}
	spMap[27] = Spell{TraditionHighMagic, 3, "Exhalation of Congelating Cold", ""}
	spMap[28] = Spell{TraditionHighMagic, 3, "Foresightful Apprehension", ""}
	spMap[29] = Spell{TraditionHighMagic, 3, "Glass Chimes of the Bamboo Terrace", ""}
	spMap[30] = Spell{TraditionHighMagic, 3, "The Howl of Light", ""}
	spMap[31] = Spell{TraditionHighMagic, 3, "Phobic Storm", ""}
	spMap[32] = Spell{TraditionHighMagic, 3, "Scorn the Fetters of Earth", ""}
	spMap[33] = Spell{TraditionHighMagic, 3, "The Torment of Tumefaction", ""}
	spMap[34] = Spell{TraditionHighMagic, 3, "Touch of Elucidating Intangibility", ""}
	spMap[35] = Spell{TraditionHighMagic, 3, "Vallation of Specified Exclusion", ""}
	spMap[36] = Spell{TraditionHighMagic, 4, "Calculation of the Phantasmal Eidolon", ""}
	spMap[37] = Spell{TraditionHighMagic, 4, "Contingent Excision of Arcana", ""}
	spMap[38] = Spell{TraditionHighMagic, 4, "Disjunctive Temporal Reversion", ""}
	spMap[39] = Spell{TraditionHighMagic, 4, "Evert the Inwardness", ""}
	spMap[40] = Spell{TraditionHighMagic, 4, "The Grinding Geas", ""}
	spMap[41] = Spell{TraditionHighMagic, 4, "Obnubilation of the Will", ""}
	spMap[42] = Spell{TraditionHighMagic, 4, "Ochre Sigil of Juxtaposition", ""}
	spMap[43] = Spell{TraditionHighMagic, 4, "Pierce the Pallid Gate", ""}
	spMap[44] = Spell{TraditionHighMagic, 4, "Sigil of Aeolian Auctoritas", ""}
	spMap[45] = Spell{TraditionHighMagic, 5, "Abdication of Temporal Presence", ""}
	spMap[46] = Spell{TraditionHighMagic, 5, "Banishment to the Black Glass Labyrinth", ""}
	spMap[47] = Spell{TraditionHighMagic, 5, "The Dazzling Prismatic Hemicycle", ""}
	spMap[48] = Spell{TraditionHighMagic, 5, "Deluge of Hell", ""}
	spMap[49] = Spell{TraditionHighMagic, 5, "The Earth as Clay", ""}
	spMap[50] = Spell{TraditionHighMagic, 5, "Invocation of the Invincible Citadel", ""}
	spMap[51] = Spell{TraditionHighMagic, 5, "Open the High Road", ""}
	spMap[52] = Spell{TraditionElementalist, 1, "Aqueous Harmony Level", ""}
	spMap[53] = Spell{TraditionElementalist, 1, "Flame Scrying Level", ""}
	spMap[54] = Spell{TraditionElementalist, 1, "Elemental Favor Level", ""}
	spMap[55] = Spell{TraditionElementalist, 1, "Elemental Spy Level", ""}
	spMap[56] = Spell{TraditionElementalist, 2, "Boreal Wings Level", ""}
	spMap[57] = Spell{TraditionElementalist, 2, "The Burrower Below Level", ""}
	spMap[58] = Spell{TraditionElementalist, 2, "Flame Without End Level", ""}
	spMap[59] = Spell{TraditionElementalist, 2, "Pact of Stone and Sea Level", ""}
	spMap[60] = Spell{TraditionElementalist, 3, "Elemental Vallation Level", ""}
	spMap[61] = Spell{TraditionElementalist, 3, "Like the Stones Level", ""}
	spMap[62] = Spell{TraditionElementalist, 3, "Wind Walking Level", ""}
	spMap[63] = Spell{TraditionElementalist, 4, "Calcifying Scourge Level", ""}
	spMap[64] = Spell{TraditionElementalist, 4, "Elemental Guardian Level", ""}
	spMap[65] = Spell{TraditionElementalist, 5, "Fury of the Elements Level", ""}
	spMap[66] = Spell{TraditionElementalist, 5, "Tremors of the Depths Level", ""}
	spMap[67] = Spell{TraditionNecromancer, 1, "Command the Dead", ""}
	spMap[68] = Spell{TraditionNecromancer, 1, "Query the Skull", ""}
	spMap[69] = Spell{TraditionNecromancer, 1, "Smite the Dead", ""}
	spMap[70] = Spell{TraditionNecromancer, 1, "Terrible Liveliness", ""}
	spMap[71] = Spell{TraditionNecromancer, 2, "Augment Mortal Vitality", ""}
	spMap[72] = Spell{TraditionNecromancer, 2, "Enfeebling Wave", ""}
	spMap[73] = Spell{TraditionNecromancer, 2, "Final Death", ""}
	spMap[74] = Spell{TraditionNecromancer, 2, "Raise Corpse", ""}
	spMap[75] = Spell{TraditionNecromancer, 3, "Compel Flesh", ""}
	spMap[76] = Spell{TraditionNecromancer, 3, "Festering Curse", ""}
	spMap[77] = Spell{TraditionNecromancer, 3, "Forgetting the Grave", ""}
	spMap[78] = Spell{TraditionNecromancer, 3, "Merge Souls", ""}
	spMap[79] = Spell{TraditionNecromancer, 4, "Boneshaper", ""}
	spMap[80] = Spell{TraditionNecromancer, 4, "Raise Grave Knight", ""}
	spMap[81] = Spell{TraditionNecromancer, 5, "Call of the Tomb", ""}
	spMap[82] = Spell{TraditionNecromancer, 5, "Everlasting", ""}
	return spMap
}

func FilterSpellsByTradition(tradition string, maxLevel int) []string {
	spMap := allSpells()

	filteredMap := make(map[int]Spell)
	for k, v := range spMap {
		if tradition != "Any" && tradition != v.Tradition {
			continue
		}
		if v.Level <= maxLevel {
			filteredMap[k] = v
		}
	}
	spells := []string{}
	for i := 0; i < 100; i++ {
		if val, ok := filteredMap[i]; ok {
			spells = append(spells, val.Name)
		}
	}
	return spells
}
