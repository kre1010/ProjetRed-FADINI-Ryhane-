package character

import "fmt"

func (c *Character) OpenArmurier() {
	fmt.Println("\033[1;33m=== ARMURIER ===\033[0m")
	fmt.Println("\033[36m1.\033[0m bandage - 3 eddie")
	fmt.Println("\033[36m2.\033[0m medkit- 6 eddie")
	fmt.Println("\033[36m3.\033[0m cache flammes : Boule de Feu - 25 eddie")
	fmt.Println("\033[36m4.\033[0m viseur holographique - 4 eddie")
	fmt.Println("\033[36m5.\033[0m gachette rapide - 7 eddie")
	fmt.Println("\033[36m6.\033[0m habiliter au tir - 3 eddie")
	fmt.Println("\033[36m7.\033[0m munitions - 1 eddie")
	fmt.Println("\033[36m8.\033[0m Augmentation d'inventaire - 30 eddie")
	fmt.Println("\033[36m0.\033[0m Quitter l'armurier")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.buyItem("Bandage", 3)
	case 2:
		c.buyItem("redbulle", 6)
	case 3:
		if c.Argent < 25 {
			fmt.Println("\033[31mPas assez d'eddie pour acheter le Cache Flammes\033[0m")
			return
		}
		c.Argent -= 25
		c.OpenArmurier()
	case 4:
		c.buyItem("viseur holographique", 4)
	case 5:
		c.buyItem("gachette rapide", 7)
	case 6:
		c.buyItem("habiliter au tir", 3)
	case 7:
		c.buyItem("munitions", 1)
	case 8:
		if c.Argent < 30 {
			fmt.Println("\033[31mPas assez d'eddie pour cette amélioration\033[0m")
			return
		}
		c.Argent -= 30
		c.UpgradeInventorySlot()
	case 0:
		return
	default:
		fmt.Println("\033[31mChoix invalide\033[0m")
	}
}
