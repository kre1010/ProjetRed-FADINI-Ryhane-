package character

import (
	"fmt"
)

type Character struct {
	Name    string
	Class   string
	Level   int
	maxHp   int
	Hp      int
	Inv     map[string]int
	Argent  int
	buyItem int
}

func (c *Character) MenuInitCharacter() {
	fmt.Println("Début de la création de votre personnage...")
	fmt.Println("Choisissez votre nom :")
	var nameInput string
	fmt.Scan(&nameInput)
	fmt.Println("Choisissez votre classe :")
	fmt.Println("\t 1 - Sniper : HP = 50, max HP = 100, Sniper rifle, Heavy Bullet = 7")
	fmt.Println("\t 2 - Assault : HP = 50, max HP = 100, Assault rifle, Medium bullets = 44")
	fmt.Println("\t 3 - Sniper : HP = 50, max HP = 100, Shotgun, Shotgun shells = 9")
	var classInput string
	fmt.Scan(&classInput)

	c.InitCharacter(nameInput, classInput)
	fmt.Println("fin de la création de votre personnage....")
}

func (c *Character) InitCharacter(name string, class string) {

	c.Name = name
	c.Class = class

	switch c.Class {
	case "1":
		c.Level = 1
		c.maxHp = 100
		c.Hp = c.maxHp / 2
		c.Inv = map[string]int{"Sniper Rifle": 1, "Heavy Bullets": 7}
	case "2":
		c.Level = 1
		c.maxHp = 130
		c.Hp = c.maxHp / 2
		c.Inv = map[string]int{"Assault Rifle": 1, "Medium bullets": 44}
	case "3":
		c.Level = 1
		c.maxHp = 150
		c.Hp = c.maxHp / 2
		c.Inv = map[string]int{"Shotgun": 1, "Shotgun shells": 9}
	case "???":
		c.Level = 1
		c.maxHp = 1
		c.Hp = 1
		c.Inv = map[string]int{"Atomic": 999}
	}
}
