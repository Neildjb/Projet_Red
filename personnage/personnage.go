package personnage

import (
	"fmt"
	"strings"
)

type Etudiant struct {
	Nom                  string
	Classe               string
	ExperienceCombat     int
	MaxVie               int
	Vie                  int
	Inventaire           []string
	CapaciteInventaire   int
	Argent               int
	Equipement           Equipment
	Skills               []string
	AchatsMarchand       map[string]bool
	EquipementsFabriques map[string]bool
	Poison               bool
	Hardcore             bool
	GameOver             bool
	TutorielFait         bool
}

type Equipment struct {
	Headgear  string
	BodyArmor string
	LegArmor  string
	FeetArmor string
}

func InitCharacter(nom, classe string, experienceCombat, maxVie, vie, argent int, inventaire []string, equipement Equipment, skills []string) Etudiant {
	bonusVie := bonusEquipement(equipement.Headgear) + bonusEquipement(equipement.BodyArmor) + bonusEquipement(equipement.LegArmor) + bonusEquipement(equipement.FeetArmor)
	maxVie += bonusVie
	vie += bonusVie

	return Etudiant{
		Nom:                  nom,
		Classe:               classe,
		ExperienceCombat:     experienceCombat,
		MaxVie:               maxVie,
		Vie:                  vie,
		Inventaire:           inventaire,
		CapaciteInventaire:   10,
		Argent:               argent,
		Equipement:           equipement,
		Skills:               skills,
		AchatsMarchand:       achatsMarchandDepuisInventaire(inventaire),
		EquipementsFabriques: equipementsFabriquesDepuisEtat(inventaire, equipement),
	}
}

func achatsMarchandDepuisInventaire(inventaire []string) map[string]bool {
	achats := make(map[string]bool)
	for _, item := range inventaire {
		switch item {
		case "Lot de Kunaï (nouveau sort)", "Rasengan (nouveau sort)", "upgradeinventoryslot1", "upgradeinventoryslot2":
			achats[item] = true
		}
	}
	return achats
}

func equipementsFabriquesDepuisEtat(inventaire []string, equipement Equipment) map[string]bool {
	fabriques := make(map[string]bool)
	for _, item := range inventaire {
		if estEquipement(item) {
			fabriques[item] = true
		}
	}
	for _, item := range []string{equipement.Headgear, equipement.BodyArmor, equipement.LegArmor, equipement.FeetArmor} {
		if estEquipement(item) {
			fabriques[item] = true
		}
	}
	return fabriques
}

func estEquipement(item string) bool {
	switch item {
	case "bandeau frontale ninja", "Manteau Akatsuki", "Pantalon des Six Chemins", "Sandales du Shinobi":
		return true
	default:
		return false
	}
}

func DisplayInfo(c Etudiant) {
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Expérience de combat :", c.ExperienceCombat)
	fmt.Println("Vie :", c.Vie, "/", c.MaxVie)
	fmt.Println("Inventaire :", strings.Join(c.Inventaire, ", "))
	fmt.Println("Capacité inventaire :", len(c.Inventaire), "/", c.CapaciteInventaire)
	fmt.Println("Argent :", c.Argent)
	fmt.Println("Équipement :")
	fmt.Println("  Tête :", afficherEquipement(c.Equipement.Headgear))
	fmt.Println("  Corps :", afficherEquipement(c.Equipement.BodyArmor))
	fmt.Println("  Jambes :", afficherEquipement(c.Equipement.LegArmor))
	fmt.Println("  Pieds :", afficherEquipement(c.Equipement.FeetArmor))
	fmt.Println("Compétences :", strings.Join(c.Skills, ", "))
}

func afficherEquipement(nom string) string {
	bonus := bonusEquipement(nom)
	if bonus == 0 || nom == "" || nom == "rien" {
		return nom
	}
	return fmt.Sprintf("%s (+%d PV)", nom, bonus)
}

func bonusEquipement(nom string) int {
	return map[string]int{
		"bandeau frontale ninja":   15,
		"Manteau Akatsuki":         20,
		"Pantalon des Six Chemins": 25,
		"Sandales du Shinobi":      30,
	}[nom]
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
	Drop           []string
	ArgentDrop     int
	ExperienceDrop int
	Poison         bool
}

func NomMonstrePourClasse(nomMonstre, classe string) string {
	noms := map[string]map[string]string{
		"Naruto_Prime": {
			"Ninjas déserteurs": "Neil",
			"Golems de chakra":  "Andy",
			"Demon a queue":     "Mathias",
			"Madara":            "Maxime_boss",
		},
		"Admin4416": {
			"Ninjas déserteurs": "Neil",
			"Golems de chakra":  "Andy",
			"Demon a queue":     "Mathias",
			"Madara":            "Maxime_boss",
		},
		"Kage": {
			"Ninjas déserteurs": "Pain",
			"Golems de chakra":  "Itachi",
			"Demon a queue":     "Obito",
			"Madara":            "Madara",
		},
		"Jonin": {
			"Ninjas déserteurs": "Deidara",
			"Golems de chakra":  "Kisame",
			"Demon a queue":     "Sasori",
			"Madara":            "Orochimaru",
		},
		"Genin": {
			"Ninjas déserteurs": "Asuma",
			"Golems de chakra":  "Gaara",
			"Demon a queue":     "Shino",
			"Madara":            "Kakshi",
		},
		"Ninja": {
			"Ninjas déserteurs": "Ninja déserteur",
			"Golems de chakra":  "Sai",
			"Demon a queue":     "Yamato",
			"Madara":            "Danzo",
		},
	}

	if nomsParClasse, ok := noms[classe]; ok {
		if nom, ok := nomsParClasse[nomMonstre]; ok {
			return nom
		}
	}
	return nomMonstre
}

func Init_Maxime() Monster {
	return Monster{
		Nom:            "Maxime",
		Max_vie:        35,
		Vie:            35,
		Points_attaque: 10,
		Drop:           []string{"Baton de bois"},
		ArgentDrop:     10,
		ExperienceDrop: 1,
	}
}

func Init_Ninjas_déserteurs() Monster {
	return Monster{
		Nom:            "Ninjas déserteurs",
		Max_vie:        60,
		Vie:            60,
		Points_attaque: 15,
		Drop:           []string{"fer"},
		ArgentDrop:     20,
		ExperienceDrop: 1,
	}
}

func Init_Golems_de_chakra() Monster {
	return Monster{
		Nom:            "Golems de chakra",
		Max_vie:        85,
		Vie:            85,
		Points_attaque: 20,
		Drop:           []string{"chakra"},
		ArgentDrop:     45,
		ExperienceDrop: 1,
	}
}

func Init_Demon_a_queue() Monster {
	return Monster{
		Nom:            "Demon a queue",
		Max_vie:        120,
		Vie:            120,
		Points_attaque: 25,
		Drop:           []string{"queue de demon"},
		ArgentDrop:     65,
		ExperienceDrop: 2,
	}
}

func Init_Madara() Monster {
	return Monster{
		Nom:            "Madara",
		Max_vie:        200,
		Vie:            200,
		Points_attaque: 35,
		Drop:           []string{"Sharingan", "Trophee en or"},
		ArgentDrop:     150,
		ExperienceDrop: 5,
	}
}
