package potvie

import c "red/src/character"

func Medkit(c c.Character) {
	c.Hp += 20

	if c.Hp > c.maxHp {
		c.Hp = c.maxHp
	}
}
