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
