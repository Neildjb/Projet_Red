package npc

// TACHE 7 : Marchand

import (
	"Projet_Red/personnage"
	"fmt"
	"strconv"
)

func Marchand(c *personnage.Etudiant) {
	fmt.Println("Je suis le marchand, je peux vous vendre des objets pour améliorer vos compétences.")

	for {
		fmt.Println("Tu as l'embarras du choix. Il te reste", c.Argent, "euros.")
		for i, it := range boutique {
			fmt.Printf("%d - %s (%d euros)\n", i+1, it.Nom, prixActuel(it))
		}
		fmt.Println("0 - Retour")

		var saisie string
		fmt.Scanln(&saisie)

		n, err := strconv.Atoi(saisie)
		if err != nil || n < 0 || n > len(boutique) {
			fmt.Println("Gamin t'es taré ou quoi. Ce dont tu parles n'existe pas")
			continue
		}
		if n == 0 {
			return
		}

		item := boutique[n-1]
		if articleDejaAchete(c, item.Nom) {
			fmt.Println("Cet article ne peut être acheté qu'une seule fois.")
			continue
		}
		prix := prixActuel(item)
		if c.Argent < prix {
			fmt.Println("T'as même pas assez pour " + item.Nom + ", c'est ridicule")
			continue
		}

		if ajouterAchatMarchand(c, item.Nom) {
			c.Argent -= prix
			switch item.Nom {
			case "Potion de soin":
				nombrePotionsSoin++
			case "Potion de poison":
				nombrePotionsPoison++
			}
			fmt.Println("C'est une bonne affaire d'acheter " + item.Nom)
			if item.Nom == "Lot de Kunaï (nouveau sort)" || item.Nom == "Rasengan (nouveau sort)" {
				fmt.Println("Allez dans l'inventaire et utilisez l'objet pour apprendre une nouvelle technique.")
			}
		} else {
			fmt.Println("Il y'en a déja un autre ne sois pas gourmand")
		}

	}
}

func AddInventory(c *personnage.Etudiant, item string) bool {
	doublon := false
	if articleVenduUneSeuleFois(item) {

		for _, v := range c.Inventaire {
			if v == item {
				doublon = true
				fmt.Println("Il y'en a déja un autre ne sois pas gourmand")
			}
			if doublon {
				return false
			}
		}
	}

	c.Inventaire = append(c.Inventaire, item)
	return true
}

func articleVenduUneSeuleFois(item string) bool {
	switch item {
	case "Lot de Kunaï (nouveau sort)", "Rasengan (nouveau sort)", "upgradeinventoryslot1", "upgradeinventoryslot2":
		return true
	default:
		return false
	}
}

func articleDejaAchete(c *personnage.Etudiant, item string) bool {
	if !articleVenduUneSeuleFois(item) {
		return false
	}
	return c.AchatsMarchand != nil && c.AchatsMarchand[item]
}

func enregistrerAchatMarchand(c *personnage.Etudiant, item string) {
	if c.AchatsMarchand == nil {
		c.AchatsMarchand = make(map[string]bool)
	}
	c.AchatsMarchand[item] = true
}

func ajouterAchatMarchand(c *personnage.Etudiant, item string) bool {
	if articleDejaAchete(c, item) {
		return false
	}
	if !AddInventory(c, item) {
		return false
	}
	enregistrerAchatMarchand(c, item)
	return true
}

func RemoveInventory(c *personnage.Etudiant, item string) {
	for i, v := range c.Inventaire {
		if v == item {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			return
		}
	}
}

type Item struct {
	Nom  string
	Prix int
}

var boutique = []Item{
	{Nom: "Lot de Kunaï (nouveau sort)", Prix: 30},
	{Nom: "Rasengan (nouveau sort)", Prix: 75},
	{Nom: "Potion de soin", Prix: 20},
	{Nom: "Potion de poison", Prix: 25},
	{Nom: "Potion de guérison du poison", Prix: 30},
	{Nom: "Potion de PV total", Prix: 60},
	{Nom: "upgradeinventoryslot1", Prix: 40},
	{Nom: "upgradeinventoryslot2", Prix: 80},
}

var nombrePotionsSoin int
var nombrePotionsPoison int

func prixActuel(it Item) int {
	switch it.Nom {
	case "Potion de soin":
		switch {
		case nombrePotionsSoin == 0:
			return 0
		case nombrePotionsSoin == 1:
			return 10
		case nombrePotionsSoin <= 3:
			return 15
		default:
			return 20
		}
	case "Potion de poison":
		if nombrePotionsPoison == 0 {
			return 15
		}
		return 25
	}
	return it.Prix
}
