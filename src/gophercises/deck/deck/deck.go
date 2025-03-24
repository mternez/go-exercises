package deck

type family interface {
	GetName() string
	SetName(name string)
}

type card interface {
	GetName() string
	SetName(name string)
	GetValue() int
	SetValue(value int)
	GetFamily() family
	SetFamily(family family)
}

type Family struct {
	name string
}

func (c *Family) GetName() string {
	return c.name
}

func (c *Family) SetName(name string) {
	c.name = name
}

type Card struct {
	name   string
	value  int
	family family
}

func (c *Card) GetName() string {
	return c.name
}

func (c *Card) SetName(name string) {
	c.name = name
}

func (c *Card) GetValue() int {
	return c.value
}

func (c *Card) SetValue(value int) {
	c.value = value
}

func (c *Card) GetFamily() family {
	return c.family
}

func (c *Card) SetFamily(family family) {
	c.family = family
}
