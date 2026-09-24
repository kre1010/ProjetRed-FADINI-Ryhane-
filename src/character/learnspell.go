package character

import (
	"fmt"
)

func (c *Character) DisplaySpell() {
	fmt.Println("Voulez-vous apprendre à lancer une grenade ?")
	fmt.Println("1. Oui")
	fmt.Println("2. Non")

	var choix1 int
	fmt.Print("\nChoix : ")
	fmt.Scan(&choix1)

	switch choix1 {
	case 1 :
		fmt.Println("Vous avez appris à lancer une grenade !")
	} 
}


func (c *Character)CheckSpell(name string) bool{
	for _, skill := range c.skill {
		if skill == name {
			return true
		}
	}
	return false 
}


func (c *Character)LearnSpell(name string){
	if c.CheckSpell(name) == true {
		c.skill = append(c.skill, "Grenade")
	}
	
}