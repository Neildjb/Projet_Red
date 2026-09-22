package tache7101111suite

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
		prix := prixActuel(item)
		if c.Argent < prix {
			fmt.Println("T'as même pas assez pour " + item.Nom + ", c'est ridicule")
			continue
		}

		c.Argent -= prix
		AddInventory(c, item.Nom)
		if item.Nom == "Potion de soin" {
			potionGratuiteUtilisee = true
		}
		fmt.Println("C'est une bonne affaire d'acheter " + item.Nom)
	}
}

func AddInventory(c *personnage.Etudiant, item string) {
	c.Inventaire = append(c.Inventaire, item)
}

func RemoveInventory(c *personnage.Etudiant, item string) {
	var res []string
	for _, v := range c.Inventaire {
		if v != item {
			res = append(res, v)
		}
	}
	c.Inventaire = res
}

type Item struct {
	Nom  string
	Prix int
}

var boutique = []Item{
	{Nom: "Kunaï", Prix: 25},
	{Nom: "Rasengan", Prix: 50},
	{Nom: "Potion de soin", Prix: 20},
}

var potionGratuiteUtilisee = false

// prixActuel donne le prix réel d'un objet à cet instant
func prixActuel(it Item) int {
	if it.Nom == "Potion de soin" && !potionGratuiteUtilisee {
		return 0
	}
	return it.Prix
}
