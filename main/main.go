package main

import (
	"Projet_Red/Menu"
	"Projet_Red/personnage"
	"Projet_Red/tache7101111suite"
	"fmt"
)

func main() {
	fmt.Println("================Shinobi Genesis================")
	fmt.Println("")
	fmt.Println("Bienvenue dans Shinobi Genesis jeune apprenti Ninja,")
	fmt.Println("vous êtes un Ninja et vous devez combattre des monstres pour obtenir des pièces afin d'améliorer vos compétences pour devenir plus fort!")
	fmt.Println("")
	fmt.Println("Tout d'abord, commençons par la création de votre personnage.")
	fmt.Println("Écrivez votre pseudo :")	
	
	var choix string
	var choix2 string

	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Erreur de lecture de votre pseudo")
	}

	fmt.Println("================Sélection de votre Classe================")
	fmt.Println("")	
	fmt.Println("La classe que tu choisis définit le mode de difficulté du jeu.")
	fmt.Println("")
	fmt.Println("Kage = facile")
	fmt.Println("Jonin = moyen")
	fmt.Println("Genin = difficile")
	fmt.Println("ninja = impossible")
	fmt.Println("Moderateur = mode de triche")
	fmt.Println("")
	fmt.Println("Écrivez la class de votre choix :")	

	_, err2 := fmt.Scanln(&choix2)

	if err2 != nil {
		fmt.Println("Erreur de lecture de votre classe")
	}

	c1 := tache7101111suite.CharacterCreation(choix, choix2)

	fmt.Println("================Character================")
	fmt.Println("")
	fmt.Println("Voici les données de votre personnage :")
	fmt.Println("")
	personnage.DisplayInfo(c1)
	Menu.Menu(c1)
}
