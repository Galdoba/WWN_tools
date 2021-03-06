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
	DualMage         = "Dual Mage"
	FullHighMage     = "Full High Mage"
	FullElementalist = "Full Elementalist"
	FullNecromancer  = "Full Necromancer"
	HighMage         = "Partial High Mage"
	Elementalist     = "Partial Elementalist"
	Necromancer      = "Partial Necromancer"
	Healer           = "Healer"
	Vowed            = "Vowed"
)

type Character struct {
	Name       string
	Race       string
	FlagAuto   bool
	Dice       *dice.Dicepool
	Level      int
	Stat       map[string]string
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
	chr.Stat = make(map[string]string)
	for _, val := range []string{"HP", "Effort", "BA", "MA", "RA", "Init", "SPhy", "SEva", "SMen", "Lk"} {
		chr.Stat[val] = "  "
	}
	chr.Attribute = make(map[string]asset.Attribute)
	atrArrayStr := []string{STR, DEX, CON, INT, WIS, CHA}
	for _, val := range atrArrayStr {
		chr.Attribute[val] = asset.NewAttribute(val)
	}

	return chr
}

func (chr *Character) SetAttributes() {
	atrArrayStr := []string{STR, DEX, CON, INT, WIS, CHA}
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
	validOptions := allBackgrounds()
	validOptions = append(validOptions, "Naaah... just roll it")
	chsen := chr.ChooseOption("Select background for a character:\n*check description and bonuses in CRB p.12-17", validOptions)
	if chsen == "Naaah... just roll it" {
		chsen = chr.Dice.RollFromList(allBackgrounds())
	}
	chr.Background = asset.NewBackground(chsen)
}

////////////////////////SKILLS
func (chr *Character) SetSkills() {
	chr.Skill = make(map[string]asset.Skill)
	freeSkill := chr.Background.FreeSkill()
	method := ""
	chr.Train(freeSkill)
	options := []string{"Gain the background???s listed quick skills", "Pick two skills from the background???s Learning table", "Roll three times, splitting the rolls as you wish between the Growth and Learning tables for your background"}
	method = chr.ChooseOption("What will you do?", options)
	switch method {
	case options[0]:
		quickSkills := chr.Background.QuickSkills()
		for _, skl := range quickSkills {
			chr.Train(skl)
		}
		return
	case options[1]:
		learn := chr.Background.Learning()
		learn = pickValid(learn)
		chosen := []string{}
		chosen = append(chosen, chr.ChooseOption("Chose first skill:", learn))
		chosen = append(chosen, chr.ChooseOption("Chose second skill:", learn))
		for _, skl := range chosen {
			chr.Train(skl)
		}
		return
	case options[2]:
		chosen := []string{}
		for len(chosen) < 3 {
			i := len(chosen)
			chosen = append(chosen, chr.ChooseOption("Select from table :"+chr.Background.Tables(), []string{"Growth", "Learning"}))
			switch chosen[i] {
			case "Growth":
				chr.Train(chr.Dice.RollFromList(chr.Background.Growth()))
			case "Learning":
				chr.Train(chr.Dice.RollFromList(chr.Background.Learning()))
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
	switch chr.Class.Name() { //TODO: ?????????? ?????????????????? ???????????? ???? ????????????
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
	chr.benefitFromTraditions()

}

func (chr *Character) maxLevelSpell() int {
	if chr.Tradition == nil {
		return 0
	}
	switch chr.Tradition.SATable() {
	case "FHM", "FEM", "FNM":
		return ((chr.Level - 1) / 2) + 1
	case "PHM", "PEM", "PNM":
		return ((chr.Level - 1) / 4) + 1

	}
	if chr.Tradition.DualMage() {
		return ((chr.Level - 1) / 3) + 1
	}
	return 0
}

func (chr *Character) benefitFromTraditions() {
	maxLevel := chr.maxLevelSpell()
	addSpells := false
	spellList := magic.FilterSpellsByTradition(magic.TraditionHighMagic, maxLevel)
	if strings.Contains(chr.Tradition.Name(), "High Mage") {
		chr.Train(Magic)
		addSpells = true
	}
	if strings.Contains(chr.Tradition.Name(), "Elementalist") {
		chr.Train(Magic)
		spellList = append(spellList, magic.FilterSpellsByTradition(magic.TraditionElementalist, maxLevel)...)
		addSpells = true
		chr.Tradition.AddArt("Elemental Resilence")
		chr.Tradition.AddArt("Elemental Sparks")
	}
	if strings.Contains(chr.Tradition.Name(), "Necromancer") {
		chr.Train(Magic)
		spellList = append(spellList, magic.FilterSpellsByTradition(magic.TraditionNecromancer, maxLevel)...)
		addSpells = true
	}
	if addSpells {
		for len(chr.Tradition.AvailableSpells()) < 4 {
			spellList = cleanOptions(spellList, chr.Tradition.AvailableSpells())
			switch chr.FlagAuto {
			case true:
				chr.Tradition.AddSpell(chr.Dice.RollFromList(spellList))
			case false:
				n := strconv.Itoa(len(chr.Tradition.AvailableSpells()) + 1)
				chr.Tradition.AddSpell(spellList[chooseOption("Select Spell ("+n+"/4):", spellList)])
				fmt.Print(chr.Sheet())
			}

		}
	}
	if strings.Contains(chr.Tradition.Name(), "Healer") {
		chr.Train(Heal)
		chr.Tradition.AddArt("Healing Touch")
	}
	if strings.Contains(chr.Tradition.Name(), "Vowed") {
		chr.Train("Non-Combat")
		chr.Tradition.AddArt("Martial Style")
		chr.Tradition.AddArt("Unarmed Might")
		chr.Tradition.AddArt("Unarmored Defense")
	}
	artList := []string{} //?????????? ?????????????????????? ?? ???????? - ?? ?????????????????????? ???????? ???????? ?????????? ??????????
	//					  //???????? ?????????? ?????????????? ???? ?????????????????? ????????????.
	if strings.Contains(chr.Tradition.Name(), "High Mage") {
		artList = magic.ArtListByTradition(magic.TraditionHighMagic)
		switch chr.FlagAuto {
		case true:
			chr.Tradition.AddArt(chr.Dice.RollFromList(artList))
		case false:
			chr.Tradition.AddArt(artList[chooseOption("Choose Art:", artList)])
		}
		fmt.Print(chr.Sheet())
	}
	if strings.Contains(chr.Tradition.Name(), "Elementalist") {
		artList = magic.ArtListByTradition(magic.TraditionElementalist)
		switch chr.FlagAuto {
		case true:
			chr.Tradition.AddArt(chr.Dice.RollFromList(artList))
		case false:
			chr.Tradition.AddArt(artList[chooseOption("Choose Art:", artList)])
		}
		fmt.Print(chr.Sheet())
	}
	if strings.Contains(chr.Tradition.Name(), "Necromancer") {
		artList = magic.ArtListByTradition(magic.TraditionNecromancer)
		switch chr.FlagAuto {
		case true:
			chr.Tradition.AddArt(chr.Dice.RollFromList(artList))
		case false:
			chr.Tradition.AddArt(artList[chooseOption("Choose Art:", artList)])
		}
		fmt.Print(chr.Sheet())
	}
	if strings.Contains(chr.Tradition.Name(), "Healer") {
		artList = magic.ArtListByTradition(magic.TraditionHealer)
		switch chr.FlagAuto {
		case true:
			chr.Tradition.AddArt(chr.Dice.RollFromList(artList))
		case false:
			chr.Tradition.AddArt(artList[chooseOption("Choose Art:", artList)])
		}
		fmt.Print(chr.Sheet())
	}
	if strings.Contains(chr.Tradition.Name(), "Vowed") {
		artList = magic.ArtListByTradition(magic.TraditionVowed)
		switch chr.FlagAuto {
		case true:
			chr.Tradition.AddArt(chr.Dice.RollFromList(artList))
		case false:
			chr.Tradition.AddArt(artList[chooseOption("Choose Art:", artList)])
		}
		fmt.Print(chr.Sheet())
	}
	if strings.Contains(chr.Tradition.Name(), FullHighMage) {
		artList = magic.ArtListByTradition(magic.TraditionHighMagic)
		artList = cleanOptions(artList, chr.Tradition.AvailableArts())
		switch chr.FlagAuto {
		case true:
			chr.Tradition.AddArt(chr.Dice.RollFromList(artList))
		case false:
			chr.Tradition.AddArt(artList[chooseOption("Choose Art:", artList)])
		}
		fmt.Print(chr.Sheet())
	}

}

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

////////////////////////FINAL TOUCH
func (chr *Character) SetHitPoints() {
	hdMod := chr.Attribute[CON].Modifer()
	chr.Stat["BA"] = " 0"
	switch chr.Class.Name() {
	case Warrior, AdventurerEW, AdventurerMW:
		hdMod = hdMod + (chr.Level * 2)
		chr.Stat["BA"] = "+1"
	case Expert, AdventurerEM:
		hdMod = hdMod + (chr.Level * 0)
	case Mage:
		hdMod = hdMod + (chr.Level * -1)
	}
	if _, ok := chr.Foci[asset.DieHard]; ok {
		hdMod += (2 * chr.Level)
	}
	hp := chr.Dice.RollNext("1d6").DM(hdMod).Sum()
	if hp < 1 {
		hp = 1
	}
	chr.Stat["HP"] = strconv.Itoa(hp)
	for len(chr.Stat["HP"]) < 2 {
		chr.Stat["HP"] = " " + chr.Stat["HP"]
	}
	chr.Stat["Init"] = strconv.Itoa(chr.Attribute[DEX].Modifer())
	chr.Stat["MA"] = strconv.Itoa(chr.Attribute[STR].Modifer())
	chr.Stat["RA"] = strconv.Itoa(chr.Attribute[DEX].Modifer())
	chr.Stat["Effort"] = " 0"
	if _, ok := chr.Skill[Magic]; ok {
		chr.Stat["Effort"] = strconv.Itoa(utils.Max(chr.Attribute[INT].Modifer(), chr.Attribute[CHA].Modifer()) + 1 + chr.Skill[Magic].Level())
	}

	chr.Stat["SPhy"] = strconv.Itoa(15 - utils.Max(chr.Attribute[STR].Modifer(), chr.Attribute[CON].Modifer()))
	chr.Stat["SEva"] = strconv.Itoa(15 - utils.Max(chr.Attribute[DEX].Modifer(), chr.Attribute[INT].Modifer()))
	chr.Stat["SMen"] = strconv.Itoa(15 - utils.Max(chr.Attribute[WIS].Modifer(), chr.Attribute[CHA].Modifer()))
	chr.Stat["Lk"] = strconv.Itoa(15)
	for k := range chr.Stat {
		for len(chr.Stat[k]) < 2 {
			chr.Stat[k] = " " + chr.Stat[k]
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

func cleanOptions(options, picked []string) []string {
	valid := []string{}
	for _, val := range options {
		met := false
		for _, test := range picked {
			if val == test {
				met = true
			}
		}
		if !met {
			valid = append(valid, val)
		}
	}
	return valid
}

func (chr *Character) ChooseOption(msg string, options []string) string {
	decision := "NOT MADE"
	switch chr.FlagAuto {
	case true:
		decision = chr.Dice.RollFromList(options)
	case false:
		ch := chooseOption(msg, options)
		decision = options[ch]
	}
	return decision
}
