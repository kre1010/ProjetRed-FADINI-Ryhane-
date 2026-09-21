package character

import "fmt"

type character struct {
      Name string
      Pv int
      Alive bool
}

func Damage(c *character, amount int) {

if !c.Alive {
    fmt.Println(c.Name, "is already dead.")
              return
        }
             c.Pv -= amount
             if c.Pv <= 0 {
                c.Pv = 0
                c.Alive = false

            fmt.Println(c.Name, "you died !")
            } else {
             fmt.Printf("%s has %d HP remaining.\n", c.Name, c.Pv)
        }
    }

    func main() {
	    player := character{
            Name: "c.class",
               Pv: 100,
               Alive: true,
		}
      		Damage(&player, 30)
            Damage(&player, 50)
            Damage(&player, 25)
            Damage(&player, 10)
}