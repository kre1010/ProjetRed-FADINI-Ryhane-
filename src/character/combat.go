package character

import (
	"fmt"
	"math/rand"
)

func Combat(c *Character, ennemi *Operateur) {

	for c.Hp > 0 && ennemi.Hpactuel > 0 {

		fmt.Println()
		fmt.Println("========== COMBAT ==========")
		fmt.Println(c.Name, ":", c.Hp, "/", c.MaxHp, "HP")
		fmt.Println(ennemi.Nom, ":", ennemi.Hpactuel, "/", ennemi.Hpmax, "HP")
		fmt.Println("============================")

		fmt.Println("1. Attaquer")
		fmt.Println("2. Défendre")
		fmt.Println("3. Inventaire")

		var choix int
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {

		// ATTAQUE
		case 1:

			de := rand.Intn(9) + 1
			degats := 0

			switch c.Class {

			// Sniper
			case 1:
				if de <= 4 {
					degats = 20
				} else if de <= 7 {
					degats = 30
				} else {
					degats = 45
				}

			// Assault
			case 2:
				if de <= 4 {
					degats = 15
				} else if de <= 7 {
					degats = 25
				} else {
					degats = 35
				}

			// Shotgunner
			case 3:
				if de <= 4 {
					degats = 10
				} else if de <= 7 {
					degats = 20
				} else {
					degats = 40
				}
			}

			fmt.Println("Dé :", de)

			if de >= 8 {
				fmt.Println("COUP CRITIQUE !")
			}

			fmt.Println("Vous infligez", degats, "dégâts.")

			ennemi.Hpactuel -= degats

			if ennemi.Hpactuel < 0 {
				ennemi.Hpactuel = 0
			}

			if ennemi.Hpactuel <= 0 {
				fmt.Println("Vous avez vaincu", ennemi.Nom, "!")
				c.Argent += 75

				fmt.Println("Vous gagnez 75 Eddies !")
				fmt.Println("Eddies :", c.Argent)
				return
			}

		// DÉFENSE
		case 2:

			c.Hp += ennemi.Attaque

			if c.Hp > c.MaxHp {
				c.Hp = c.MaxHp
			}

			fmt.Println("Vous vous défendez.")
			fmt.Println("HP :", c.Hp, "/", c.MaxHp)

		// INVENTAIRE
		case 3:

			fmt.Println()
			fmt.Println("1. Potion de Soin")
			fmt.Println("2. MedKit")
			fmt.Println("3. Cocktail Molotov")
			fmt.Println("4. Retour")

			var choixInv int
			fmt.Print("Choix : ")
			fmt.Scan(&choixInv)

			switch choixInv {

			case 1:
				if c.Inv["Potion de Soin"] <= 0 {
					fmt.Println("Vous n'avez pas de Potion de Soin.")
					continue
				}

				c.Hp += 50

				if c.Hp > c.MaxHp {
					c.Hp = c.MaxHp
				}

				c.Inv["Potion de Soin"]--

				fmt.Println("Potion utilisée.")
				fmt.Println("HP :", c.Hp, "/", c.MaxHp)

			case 2:
				if c.Inv["MedKit"] <= 0 {
					fmt.Println("Vous n'avez pas de MedKit.")
					continue
				}

				c.Hp += 200

				if c.Hp > c.MaxHp {
					c.Hp = c.MaxHp
				}

				c.Inv["MedKit"]--

				fmt.Println("MedKit utilisé.")
				fmt.Println("HP :", c.Hp, "/", c.MaxHp)

			case 3:
				if c.Inv["Cocktail Molotov"] <= 0 {
					fmt.Println("Vous n'avez pas de Cocktail Molotov.")
					continue
				}

				c.Inv["Cocktail Molotov"]--

				CocktailMolotov(c, ennemi)

				if ennemi.Hpactuel < 0 {
					ennemi.Hpactuel = 0
				}

				if c.Hp < 0 {
					c.Hp = 0
				}

				fmt.Println("Cocktail Molotov lancé !")
				fmt.Println(ennemi.Nom, "perd 30 HP.")
				fmt.Println(c.Name, "perd aussi 30 HP.")

				if ennemi.Hpactuel <= 0 {
					c.Level += 1
					fmt.Println("Level:", c.Level)
					fmt.Println("Vous avez vaincu", ennemi.Nom, "!")

					c.Argent += 75
					fmt.Println("Vous gagnez 75 Eddies !")
					fmt.Println("Eddies :", c.Argent)
					return
				}

				if c.Hp <= 0 {
					return
				}

			case 4:
				continue

			default:
				fmt.Println("Choix invalide.")
				continue
			}

		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if ennemi.Hpactuel > 0 {

			fmt.Println()
			fmt.Println(ennemi.Nom, "attaque !")

			Damage(c, ennemi.Attaque)

			if c.Hp <= 0 {
				return
			}
		}
	}
}
