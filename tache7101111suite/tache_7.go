package tache7101111suite

// TACHE 7 : Marchand

import (
	"Projet_Red/personnage"
	"fmt"
)

func Marchand() {
	fmt.Println("Je suis le marchand, je peux vous vendre des objets pour améliorer vos compétences.")
	fmt.Println("Je peux vous vendre une potion de vie gratuitement.")
	fmt.Println("Tapez AchatPotionDeVie() pour récupérer votre potion de vie.")
}

func AchatPotionDeVie(c *personnage.Etudiant) {
	c.Inventaire = append(c.Inventaire, "potion de vie")
}

func AddInventory(c *personnage.Etudiant, item string){
	c.Inventaire= append(c.Inventaire, item)
}

func RemoveInventory(c *personnage.Etudiant, item string){
	var res []string
	for i,_ :=range c.Inventaire{
		if item == c.Inventaire[i]{

			continue
			
		} else {
			res= append(res,  c.Inventaire[i] )
			fmt.Println(res[i])
		}
	}
	c.Inventaire=res
}

