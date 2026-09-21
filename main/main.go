package main

import (
	"Projet_Red/personnage"
	"fmt"
	"Projet_Red/tache7101111suite"
)

func main() {
	fmt.Println("Bienvenue dans le jeu _")
	fmt.Println("Vous êtes un Ninja et vous devez combattre des monstres pour devenir le plus fort.")
	fmt.Println("Commencons par la creation de votre personnage.")
	fmt.Println("Ecrit moi ton nom :")	
	
	var choix string
	var choix2 string

	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Erreur de lecture sur le blaze")
	}

	fmt.Println("La classe que tu choisi difini le mode de difficulte du Jeu, Kage = facile, Jonin = moyen, Genin = difficile, ninja = impossible")
	fmt.Println("ecrit moi ton class(difficulté du jeu) :")	

	_, err2 := fmt.Scanln(&choix2)

	if err2 != nil {
		fmt.Println("Erreur de lecture sur le deuxième choix")
	}

	c1 := tache7101111suite.CharacterCreation(choix, choix2)

	fmt.Println("Voici ton personnage :")
	personnage.DisplayInfo(c1)

	fmt.Print("Commencons par vous expliquer coment acceder au menu")
	fmt.Println("il faut ecrire 'menu()' et appuyer sur entrer")
}
