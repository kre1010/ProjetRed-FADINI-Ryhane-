package character

import "fmt"

func (c *Character) OpenArmurier() {
	fmt.Println("\033[1;33m=== ARMURIER ===\033[0m")
	fmt.Println("\033[36m1.\033[0m banadage - 3 eddie")
	fmt.Println("\033[36m2.\033[0m Potion de poison - 6 eddie")
	fmt.Println("\033[36m3.\033[0m Livre de Sort : Boule de Feu - 25 eddie")
	fmt.Println("\033[36m4.\033[0m Fourrure de Loup - 4 eddie")
	fmt.Println("\033[36m5.\033[0m Peau de Troll - 7 eddie")
	fmt.Println("\033[36m6.\033[0m Cuir de Sanglier - 3 eddie")
	fmt.Println("\033[36m7.\033[0m Plume de Corbeau - 1 eddie")
	fmt.Println("\033[36m8.\033[0m Augmentation d'inventaire - 30 eddie")
	fmt.Println("\033[36m0.\033[0m Quitter l'armurier")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.buyItem("Potion de vie", 3)
	case 2:
		c.buyItem("Potion de poison", 6)
	case 3:
		if c.Gold < 25 {
			fmt.Println("\033[31mPas assez d'eddie pour acheter le Livre de Sort\033[0m")
			return
		}
		c.Gold -= 25
		c.SpellBook()
	case 4:
		c.buyItem("Fourrure de Loup", 4)
	case 5:
		c.buyItem("Peau de Troll", 7)
	case 6:
		c.buyItem("Cuir de Sanglier", 3)
	case 7:
		c.buyItem("Plume de Corbeau", 1)
	case 8:
		if c.Gold < 30 {
			fmt.Println("\033[31mPas assez d'eddie pour cette amélioration\033[0m")
			return
		}
		c.Gold -= 30
		c.UpgradeInventorySlot()
	case 0:
		return
	default:
		fmt.Println("\033[31mChoix invalide\033[0m")
	}
}
