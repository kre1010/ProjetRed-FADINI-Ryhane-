package character

import (
	"fmt"
)

func gagnerArgent(j *Character, montant int) {
	j.Argent += montant
	fmt.Printf("%s a gagné %d eddie !\n", j.Name, montant)
}

func afficherArgent(j *Character) {
	fmt.Printf("eddie actuel : %d eddie\n", j.Argent)
}

func (c *Character) AcheterObjet(nomObjet string, prix int) {
	if c.Argent >= prix {
		c.Argent -= prix
		fmt.Printf("Vous avez acheté %s pour %d eddie.\n", nomObjet, prix)
	} else {
		fmt.Printf("Pas assez d'eddie pour acheter %s.\n", nomObjet)
	}
}
