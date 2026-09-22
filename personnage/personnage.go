package personnage

import "fmt"

type Etudiant struct {
	Nom                string
	Classe             string
	Niveau             int
	MaxVie             int
	Vie                int
	Inventaire         []string
	CapaciteInventaire int
	Argent             int
	Equipement         Equipment
	Skills             []string
}

type Equipment struct {
	Headgear  string
	BodyArmor string
	FeetArmor string
}

func InitCharacter(nom, classe string, niveau, maxVie, vie, argent int, inventaire []string, equipement Equipment, skills []string) Etudiant {
	return Etudiant{
		Nom:                nom,
		Classe:             classe,
		Niveau:             niveau,
		MaxVie:             maxVie,
		Vie:                vie,
		Inventaire:         inventaire,
		CapaciteInventaire: 10,
		Argent:             argent,
		Equipement:         equipement,
		Skills:             skills,
	}
}

func DisplayInfo(c Etudiant) {
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Println("Vie :", c.Vie, "/", c.MaxVie)
	fmt.Println("Inventaire :", c.Inventaire)
	fmt.Println("Capacité inventaire :", len(c.Inventaire), "/", c.CapaciteInventaire)
	fmt.Println("Argent :", c.Argent)
	fmt.Println("Équipement :")
	fmt.Println("  Tête :", c.Equipement.Headgear)
	fmt.Println("  Corps :", c.Equipement.BodyArmor)
	fmt.Println("  Pieds :", c.Equipement.FeetArmor)
	fmt.Println("Compétences :", c.Skills)
}

func AjouterItem(c *Etudiant, item string) {
	if len(c.Inventaire) >= c.CapaciteInventaire {
		fmt.Println("Inventaire plein.")
		return
	}
	c.Inventaire = append(c.Inventaire, item)
}

type Monster struct {
	Nom            string
	Max_vie        int
	Vie            int
	Points_attaque int
	Drop           string
}

func Init_Maxime() Monster {
	return Monster{
		Nom:            "Maxime",
		Max_vie:        40,
		Vie:            40,
		Points_attaque: 10,
		Drop:           "Baton de bois",
	}
}

func Init_Ninjas_déserteurs() Monster {
	return Monster{
		Nom:            "Ninjas déserteurs",
		Max_vie:        60,
		Vie:            60,
		Points_attaque: 15,
		Drop:           "fer",
	}
}

func Init_Golems_de_chakra() Monster {
	return Monster{
		Nom:            "Golems de chakra",
		Max_vie:        80,
		Vie:            80,
		Points_attaque: 15,
		Drop:           "chakra",
	}
}

func Init_Demon_a_queue() Monster {
	return Monster{
		Nom:            "Demon a queue",
		Max_vie:        100,
		Vie:            100,
		Points_attaque: 25,
		Drop:           "queue de demon",
	}
}

func Init_Madara() Monster {
	return Monster{
		Nom:            "Madara",
		Max_vie:        150,
		Vie:            150,
		Points_attaque: 35,
		Drop:           "armure detruite",
	}
}
