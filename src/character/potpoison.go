package character

import (
	"fmt"
	"time"
)

func CocktailMolotov(c *Character, ennemi *Operateur) {

	fmt.Println("Vous lancez un Cocktail Molotov !")

	for i := 0; i < 3; i++ {

		ennemi.Hpactuel -= 10
		if ennemi.Hpactuel < 0 {
			ennemi.Hpactuel = 0
		}
		
		c.Hp -= 10

		if c.Hp < 0 {
			c.Hp = 0
		}

		fmt.Println("-10 HP à", ennemi.Nom)
		fmt.Println("-10 HP à", c.Name)

		time.Sleep(1 * time.Second)
	}

	fmt.Println("Total : -30 HP à l'ennemi et -30 HP au joueur")
}
