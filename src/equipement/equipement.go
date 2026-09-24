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
	fmt.Println(Equip.Tête)
	fmt.Println("Torse :")
	fmt.Println(Equip.Torse)
	fmt.Println("Jambe :")
	fmt.Println(Equip.Jambe)
	fmt.Println("Mains :")
	fmt.Println(Equip.Mains)
	fmt.Println("Pieds :")
	fmt.Println(Equip.Pieds)
	fmt.Println("=================================")
	fmt.Println()
	fmt.Println("Appuyer sur entrée pour revenir...")
}


