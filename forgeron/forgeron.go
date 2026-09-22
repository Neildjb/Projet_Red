package forgeron

import (
	"Projet_Red/personnage"
	"Projet_Red/tache7101111suite"
	"fmt"
	"strconv"
)

const coutFabrication = 5

type Recette struct {
	Nom        string
	Ressources map[string]int // nom de la ressource -> quantité
}

var recettes = []Recette{
	{Nom: "Chapeau de l'aventurier", Ressources: map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1}},
	{Nom: "Tunique de l'aventurier", Ressources: map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1}},
	{Nom: "Bottes de l'aventurier", Ressources: map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1}},
}

func Forgeron(c *personnage.Etudiant) {
	fmt.Println("Je suis le forgeron, je fabrique ton équipement contre des ressources.")

	for {
		fmt.Println("\nIl te reste", c.Argent, "pièces d'or. Chaque fabrication coûte", coutFabrication, "pièces.")
		for i, r := range recettes {
			fmt.Printf("%d - %s\n", i+1, r.Nom)
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
			fmt.Println("T'as pas de quoi payer, il me faut", coutFabrication, "pièces d'or.")
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