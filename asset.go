package asset

const (
	TypeAttribute  = 1
	TypeSkill      = 2
	TypeBackground = 3
	TypeClass      = 4
	TypeFoci       = 5
)

type Asset struct {
	Name          string
	Group         int
	Rank          int
	LearningCost  int
	BonusScore    map[string]int
	BonusModifier map[string]int
	Description   string
	Bonus         string
}

func New(assetType int, name string) Asset {
	a := Asset{}
	a.Name = name
	a.Group = assetType
	a.Description = callDescription(name)
	return a
}

func callDescription(name string) string {
	switch name {
	default:
		return "No description for '" + name + "'"
	case "Strength":
		return "Lifting heavy weights, breaking things, melee combat, carrying gear"
	case "Dexterity":
		return "Speed, evasion, manual dexterity, reaction time, combat initiative"
	case "Constitution":
		return "Hardiness, enduring injury, resisting poisons, going without food or rest"
	case "Intelligence":
		return "Memory, reasoning, intellectual skills, general education"
	case "Wisdom":
		return "Noticing things, making judgments, reading situations, intuition"
	case "Charisma":
		return "Force of character, charming others, attracting attention, winning loyalty"
	}
}

//SETTERS
