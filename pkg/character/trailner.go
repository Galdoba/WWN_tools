package character

import (
	"errors"
	"fmt"

	"github.com/Galdoba/WWN_tools/pkg/character/asset"
	"github.com/Galdoba/WWN_tools/pkg/magic"
)

func (chr *Character) Train(skl string) {
	if skl == "" {
		return
	}
	if isSkill(skl) {
		fmt.Println("Train ", skl)
		if _, ok := chr.Skill[skl]; !ok {
			chr.Skill[skl] = asset.NewSkill(skl)
			fmt.Print(chr.Sheet())
			return
		}
		if chr.Skill[skl].Level() >= 1 {
			fmt.Println("Can't have level 2 skill for novice character")
			return
		}
		chr.Skill[skl].SetLevel(chr.Skill[skl].Level() + 1)
		fmt.Print(chr.Sheet())
		return
	}
	if isAttribute(skl) {
		chr.Attribute[skl].SetScore(chr.Attribute[skl].Score() + 1)
		fmt.Print(chr.Sheet())
		return
	}
	expanded := chr.handleSpecial(skl)
	for _, skl := range expanded {
		chr.Train(skl)
	}
	//fmt.Print(chr.Sheet())
}

func (chr *Character) handleSpecial(spec string) []string {
	automod := chr.FlagAuto
	addon := []string{}
	switch spec {
	default:
		panic("Unknown Special: " + spec)
	case "Any Combat":
		anyCombatList := []string{"Stab", "Shoot", "Punch"}
		if automod {
			addon = append(addon, chr.Dice.RollFromList(anyCombatList))
			return addon
		}
		choise := chooseOption("'Any Combat' - pick one:", anyCombatList)
		addon = append(addon, anyCombatList[choise])
		return addon
	case "Non-Combat":
		nonCombatList := []string{}
		for _, val := range allSkills() {
			switch val {
			case "Stab", "Shoot", "Punch":
				continue
			}
			nonCombatList = append(nonCombatList, val)
		}
		if automod {
			addon = append(addon, chr.Dice.RollFromList(nonCombatList))
			return addon
		}
		choise := chooseOption("'Non-Combat' - pick one:", nonCombatList)
		addon = append(addon, nonCombatList[choise])
		return addon
	case "Any Skill":
		if automod {
			addon = append(addon, chr.Dice.RollFromList(allSkills()))
			return addon
		}
		choise := chooseOption("'Any Skill' - pick one:", allSkills())
		addon = append(addon, allSkills()[choise])
		return addon
	case "+1 Any Stat":
		if automod {
			addon = append(addon, chr.Dice.RollFromList(allAttributes()))
			return addon
		}
		choise := chooseOption("'+1 Any Stat' - pick one:", allAttributes())
		addon = append(addon, allAttributes()[choise])
		return addon
	case "+2 Physical":
		physicalList := []string{STR, DEX, CON}
		if automod {
			addon = append(addon, chr.Dice.RollFromList(physicalList))
			addon = append(addon, chr.Dice.RollFromList(physicalList))
			return addon
		}
		choise := chooseOption("'+2 Physical' - pick one:", physicalList)
		addon = append(addon, physicalList[choise])
		choise = chooseOption("and another one:", physicalList)
		addon = append(addon, physicalList[choise])
		return addon
	case "+2 Mental":
		mentalList := []string{INT, WIS, CHA}
		if automod {
			addon = append(addon, chr.Dice.RollFromList(mentalList))
			addon = append(addon, chr.Dice.RollFromList(mentalList))
			return addon
		}
		choise := chooseOption("'+2 Mental' - pick one:", mentalList)
		addon = append(addon, mentalList[choise])
		choise = chooseOption("and another one:", mentalList)
		addon = append(addon, mentalList[choise])
		return addon
	case "Perform or Sneak":
		skillsOption := []string{"Perform", "Sneak"}
		switch automod {
		case true:
			addon = append(addon, chr.Dice.RollFromList(skillsOption))
		case false:
			pick := chooseOption("Choose bonus skill:", skillsOption)
			addon = append(addon, skillsOption[pick])
		}
		return addon
	case "Punch or Stab":
		skillsOption := []string{"Punch", "Stab"}
		switch automod {
		case true:
			addon = append(addon, chr.Dice.RollFromList(skillsOption))
		case false:
			pick := chooseOption("Choose bonus skill:", skillsOption)
			addon = append(addon, skillsOption[pick])
		}
	case "Specialist":
		skillsOption := specialistSkills()
		switch automod {
		case true:
			addon = append(addon, chr.Dice.RollFromList(skillsOption))
		case false:
			pick := chooseOption("Choose bonus skill:", skillsOption)
			addon = append(addon, skillsOption[pick])
		}
		return addon
	}
	return addon
}

func isSkill(skl string) bool {
	for _, val := range allSkills() {
		if skl == val {
			return true
		}
	}
	return false
}

func isAttribute(atr string) bool {
	for _, val := range allAttributes() {
		if atr == val {
			return true
		}
	}
	return false
}

func pickValid(fullList []string) []string {
	picked := []string{}
	for _, test := range fullList {
		if test != "Any Skill" {
			picked = append(picked, test)
		}
	}
	return picked
}

func (chr *Character) LearnSpell(spellName string) error {
	if chr.Tradition == nil {
		return errors.New("tradition not picked")
	}
	if chr.Class == nil {
		return errors.New("class not picked")
	}
	switch chr.Class.Name() {
	case Warrior, Expert, AdventurerEW:
		return errors.New("character unable to cast spells")
	}
	spell, err := magic.Grimoire(spellName)
	if err != nil {
		return err
	}
	chr.Tradition.AddSpell(spell.Name)
	return nil
}
