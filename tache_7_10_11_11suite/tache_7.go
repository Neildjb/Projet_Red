package tache7101111suite

// TACHE 7 : Marchand

import (
	"fmt"

	"Projet_Red/personnage"
)

func Marchand() {
	fmt.Println("Je suis le marchand, je peux vous vendre des objets pour améliorer vos compétences.")
	fmt.Println("Je peux vous vendre une potion de vie gratuitement.")
	fmt.Println("Tapez AchatPotionDeVie() pour récupérer votre potion de vie.")
}

func AchatPotionDeVie() {
	personnage.Etudiant.Inventaire = append(personnage.Etudiant.Inventaire, "potion de vie")
}
