package asset

import "strconv"

type Background interface {
	NameBKG() string
	FreeSkill() string
	QuickSkills() []string
	Growth() []string
	Learning() []string
	Tables() string
}

func NewBackground(name string) Background {
	bkg := Asset{}
	bkg.Group = TypeAttribute
	bkg.Name = name
	return &bkg
}

func (a *Asset) NameBKG() string {
	return a.Name
}

func (a *Asset) FreeSkill() string {
	switch a.Name {
	default:
		return "ERROR"
	case "Artisan":
		return "Craft"
	case "Barbarian":
		return "Survive"
	case "Carter":
		return "Ride"
	case "Courtesan":
		return "Perform"
	case "Criminal":
		return "Sneak"
	case "Hunter":
		return "Shoot"
	case "Laborer":
		return "Work"
	case "Merchant":
		return "Trade"
	case "Noble":
		return "Lead"
	case "Nomad":
		return "Ride"
	case "Peasant":
		return "Exert"
	case "Performer":
		return "Perform"
	case "Physician":
		return "Heal"
	case "Priest":
		return "Pray"
	case "Sailor":
		return "Sail"
	case "Scholar":
		return "Know"
	case "Slave":
		return "Sneak"
	case "Soldier":
		return "Any Combat"
	case "Thug":
		return "Any Combat"
	case "Wanderer":
		return "Survive"
	}
}

func (a *Asset) QuickSkills() []string {
	switch a.Name {
	default:
		return []string{"ERROR"}
	case "Artisan":
		return []string{
			"Trade",
			"Connect",
		}
	case "Barbarian":
		return []string{
			"Any Combat",
			"Notice",
		}
	case "Carter":
		return []string{
			"Connect",
			"Any Combat",
		}
	case "Courtesan":
		return []string{
			"Notice",
			"Connect",
		}
	case "Criminal":
		return []string{
			"Connect",
			"Convince",
		}
	case "Hunter":
		return []string{
			"Survive",
			"Sneak",
		}
	case "Laborer":
		return []string{
			"Connect",
			"Exert",
		}
	case "Merchant":
		return []string{
			"Convince",
			"Connect",
		}
	case "Noble":
		return []string{
			"Connect",
			"Administer",
		}
	case "Nomad":
		return []string{
			"Survive",
			"Any Combat",
		}
	case "Peasant":
		return []string{
			"Sneak",
			"Survive",
		}
	case "Performer":
		return []string{
			"Convince",
			"Connect",
		}
	case "Physician":
		return []string{
			"Know",
			"Notice",
		}
	case "Priest":
		return []string{
			"Convince",
			"Know",
		}
	case "Sailor":
		return []string{
			"Exert",
			"Notice",
		}
	case "Scholar":
		return []string{
			"Heal",
			"Administer",
		}
	case "Slave":
		return []string{
			"Survive",
			"Exert",
		}
	case "Soldier":
		return []string{
			"Exert",
			"Survive",
		}
	case "Thug":
		return []string{
			"Convince",
			"Connect",
		}
	case "Wanderer":
		return []string{
			"Sneak",
			"Notice",
		}
	}
}

func (a *Asset) Growth() []string {
	switch a.Name {
	default:
		return []string{"ERROR"}
	case "Artisan":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Mental",
			"Exert",
			"Any Skill",
		}
	case "Barbarian":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Mental",
			"Exert",
			"Any Skill",
		}
	case "Carter":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Mental",
			"Connect",
			"Any Skill",
		}
	case "Courtesan":
		return []string{
			"+1 Any Stat",
			"+2 Mental",
			"+2 Mental",
			"+2 Physical",
			"Connect",
			"Any Skill",
		}
	case "Criminal":
		return []string{
			"+1 Any Stat",
			"+2 Mental",
			"+2 Physical",
			"+2 Mental",
			"Connect",
			"Any Skill",
		}
	case "Hunter":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Mental",
			"Exert",
			"Any Skill",
		}
	case "Laborer":
		return []string{
			"+1 Any Stat",
			"+1 Any Stat",
			"+1 Any Stat",
			"+1 Any Stat",
			"Exert",
			"Any Skill",
		}
	case "Merchant":
		return []string{
			"+1 Any Stat",
			"+2 Mental",
			"+2 Mental",
			"+2 Mental",
			"Connect",
			"Any Skill",
		}
	case "Noble":
		return []string{
			"+1 Any Stat",
			"+2 Mental",
			"+2 Mental",
			"+2 Mental",
			"Connect",
			"Any Skill",
		}
	case "Nomad":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Mental",
			"Exert",
			"Any Skill",
		}
	case "Peasant":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Physical",
			"Exert",
			"Any Skill",
		}
	case "Performer":
		return []string{
			"+1 Any Stat",
			"+2 Mental",
			"+2 Physical",
			"+2 Physical",
			"Connect",
			"Any Skill",
		}
	case "Physician":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Mental",
			"+2 Mental",
			"Connect",
			"Any Skill",
		}
	case "Priest":
		return []string{
			"+1 Any Stat",
			"+2 Mental",
			"+2 Physical",
			"+2 Mental",
			"Connect",
			"Any Skill",
		}
	case "Sailor":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Mental",
			"Exert",
			"Any Skill",
		}
	case "Scholar":
		return []string{
			"+1 Any Stat",
			"+2 Mental",
			"+2 Mental",
			"+2 Mental",
			"Connect",
			"Any Skill",
		}
	case "Slave":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Mental",
			"Exert",
			"Any Skill",
		}
	case "Soldier":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Physical",
			"Exert",
			"Any Skill",
		}
	case "Thug":
		return []string{
			"+1 Any Stat",
			"+2 Mental",
			"+2 Physical",
			"+2 Physical",
			"Connect",
			"Any Skill",
		}
	case "Wanderer":
		return []string{
			"+1 Any Stat",
			"+2 Physical",
			"+2 Physical",
			"+2 Mental",
			"Exert",
			"Any Skill",
		}
	}
}

func (a *Asset) Learning() []string {
	switch a.Name {
	default:
		return []string{"ERROR"}
	case "Artisan":
		return []string{
			"Connect",
			"Convince",
			"Craft",
			"Craft",
			"Exert",
			"Know",
			"Notice",
			"Trade",
		}
	case "Barbarian":
		return []string{
			"Any Combat",
			"Connect",
			"Exert",
			"Lead",
			"Notice",
			"Punch",
			"Sneak",
			"Survive",
		}
	case "Carter":
		return []string{
			"Any Combat",
			"Connect",
			"Craft",
			"Exert",
			"Notice",
			"Ride",
			"Survive",
			"Trade",
		}
	case "Courtesan":
		return []string{
			"Any Combat",
			"Connect",
			"Convince",
			"Exert",
			"Notice",
			"Perform",
			"Survive",
			"Trade",
		}
	case "Criminal":
		return []string{
			"Administer",
			"Any Combat",
			"Connect",
			"Convince",
			"Exert",
			"Notice",
			"Sneak",
			"Trade",
		}
	case "Hunter":
		return []string{
			"Any Combat",
			"Exert",
			"Heal",
			"Notice",
			"Ride",
			"Shoot",
			"Sneak",
			"Survive",
		}
	case "Laborer":
		return []string{
			"Administer",
			"Any Skill",
			"Connect",
			"Convince",
			"Craft",
			"Exert",
			"Ride",
			"Work",
		}
	case "Merchant":
		return []string{
			"Administer",
			"Any Combat",
			"Connect",
			"Convince",
			"Craft",
			"Know",
			"Notice",
			"Trade",
		}
	case "Noble":
		return []string{
			"Administer",
			"Any Combat",
			"Connect",
			"Convince",
			"Know",
			"Lead",
			"Notice",
			"Ride",
		}
	case "Nomad":
		return []string{
			"Any Combat",
			"Connect",
			"Exert",
			"Lead",
			"Notice",
			"Ride",
			"Survive",
			"Trade",
		}
	case "Peasant":
		return []string{
			"Connect",
			"Exert",
			"Craft",
			"Notice",
			"Sneak",
			"Survive",
			"Trade",
			"Work",
		}
	case "Performer":
		return []string{
			"Any Combat",
			"Connect",
			"Exert",
			"Notice",
			"Perform",
			"Perform",
			"Sneak",
			"Convince",
		}
	case "Physician":
		return []string{
			"Administer",
			"Connect",
			"Craft",
			"Heal",
			"Know",
			"Notice",
			"Convince",
			"Trade",
		}
	case "Priest":
		return []string{
			"Administer",
			"Connect",
			"Know",
			"Lead",
			"Heal",
			"Convince",
			"Pray",
			"Pray",
		}
	case "Sailor":
		return []string{
			"Any Combat",
			"Connect",
			"Craft",
			"Exert",
			"Heal",
			"Notice",
			"Perform",
			"Sail",
		}
	case "Scholar":
		return []string{
			"Administer",
			"Heal",
			"Craft",
			"Know",
			"Notice",
			"Perform",
			"Pray",
			"Convince",
		}
	case "Slave":
		return []string{
			"Administer",
			"Any Combat",
			"Any Skill",
			"Convince",
			"Exert",
			"Sneak",
			"Survive",
			"Work",
		}
	case "Soldier":
		return []string{
			"Any Combat",
			"Any Combat",
			"Exert",
			"Lead",
			"Notice",
			"Ride",
			"Sneak",
			"Survive",
		}
	case "Thug":
		return []string{
			"Any Combat",
			"Any Combat",
			"Connect",
			"Convince",
			"Exert",
			"Notice",
			"Sneak",
			"Survive",
		}
	case "Wanderer":
		return []string{
			"Any Combat",
			"Connect",
			"Notice",
			"Perform",
			"Ride",
			"Sneak",
			"Survive",
			"Work",
		}
	}
}

func (a *Asset) Tables() string {
	tb := "\n  d6 Growth       d8 Learning\n"
	gr := a.Growth()
	ln := a.Learning()
	for i := 0; i < 8; i++ {
		switch i {
		case 0, 1, 2, 3, 4, 5:
			grStr := strconv.Itoa(i+1) + "  " + gr[i]
			for len(grStr) < 16 {
				grStr += " "
			}
			lnStr := strconv.Itoa(i+1) + "  " + ln[i]
			for len(lnStr) < 16 {
				lnStr += " "
			}
			tb += "  " + grStr + lnStr + "\n"
		case 6, 7:
			grStr := "                "
			lnStr := strconv.Itoa(i+1) + "  " + ln[i]
			for len(lnStr) < 16 {
				lnStr += " "
			}
			tb += "  " + grStr + lnStr + "\n"
		}
	}
	return tb
}
