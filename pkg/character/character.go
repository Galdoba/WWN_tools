package character

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Galdoba/WWN_tools/pkg/character/asset"
	"github.com/Galdoba/WWN_tools/pkg/dice"
	"github.com/Galdoba/WWN_tools/pkg/magic"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/devtools/cli/user"
)

const (
	//ATTRIBUTES
	STR = "Strength"
	DEX = "Dexterity"
	CON = "Constitution"
	INT = "Intelligence"
	WIS = "Wisdom"
	CHA = "Charisma"
	//SKILLS
	Administer = "Administer"
	Connect    = "Connect"
	Convince   = "Convince"
	Craft      = "Craft"
	Exert      = "Exert"
	Heal       = "Heal"
	Know       = "Know"
	Lead       = "Lead"
	Magic      = "Magic"
	Notice     = "Notice"
	Perform    = "Perform"
	Pray       = "Pray"
	Punch      = "Punch"
	Ride       = "Ride"
	Sail       = "Sail"
	Shoot      = "Shoot"
	Sneak      = "Sneak"
	Stab       = "Stab"
	Survive    = "Survive"
	Trade      = "Trade"
	Work       = "Work"
	//special
	AnyCombat = "Any Combat"
	//BACKGROUNDS
	Artisan   = "Artisan"
	Barbarian = "Barbarian"
	Carter    = "Carter"
	Courtesan = "Courtesan"
	Criminal  = "Criminal"
	Hunter    = "Hunter"
	Laborer   = "Laborer"
	Merchant  = "Merchant"
	Noble     = "Noble"
	Nomad     = "Nomad"
	Peasant   = "Peasant"
	Performer = "Performer"
	Physician = "Physician"
	Priest    = "Priest"
	Sailor    = "Sailor"
	Scholar   = "Scholar"
	Slave     = "Slave"
	Soldier   = "Soldier"
	Thug      = "Thug"
	Wanderer  = "Wanderer"
	//CLASSES
	Warrior      = "Warrior"
	Expert       = "Expert"
	Mage         = "Mage"
	AdventurerEW = "Adventurer (ew)"
	AdventurerEM = "Adventurer (em)"
	AdventurerMW = "Adventurer (mw)"
	//FOCI
	//TRADITIONS
	HighMage     = "Partial High Mage"
	Elementalist = "Partial Elementalist"
	Necromancer  = "Partial Necromancer"
	Healer       = "Healer"
	Vowed        = "Vowed"
)

type Character struct {
	Name       string
	Race       string
	FlagAuto   bool
	Dice       *dice.Dicepool
	Level      int
	Attribute  map[string]asset.Attribute
	Background asset.Background
	Skill      map[string]asset.Skill
	Class      asset.Class
	Foci       map[string]asset.Foci
	Tradition  asset.Tradition
	//Background
}

func New(auto bool, seed ...string) Character {
	chr := Character{}
	chr.FlagAuto = auto
	chr.Dice = dice.New()
	if len(seed) > 0 {
		chr.Dice.SetSeed(seed[0])
		chr.Name = seed[0]
	}
	if chr.Name == "" {
		chr.Name = "Unknown Hero"
	}
	chr.Race = "Human"
	chr.Level = 1
	return chr
}

func (chr *Character) SetAttributes() {
	chr.Attribute = make(map[string]asset.Attribute)
	atrArrayStr := []string{STR, DEX, CON, INT, WIS, CHA}
	for _, val := range atrArrayStr {
		chr.Attribute[val] = asset.NewAttribute(val)
	}
	if chr.FlagAuto {
		to14 := []string{}
		for k := range chr.Attribute {
			chr.Attribute[k].SetScore(dice.Roll3D())
			if chr.Attribute[k].Score() < 14 {
				to14 = append(to14, k)
			}
			fmt.Println(chr.Sheet())
		}
		r := chr.Dice.RollFromList(to14)
		chr.Attribute[r].SetScore(14)
		fmt.Println(chr.Sheet())
		return
	}
	//atrArray := []int{14, 12, 11, 10, 9, 7}
	method := chooseOption("Attributes reflect the basic potential of your hero. You can:", []string{
		"Roll 3d6 six times and assign them in order",
		"Use an array of 14, 12, 11, 10, 9, 7 assigned as you wish",
	})
	switch method {
	case 0:
		for k := range chr.Attribute {
			chr.Attribute[k].SetScore(dice.Roll3D())
			fmt.Println(chr.Sheet())
		}
		raise := append(atrArrayStr, "Naaah... All is fine")
		r := chooseOption("You can set one of the attribute to 14...", raise)
		if r < len(atrArrayStr) {
			chr.Attribute[atrArrayStr[r]].SetScore(14)
		}
		fmt.Println(chr.Sheet())
	case 1:
		atrArray := []int{14, 12, 11, 10, 9, 7}
		for _, suggest := range atrArray {
			validOptions := []string{}
			for _, val := range atrArrayStr {
				if chr.Attribute[val].Score() == 0 {
					validOptions = append(validOptions, val)
				}
			}
			fmt.Println(chr.Sheet())
			chosen := chooseOption("Choose Attribute to set "+strconv.Itoa(suggest), validOptions)
			chr.Attribute[validOptions[chosen]].SetScore(suggest)
		}
		fmt.Println(chr.Sheet())

	}
}

////////////////////////Backgrounds
func (chr *Character) SetBackground() {
	if chr.FlagAuto {
		bkg := chr.Dice.RollFromList(allBackgrounds())
		chr.Background = asset.NewBackground(bkg)
		return
	}
	validOptions := allBackgrounds()
	validOptions = append(validOptions, "Naaah... just roll it")
	chsen := chooseOption("Select background for a character:\n*check description and bonuses in CRB p.12-17", validOptions)
	bkg := ""
	if chsen >= len(allBackgrounds()) {
		bkg = chr.Dice.RollFromList(allBackgrounds())
	} else {
		bkg = allBackgrounds()[chsen]
	}
	chr.Background = asset.NewBackground(bkg)
}

////////////////////////SKILLS
func (chr *Character) SetSkills() {
	chr.Skill = make(map[string]asset.Skill)
	freeSkill := chr.Background.FreeSkill()
	method := 0
	chr.Train(freeSkill)
	if chr.FlagAuto {
		method = chr.Dice.RollNext("1d3").Sum()
	}
	if method == 0 {
		method = chooseOption("Pick one of the three options below", []string{"Gain the background’s listed quick skills", "Pick two skills from the background’s Learning table", "Roll three times, splitting the rolls as you wish between the Growth and Learning tables for your background"}) + 1
	}
	switch method {
	case 1:
		quickSkills := chr.Background.QuickSkills()
		for _, skl := range quickSkills {
			chr.Train(skl)
		}
		return
	case 2:
		learn := chr.Background.Learning()
		learn = pickValid(learn)
		chosen := []string{}
		switch chr.FlagAuto {
		case true:
			chosen = append(chosen, chr.Dice.RollFromList(learn))
			chosen = append(chosen, chr.Dice.RollFromList(learn))
		case false:
			for len(chosen) < 2 {
				c := chooseOption(strconv.Itoa(len(chosen)+1)+"/2: "+"Pick one skill from options below ", learn)
				chosen = append(chosen, learn[c])
			}
		}
		for _, skl := range chosen {
			chr.Train(skl)
		}
		return
	case 3:
		chosen := []string{}
		switch chr.FlagAuto {
		case true:
			for len(chosen) < 3 {
				switch chr.Dice.RollNext("1d2").Sum() {
				case 1:
					chosen = append(chosen, chr.Dice.RollFromList(chr.Background.Growth()))
				case 2:
					chosen = append(chosen, chr.Dice.RollFromList(chr.Background.Learning()))
				}
			}
			for _, skl := range chosen {
				chr.Train(skl)
			}
		case false:
			for i := 1; i < 4; i++ {
				fmt.Println("Roll", i, "of 3:")
				options := []string{"Growth", "Learning"}
				c := chooseOption("Pick table you with to roll:"+chr.Background.Tables(), options)
				//chosen = append(chosen, options[c])
				switch options[c] {
				case "Growth":
					chr.Train(chr.Dice.RollFromList(chr.Background.Growth()))
				case "Learning":
					chr.Train(chr.Dice.RollFromList(chr.Background.Learning()))
				}
				//fmt.Println(chr.Sheet())
			}
		}
	}
}

////////////////////////CLASS
func (chr *Character) SetClass() {
	classInt := 0
	classList := []string{Warrior, Expert, Mage, AdventurerEM, AdventurerEW, AdventurerMW}
	switch chr.FlagAuto {
	case true:
		classInt = chr.Dice.RollNext("1d6").DM(-1).Sum()
	case false:
		classInt = chooseOption("Choose your Class:", classList)
	}
	chr.Class = asset.NewClass(classList[classInt])
}

////////////////////////FOCI
func (chr *Character) SetFoci() {
	chr.Foci = make(map[string]asset.Foci)
	fociExpected := []string{"Any"}
	switch chr.Class.Name() { //TODO: тутже добавдять фокусы за уровни
	case Warrior:
		fociExpected = append(fociExpected, Warrior)
	case Expert:
		fociExpected = append(fociExpected, Expert)
	case AdventurerEW:
		fociExpected = append(fociExpected, Warrior)
		fociExpected = append(fociExpected, Expert)
	case AdventurerEM:
		fociExpected = append(fociExpected, Expert)
	case AdventurerMW:
		fociExpected = append(fociExpected, Warrior)
	}
	for _, fociType := range fociExpected {
		fList := []string{}
		switch fociType {
		default:
			fList = allFociList()
		case Warrior:
			fList = warriorFociList()
		case Expert:
			fList = expertFociList()
		}
		fName := ""
		switch chr.FlagAuto {
		case true:
			fName = chr.Dice.RollFromList(fList)
		case false:
			pick := chooseOption("Pick Foci:", fList)
			fName = fList[pick]
		}
		if _, ok := chr.Foci[fName]; ok {
			chr.Foci[fName].UpgradeFoci()
			fmt.Println(chr.Sheet())
			continue
		}
		newFoci := asset.NewFoci(fName)
		chr.Foci[fName] = newFoci
		chr.Train(newFoci.BonusSkill())
		fmt.Println(chr.Sheet())
	}

}

////////////////////////MAGIC

func (chr *Character) SetMagicTraditions() {
	mtp := 0
	switch chr.Class.Name() {
	case Mage:
		mtp = 2
	case AdventurerEM, AdventurerMW:
		mtp = 1
	}
	traditionsPicked := []string{}
	for mtp > 0 {
		if mtp == 2 {
			fmt.Println("As a Full Mage you can select two traditions. Picking the same tradions twise will boost it's effect...")
		}
		options := traditionsListDynamic(mtp)
		switch chr.FlagAuto {
		case true:
			traditionsPicked = append(traditionsPicked, chr.Dice.RollFromList(options))
		case false:
			chosen := chooseOption("Select tradition:", options)
			traditionsPicked = append(traditionsPicked, options[chosen])
		}
		mtp--
	}
	if len(traditionsPicked) == 0 {
		return
	}
	chr.Tradition = asset.NewTradition(traditionsPicked)
	avalableSpells := []string{}
	if strings.Contains(chr.Tradition.Name(), HighMage) {
		chr.Train(Magic)
		traditionSpells := magic.FilterSpellsByTradition("High Mage", 1)
		for _, sp := range traditionSpells {
			utils.AppendUniqueStr(avalableSpells, sp)
		}
	}
}

/*
switch Tradition {
case Full High Mage:
	add full bonus
case Full Elementalist:
	add full bonus
case Full Necromancer:
	add full bonus
default:
	if Tradition Contains {
		add partial bonus
	}
}
learn spells:
for i := 0; i < chr.Level {
	maxLevel := from chr.Level and Tradition
	constructList(Tradition)
	pick := from chr.Tradition

}

*/

func traditionsListDynamic(mtpLeft int) []string {
	switch mtpLeft {
	default:
		return []string{"ERROR"}
	case 2:
		return []string{
			HighMage,
			Elementalist,
			Necromancer,
		}
	case 1:
		return []string{
			HighMage,
			Elementalist,
			Necromancer,
			Healer,
			Vowed,
		}
	}
}

////////////////////////HELPERS

func clearTerm() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func chooseOption(q string, ans []string) int {
	a, err := user.ChooseOne(q, ans)
	for err != nil {
		fmt.Println(err.Error())
		a, err = user.ChooseOne(q, ans)
	}
	return a
}

func allAttributes() []string {
	return []string{
		STR,
		DEX,
		CON,
		INT,
		WIS,
		CHA,
	}
}

func allSkills() []string {
	return []string{
		Administer,
		Connect,
		Convince,
		Craft,
		Exert,
		Heal,
		Know,
		Lead,
		Magic,
		Notice,
		Perform,
		Pray,
		Punch,
		Ride,
		Sail,
		Shoot,
		Sneak,
		Stab,
		Survive,
		Trade,
		Work,
	}
}

func specialistSkills() []string {
	return []string{
		Administer,
		Connect,
		Convince,
		Craft,
		Exert,
		Heal,
		Know,
		Lead,
		Notice,
		Perform,
		Pray,
		Ride,
		Sail,
		Sneak,
		Survive,
		Trade,
		Work,
	}
}

func allBackgrounds() []string {
	return []string{
		Artisan,
		Barbarian,
		Carter,
		Courtesan,
		Criminal,
		Hunter,
		Laborer,
		Merchant,
		Noble,
		Nomad,
		Peasant,
		Performer,
		Physician,
		Priest,
		Sailor,
		Scholar,
		Slave,
		Soldier,
		Thug,
		Wanderer,
	}
}

func allFociList() []string {
	return []string{
		"Alert",
		"Armored Magic",
		"Armsmaster",
		"Artisan",
		"Assassin",
		"Authority",
		"Close Combatant",
		"Connected",
		"Cultured",
		"Die Hard",
		"Deadeye",
		"Dealmaker",
		"Developed Attribute",
		"Diplomatic Grace",
		"Gifted Chirurgeon",
		"Henchkeeper",
		"Impervious Defense",
		"Impostor",
		"Lucky",
		"Nullifier",
		"Poisoner",
		"Polymath",
		"Rider",
		"Shocking Assault",
		"Sniper's Eye",
		"Special Origin",
		"Specialist",
		"Spirit Familiar",
		"Trapmaster",
		"Unarmed Combatant",
		"Unique Gift",
		"Valiant Defender",
		"Well Met",
		"Whirlwind Assault",
		"Xenoblooded",
	}
}

func warriorFociList() []string {
	return []string{
		"Alert",
		"Armsmaster",
		"Assassin",
		"Close Combatant",
		"Die Hard",
		"Deadeye",
		"Impervious Defense",
		"Rider",
		"Shocking Assault",
		"Sniper's Eye",
		"Unarmed Combatant",
		"Valiant Defender",
		"Whirlwind Assault",
	}
}

func expertFociList() []string {
	return []string{
		"Alert",
		"Artisan",
		"Authority",
		"Connected",
		"Cultured",
		"Die Hard",
		"Dealmaker",
		"Diplomatic Grace",
		"Gifted Chirurgeon",
		"Henchkeeper",
		"Impervious Defense",
		"Impostor",
		"Lucky",
		"Poisoner",
		"Rider",
		"Trapmaster",
		"Well Met",
	}
}
