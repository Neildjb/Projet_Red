package personnage

import "fmt"

type Etudiant struct {
	Nom        string
	Classe     string
	Niveau     int
	MaxVie     int
	Vie        int
	Inventaire []string
}

func InitCharacter(nom, classe string, niveau, maxVie, vie int, inventaire []string) Etudiant {
	return Etudiant{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		MaxVie:     maxVie,
		Vie:        vie,
		Inventaire: inventaire,
	}
}
func DisplayInfo(c Etudiant) {
	fmt.Print(c)
}


