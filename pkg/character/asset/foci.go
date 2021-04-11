package asset

import (
	"strconv"
	"strings"
)

const (
	Alert              = "Alert"
	ArmoredMagic       = "Armored Magic"
	Armsmaster         = "Armsmaster"
	Artisan            = "Artisan"
	Assassin           = "Assassin"
	Authority          = "Authority"
	CloseCombatant     = "Close Combatant"
	Connected          = "Connected"
	Cultured           = "Cultured"
	DieHard            = "Die Hard"
	Deadeye            = "Deadeye"
	Dealmaker          = "Dealmaker"
	DevelopedAttribute = "Developed Attribute"
	DiplomaticGrace    = "Diplomatic Grace"
	GiftedChirurgeon   = "Gifted Chirurgeon"
	Henchkeeper        = "Henchkeeper"
	ImperviousDefense  = "Impervious Defense"
	Impostor           = "Impostor"
	Lucky              = "Lucky"
	Nullifier          = "Nullifier"
	Poisoner           = "Poisoner"
	Polymath           = "Polymath"
	Rider              = "Rider"
	ShockingAssault    = "Shocking Assault"
	SnipersEye         = "Sniper’s Eye"
	SpecialOrigin      = "Special Origin"
	Specialist         = "Specialist"
	SpiritFamiliar     = "Spirit Familiar"
	Trapmaster         = "Trapmaster"
	UnarmedCombatant   = "Unarmed Combatant"
	UniqueGift         = "Unique Gift"
	ValiantDefender    = "Valiant Defender"
	WellMet            = "Well Met"
	WhirlwindAssault   = "Whirlwind Assault"
	Xenoblooded        = "Xenoblooded"
)

type Foci interface {
	NameFoci() string
	Level() int
	GetDescription() string
	BonusSkill() string
	UpgradeFoci()
}

func NewFoci(name string) Foci {
	fc := Asset{}
	fc.Name = name
	if strings.Contains(name, Specialist) || strings.Contains(name, DevelopedAttribute) {
		fc.Name = name + " (not implemented fully)"
	}
	fc.Rank = 1
	fc.Bonus = acquisitionBonusSkill(name)
	fc.Description = "[TODO: " + name + " foci descrioption based on level " + strconv.Itoa(fc.Level()) + "]"
	return &fc
}

func (a *Asset) BonusSkill() string {
	return a.Bonus
}

func (a *Asset) GetDescription() string {
	return a.Description
}

func (a *Asset) UpgradeFoci() {
	if a.Rank == 1 { //if upgradeable - upgrade
		a.Rank = 2
	}
}

func (a *Asset) NameFoci() string {
	return a.Name + " (lvl. " + strconv.Itoa(a.Rank) + ")"
}

func acquisitionBonusSkill(fociName string) string {
	abs := make(map[string]string)
	abs[Alert] = "Notice"
	abs[Armsmaster] = "Stab"
	abs[Artisan] = "Craft"
	abs[Assassin] = "Sneak"
	abs[Authority] = "Lead"
	abs[CloseCombatant] = "Any Combat"
	abs[Connected] = "Connect"
	abs[Cultured] = "Connect"
	abs[Deadeye] = "Shoot"
	abs[Dealmaker] = "Trade"
	abs[DiplomaticGrace] = "Convince"
	abs[GiftedChirurgeon] = "Heal"
	abs[Henchkeeper] = "Lead"
	abs[Impostor] = "Perform or Sneak"
	abs[Poisoner] = "Heal"
	abs[Polymath] = "Any Skill"
	abs[Rider] = "Ride"
	abs[ShockingAssault] = "Punch or Stab"
	abs[SnipersEye] = "Shoot"
	abs[Specialist] = "Specialist" //TODO:нужно как-то вставить уточнялку хочет ли игрок прокачать имеющийся фокус или создать новый
	abs[Trapmaster] = "Notice"
	abs[UnarmedCombatant] = "Punch"
	abs[ValiantDefender] = "Punch or Stab"
	abs[WhirlwindAssault] = "Stab"

	return abs[fociName]
}
