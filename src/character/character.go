package character

import (
	"fmt"
)

type Character struct {
	Name   string
	Class  int
	Level  int
	Hp     int
	MaxHp  int
	Inv    map[string]int
	MaxInv int
	Argent int
	skill []string
	Attaque int 
}

func (c *Character) MenuInitCharacter() {

	fmt.Println("Début de la création de votre personnage...")

	fmt.Println("Choisissez votre nom :")

	var nameInput string
	fmt.Scan(&nameInput)

	for {
		fmt.Println("Choisissez votre classe :")
		fmt.Println("\t1 - Sniper : HP = 75, max HP = 150, Sniper rifle, Heavy Bullet = 7")
		fmt.Println("\t2 - Assault : HP = 65, max HP = 130, Assault rifle, Medium bullets = 44")
		fmt.Println("\t3 - Shotgunner : HP = 75, max HP = 150, Shotgun, Shotgun shells = 9")

		var classInput int
		fmt.Scan(&classInput)

		if classInput >= 1 && classInput <= 3 {
			c.InitCharacter(nameInput, classInput)
			break
		}
		fmt.Println("Choix invalide, veuillez choisir 1, 2 ou 3.")
	}

	fmt.Println("Fin de la création de votre personnage....")
}

func (c *Character) InitCharacter(name string, class int) {

	c.Name = name
	c.Class = class
	c.MaxInv = 10

	switch c.Class {

	case 1:
		c.Level = 1
		c.MaxHp = 100
		c.Hp = c.MaxHp / 2
		c.Attaque = 50 
		c.Inv = map[string]int{
			"Sniper Rifle":   1,
			"Heavy Bullets": 7,
		}
		c.Argent = 100

	case 2:
		c.Level = 1
		c.MaxHp = 130
		c.Hp = c.MaxHp / 2
		c.Attaque = 30
		c.Inv = map[string]int{
			"Assault Rifle":  1,
			"Medium bullets": 44,
		}
		c.Argent = 100

	case 3:
		c.Level = 1
		c.MaxHp = 150
		c.Hp = c.MaxHp / 2
		c.Attaque = 40
		c.Inv = map[string]int{
			"Shotgun":        1,
			"Shotgun shells": 9,
		}
		c.Argent = 100
	}
}

