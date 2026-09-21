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

type Monster struct{
	Nom		string
	Max_vie	int
	Vie		int
	Points_attaque	int
}


func InitCharacter(nom, classe string, niveau, maxVie, vie int, inventaire []string,Equipement Equipment) Etudiant {
	return Etudiant{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		MaxVie:     maxVie,
		Vie:        vie,
		Inventaire: inventaire,
		Argent:     100,
		Equipement: Equipment{Headgear: "chapeau très impressionant",BodyArmor: "diamond body",FeetArmor: "chaussures qui court vite"},
	}
}

func InitGoblin() Monster{
	return Monster{
		Nom: 	"Goblin",
		Max_vie: 40,
		Vie: 40,
		Points_attaque: 5,
	}

}
	
func DisplayInfo(c Etudiant) {
	fmt.Print(c)
}
