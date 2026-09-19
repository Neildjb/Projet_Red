package personnage

import "fmt"

type Etudiant struct {
	Nom        string
	Classe     string
	Niveau     int
	MaxVie     int
	Vie        int
	Inventaire []string
	Argent	   int
	Equipement Equipment
}

type Equipment struct{
	Headgear string
	BodyArmor string
	FeetArmor string
}

func InitCharacter(nom, classe string, niveau, maxVie, vie int, inventaire []string) Etudiant {
	return Etudiant{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		MaxVie:     maxVie,
		Vie:        vie,
		Inventaire: inventaire,
		Argent:     100,
	}
}
func DisplayInfo(c Etudiant) {
	fmt.Print(c)
}


