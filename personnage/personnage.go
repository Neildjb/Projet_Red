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
	Skills     []string
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

func InitCharacter(nom, classe string, niveau, maxVie, vie, argent int, inventaire []string, equipement Equipment, skills []string) Etudiant {
	return Etudiant{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		MaxVie:     maxVie,
		Vie:        vie,
		Inventaire: inventaire,
		Argent:     argent,
		Equipement: equipement,
		Skills:     skills,
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
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Println("Vie :", c.Vie, "/", c.MaxVie)
	fmt.Println("Inventaire :", c.Inventaire)
	fmt.Println("Argent :", c.Argent)
	fmt.Println("Équipement :")
	fmt.Println("  Tête :", c.Equipement.Headgear)
	fmt.Println("  Corps :", c.Equipement.BodyArmor)
	fmt.Println("  Pieds :", c.Equipement.FeetArmor)
	fmt.Println("Compétences :", c.Skills)
}

func AjouterItem(c *Etudiant, item string) {
	c.Inventaire = append(c.Inventaire, item)
}
