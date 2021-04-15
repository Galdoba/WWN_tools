package magic

type construct struct {
	typeOfConstruct int
	tradition       string
	level           int
	name            string
	description     string
}

const (
	ConstructTypeArt   = 1
	ConstructTypeSpell = 2
)

func (c *construct) TypeIs() string {
	switch c.typeOfConstruct {
	default:
		panic("UNKNOWN MAGIC.CONSTRUCT Type")
	case 1:
		return "Art"
	case 2:
		return "Spell"
	}
}

func (c *construct) Tradition() string {
	return c.tradition
}

func (c *construct) Level() int {
	return c.level
}

func (c *construct) Name() string {
	return c.name
}

func (c *construct) Description() string {
	return c.description
}
