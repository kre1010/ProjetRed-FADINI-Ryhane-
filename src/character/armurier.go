package character

import "fmt"

func (c *Character) OpenArmurier() {
	fmt.Println("\033[1;33m=== ARMURIER ===\033[0m")
	fmt.Println("\033[36m1.\033[0m gilet pare-balles militech - 30 eddie")
	fmt.Println("\033[36m2.\033[0m casque pare balles ops-core - 10 eddie")
	fmt.Println("\033[36m3.\033[0m pantalon renforcé crye precision g4 - 10 eddie")
	fmt.Println("\033[36m4.\033[0m gants Oakley SI - 5 eddie")
	fmt.Println("\033[36m5.\033[0m bottes Salomon Forces - 5 eddie")
	fmt.Println("\033[36m6.\033[0m plaque en kevlar - 2 eddie")
	fmt.Println("\033[36m8.\033[0m Augmentation d'inventaire - 30 eddie")
	fmt.Println("\033[36m0.\033[0m Quitter l'armurier")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.AcheterObjet("gilet pare-balles militech", 30)
	case 2:
		c.AcheterObjet("casque pare balles ops-core", 10)
	case 3:
		if c.Argent < 10 {
			fmt.Println("\033[31mPas assez d'eddie pour acheter pantalon renforcé crye precision g4\033[0m")
			return
		}
		c.Argent -= 10
		c.OpenArmurier()
	case 4:
		c.AcheterObjet("pantalon renforcé crye precision g4", 10)
	case 5:
		c.AcheterObjet("gants Oakley SI", 5)
	case 6:
		c.AcheterObjet("bottes Salomon Forces", 5)
	case 7:
		c.AcheterObjet("plaque en kevlar", 2)
	case 8:
		if c.Argent < 30 {
			fmt.Println("\033[31mPas assez d'eddie pour cette amélioration\033[0m")
			return
		}
		c.Argent -= 30
		c.AcheterObjet("Augmentation d'inventaire", 30)
		fmt.Println("\033[32mAugmentation d'inventaire achetée !\033[0m")
	case 0:
		return
	default:
		fmt.Println("\033[31mChoix invalide\033[0m")
	}
}
