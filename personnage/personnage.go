package personnage

import "fmt"

type Etudiant struct {
	Nom        string
	Classe     string
	Niveau     int
	MaxVie     int
	Vie        int
	Inventaire []string
	Argent     int
	Equipement Equipment
}

type Equipment struct {
	Headgear  string
	BodyArmor string
	FeetArmor string
}

type Monster struct {
	Nom            string
	Max_vie        int
	Vie            int
	Points_attaque int
}

func InitCharacter(nom, classe string, niveau, maxVie, vie int, inventaire []string, argent int, equipement Equipment) Etudiant {
	return Etudiant{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		MaxVie:     maxVie,
		Vie:        vie,
		Inventaire: inventaire,
		Argent:     argent,
		Equipement: equipement,
	}
}

func InitGoblin() Monster {
	return Monster{
		Nom:            "Goblin",
		Max_vie:        40,
		Vie:            40,
		Points_attaque: 5,
	}

}

func DisplayInfo(c Etudiant) {
	info := fmt.Sprintf(
		"Nom: %s\nClasse: %s\nNiveau: %d\nMax vie: %d\nVie: %d\nArgent: %d\nInventaire: %v\nEquipement:\n  Tete: %s\n  Armure: %s\n  Pieds: %s\n",
		c.Nom,
		c.Classe,
		c.Niveau,
		c.MaxVie,
		c.Vie,
		c.Argent,
		c.Inventaire,
		c.Equipement.Headgear,
		c.Equipement.BodyArmor,
		c.Equipement.FeetArmor,
	)
	fmt.Print(info)
}

func AjouterItem(c *Etudiant, item string) {
	c.Inventaire = append(c.Inventaire, item)
}
