package character

import (
	"fmt"
)

func gagnerArgent(c *Character, montant int) {
	c.Argent += montant
	fmt.Printf("%s a gagné %d eddie !\n", c.Name, montant)
}

func afficherArgent(c *Character) {
	fmt.Printf("eddie actuel : %d eddie\n", c.Argent)
}

func (c *Character) AcheterObjet(nomObjet string, prix int) bool {
	if c.Argent >= prix {
		c.Argent -= prix
		fmt.Printf("Vous avez acheté %s pour %d eddies.\n", nomObjet, prix)
		return true
	}

	fmt.Printf("Pas assez d'eddies pour acheter %s.\n", nomObjet)
	return false
}


