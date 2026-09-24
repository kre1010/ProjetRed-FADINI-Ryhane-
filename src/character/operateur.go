package character

type Operateur struct {
	Nom      string
	Hpmax    int
	Hpactuel int
	Attaque  int
}

var Operateurs = map[string]Operateur{
	"Neavy Seal":    {Nom: "Neavy Seal", Hpmax: 130, Hpactuel: 130, Attaque: 40},
	"Spetnaz":       {Nom: "Spetnaz", Hpmax: 120, Hpactuel: 120, Attaque: 35},
	"Bope":          {Nom: "Bope", Hpmax: 120, Hpactuel: 120, Attaque: 30},
	"KSK":           {Nom: "KSK", Hpmax: 110, Hpactuel: 110, Attaque: 25},
	"Jaegerkorpset": {Nom: "Jaegerkorpset", Hpmax: 110, Hpactuel: 110, Attaque: 20},
	"Agent d'entrainement": {Nom: "Agent d'entrainement", Hpmax: 40, Hpactuel: 40, Attaque: 5},
}
