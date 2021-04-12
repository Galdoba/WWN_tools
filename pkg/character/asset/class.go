package asset

type Class interface {
	AvailableAbbilities() []string
	Name() string
}

func NewClass(className string) Class {
	cls := Asset{}
	cls.Group = TypeClass
	cls.AssetName = className
	return &cls
}

func (a *Asset) AvailableAbbilities() []string {
	abbi := []string{}
	switch a.AssetName {
	default:
		panic("unknown class detected '" + a.AssetName + "'")
	case "Warrior":
		abbi = append(abbi, "Killing Blow")
		abbi = append(abbi, "Veteran’s Luck")
	case "Expert":
		abbi = append(abbi, "Masterful Expertise")
		abbi = append(abbi, "Quick Learner")
	case "Mage":
		abbi = append(abbi, "Arcane Tradition")
	case "Adventurer (ew)":
		abbi = append(abbi, "Quick Learner")
	case "Adventurer (em)":
		abbi = append(abbi, "Quick Learner")
		abbi = append(abbi, "Arcane Tradition")
	case "Adventurer (mw)":
		abbi = append(abbi, "Arcane Tradition")

	}
	return abbi
}

func (a *Asset) HitDice(lvl int) string {
	switch lvl {
	case 1:
		return "1d6"
	}
	return "xd6"
}
