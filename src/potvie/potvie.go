package potvie

import "red/src/character"

func Medkit(c *character.Character) {
	c.Hp += 20
	
	if c.Hp > c.MaxHp {
		c.Hp = c.MaxHp
	}
}

