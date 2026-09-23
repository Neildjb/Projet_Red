package npc

import (
	"Projet_Red/personnage"
	"fmt"
	"strconv"
)

const coutFabrication = 20

type Recette struct {
	Nom        string
	Ressources map[string]int
}

var recettes = []Recette{
	{Nom: "bandeau frontale ninja", Ressources: map[string]int{"Baton de bois": 1}},
	{Nom: "Manteau Akatsuki", Ressources: map[string]int{"fer": 2}},
	{Nom: "Pantalon des Six Chemins", Ressources: map[string]int{"chakra": 3}},
	{Nom: "Sandales du Shinobi", Ressources: map[string]int{"queue de demon": 4}},
}

type piece struct {
	Slot  string
	Bonus int
}

var pieces = map[string]piece{
	"bandeau frontale ninja":   {Slot: "tete", Bonus: 15},
	"Manteau Akatsuki":         {Slot: "torse", Bonus: 20},
	"Pantalon des Six Chemins": {Slot: "jambes", Bonus: 25},
	"Sandales du Shinobi":      {Slot: "pieds", Bonus: 30},
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

		if c.Argent < coutFabrication {
			fmt.Println("Sale pauvre t'as même pas assez, il me faudrait", coutFabrication, "pièces d'or.")
			continue
		}

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

		c.Argent -= coutFabrication
		for nom, qte := range r.Ressources {
			for k := 0; k < qte; k++ {
				retirerUn(c, nom)
			}
		}
		AddInventory(c, r.Nom)
		fmt.Println(r.Nom, "fabriqué !")
	}
}

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
	case "jambes":
		emplacement = &c.Equipement.LegArmor
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
		c.Vie -= pieces[ancien].Bonus
		AddInventory(c, ancien)
	}

	*emplacement = nom
	c.MaxVie += p.Bonus
	c.Vie += p.Bonus
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
