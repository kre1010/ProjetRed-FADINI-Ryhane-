package character

import "fmt"

func Damage(c *Character, wewewe int) {

	if c.Hp <= 0 {
		fmt.Println(c.Name, "tu es déjà mort")
		return
	}
	c.Hp -= wewewe
	if c.Hp <= 0 {
		c.Hp = 0

		fmt.Println(c.Name, "T'es mort!")
	} else {
		fmt.Printf("%s as %d HP restant.\n", c.Name, c.Hp)
	}
}
