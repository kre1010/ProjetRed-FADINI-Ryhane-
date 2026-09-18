package red

import (
	"fmt"
)

type Character struct {
	Name  string
	Class string
	Level int
	maxHp int
	Hp    int
	inv   map[string]int
}

func (c *Character) initCharacter(name string, class string) {
	var inputclass string
	c.Name = name
	c.Class = class
	fmt.Scanln(&inputclass)
	switch c.Class {
	case "1":
		c.maxHp = 100
		c.Hp = c.maxHp / 2
		c.inv = map[string]int{"Sniper Rifle": 1, "Heavy Bullets": 7}
	case "2":
		c.maxHp = 130
		c.Hp = c.maxHp / 2
		c.inv = map[string]int{"Assault Rifle": 1, "Medium bullets": 44}
	case "3":
		c.maxHp = 150
		c.Hp = c.maxHp / 2
		c.inv = map[string]int{"Shotgun": 1, "Shotgun shells": 9}
	case "???":
		c.maxHp = 1
		c.Hp = 1
		c.inv = map[string]int{"Atomic": 999}
	}
	fmt.Println("Choisis ton charactère:", '\n')
}
