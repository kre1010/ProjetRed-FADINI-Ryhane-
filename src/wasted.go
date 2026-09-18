package red

import "fmt"

type character struct {
	Name   string
	Health int
	Alive  bool
}

func player() {
	player := character{}
	damage(&player, 30)
	damage(&player, 50)
	damage(&player, 25)
}

func damage(p *character, amount int) {
	if !p.Alive {
		return
	}
	p.Health -= amount
	if p.Health <= 0 {
		p.Health = 0
		p.Alive = false
		fmt.Println(p.Name, "You Died")
	} else {
		fmt.Println(p.Name, "a", p.Health, "HP")
	}
}
