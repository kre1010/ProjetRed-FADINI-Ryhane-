package equipement

import "fmt"

type Equipement struct {
	Tête  map[string]int
	Torse map[string]int
	Jambe map[string]int
	Mains map[string]int
	Pieds map[string]int
}

var Equip = Equipement{
	Tête:  make(map[string]int),
	Torse: make(map[string]int),
	Jambe: make(map[string]int),
	Mains: make(map[string]int),
	Pieds: make(map[string]int),
}

func AfficherEquipement() {
	fmt.Println("========== ÉQUIPEMENT ==========")

	fmt.Println("Tête :")
	for nom, quantite := range Equip.Tête {
		fmt.Println("-", nom, "x", quantite)
	}

	fmt.Println("Torse :")
	for nom, quantite := range Equip.Torse {
		fmt.Println("-", nom, "x", quantite)
	}

	fmt.Println("Jambe :")
	for nom, quantite := range Equip.Jambe {
		fmt.Println("-", nom, "x", quantite)
	}

	fmt.Println("Mains :")
	for nom, quantite := range Equip.Mains {
		fmt.Println("-", nom, "x", quantite)
	}

	fmt.Println("Pieds :")
	for nom, quantite := range Equip.Pieds {
		fmt.Println("-", nom, "x", quantite)
	}
}
