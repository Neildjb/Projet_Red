package forgeron

import (
	"Projet_Red/personnage"
	"Projet_Red/tache7101111suite"
	"fmt"
	"strconv"
)

const coutFabrication = 20

type Recette struct {
	Nom        string
	Ressources map[string]int // nom de la ressource -> quantité
}

var recettes = []Recette{
	{Nom: "bandeau frontale ninja", Ressources: map[string]int{"Baton de bois": 1}},
	{Nom: "Manteau Akatsuki", Ressources: map[string]int{"fer": 2}},
	{Nom: "Bottes de Shinobi", Ressources: map[string]int{"queue de demon": 3}},
}

type piece struct {
	Slot  string
	Bonus int
}

var pieces = map[string]piece{
	"bandeau frontale ninja": {Slot: "tete", Bonus: 15},
	"Manteau Akatsuki":       {Slot: "torse", Bonus: 20},
	"Bottes de Shinobi":      {Slot: "pieds", Bonus: 25},
}

func Forgeron(c *personnage.Etudiant) {

	for {
		fmt.Println("\nIl te reste", c.Argent, "pièces d'or. Chaque fabrication coûte", coutFabrication, "pièces.")
		for i, r := range recettes {
			fmt.Printf("%d - %s il te faudra : ", i+1, r.Nom)
			for ressource, quantite := range r.Ressources {
				fmt.Printf("%d %s ", quantite, ressource)
			}
			fmt.Println()
		}
		fmt.Println("0 - Retour")

		var saisie string
		fmt.Scanln(&saisie)

		n, err := strconv.Atoi(saisie)
		if err != nil || n < 0 || n > len(recettes) {
			fmt.Println("Ce n'est pas au menu, gamin.")
			continue
		}
		if n == 0 {
			return
		}
		r := recettes[n-1]

		// 1. assez d'argent ?
		if c.Argent < coutFabrication {
			fmt.Println("Sale pauvre t'as même pas assez, il me faudrait", coutFabrication, "pièces d'or.")
			continue
		}

		// 2. assez de ressources ?
		manque := false
		for nom, qte := range r.Ressources {
			if possede := compter(c, nom); possede < qte {
				fmt.Printf("Il te manque %d x %s\n", qte-possede, nom)
				manque = true
			}
		}
		if manque {
			continue
		}

		// 3. tout est bon : on applique
		c.Argent -= coutFabrication
		for nom, qte := range r.Ressources {
			for k := 0; k < qte; k++ {
				retirerUn(c, nom)
			}
		}
		tache7101111suite.AddInventory(c, r.Nom)
		fmt.Println(r.Nom, "fabriqué !")
	}
}

// Equiper équipe un objet de l'inventaire
func Equiper(c *personnage.Etudiant, nom string) {
	p, ok := pieces[nom]
	if !ok {
		fmt.Println("Cet objet ne s'équipe pas.")
		return
	}
	if compter(c, nom) == 0 {
		fmt.Println("Tu n'as pas", nom, "dans ton inventaire.")
		return
	}

	var emplacement *string
	switch p.Slot {
	case "tete":
		emplacement = &c.Equipement.Headgear
	case "torse":
		emplacement = &c.Equipement.BodyArmor
	case "pieds":
		emplacement = &c.Equipement.FeetArmor
	}
	if emplacement == nil {
		fmt.Println("Emplacement d'équipement invalide.")
		return
	}

	retirerUn(c, nom)
	ancien := *emplacement
	if ancien != "" && ancien != "rien" {
		c.MaxVie -= pieces[ancien].Bonus
		tache7101111suite.AddInventory(c, ancien)
	}

	*emplacement = nom
	c.MaxVie += p.Bonus
	if c.Vie > c.MaxVie {
		c.Vie = c.MaxVie
	}
	fmt.Println(nom, "équipé ! Vie maximum :", c.MaxVie)
}

func retirerUn(c *personnage.Etudiant, nom string) {
	for i, obj := range c.Inventaire {
		if obj == nom {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			return
		}
	}
}

func compter(c *personnage.Etudiant, nom string) int {
	total := 0
	for _, obj := range c.Inventaire {
		if obj == nom {
			total++
		}
	}
	return total
}
