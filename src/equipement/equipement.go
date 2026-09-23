package equipement

import "red/src/character"

type Equipement struct {
	Tête  map[string]int
	Torse map[string]int
	Pieds map[string]int
}

var Equip = Equipement{
	Tête : make(map[string]int),
	Torse : make(map[string]int),
	Pieds : make(map[string]int),
}
func Equitation() {
	if character.Inv["Kevlar"] > 0 {
		Equip.
	}
}
