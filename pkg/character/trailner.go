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
	addon := []string{}
	switch spec {
	default:
		panic("Unknown Special: " + spec)
	case "Any Combat":
		anyCombatList := []string{"Stab", "Shoot", "Punch"}
		addon = append(addon, chr.ChooseOption("Choose combat skill:", anyCombatList))
	case "Non-Combat":
		nonCombatList := []string{}
		for _, val := range allSkills() {
			switch val {
			case "Stab", "Shoot", "Punch":
				continue
			}
			nonCombatList = append(nonCombatList, val)
		}
		addon = append(addon, chr.ChooseOption("Choose non-combat skill:", nonCombatList))
	case "Any Skill":
		addon = append(addon, chr.ChooseOption("Choose any skill:", allSkills()))
	case "+1 Any Stat":
		addon = append(addon, chr.ChooseOption("Choose attribute:", allAttributes()))
	case "+2 Physical":
		physicalList := []string{STR, DEX, CON}
		addon = append(addon, chr.ChooseOption("Choose physical attribute:", physicalList))
		addon = append(addon, chr.ChooseOption("...and another one:", physicalList))
	case "+2 Mental":
		mentalList := []string{INT, WIS, CHA}
		addon = append(addon, chr.ChooseOption("Choose mental attribute:", mentalList))
		addon = append(addon, chr.ChooseOption("...and another one:", mentalList))
	case "Perform or Sneak":
		skillsOption := []string{"Perform", "Sneak"}
		addon = append(addon, chr.ChooseOption("Choose bonus skill", skillsOption))
	case "Punch or Stab":
		skillsOption := []string{"Punch", "Stab"}
		addon = append(addon, chr.ChooseOption("Choose bonus skill", skillsOption))
	case "Specialist":
		skillsOption := specialistSkills()
		addon = append(addon, chr.ChooseOption("Choose bonus skill", skillsOption))
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
