package main

import (
	menu "Projet_Red/Menu"
	"Projet_Red/jeu"
	"Projet_Red/personnage"
	"fmt"
)

func main() {
	fmt.Println("\n========================================")
	fmt.Println("           SHINOBI GENESIS")
	fmt.Println("========================================\n")
	fmt.Println("Bienvenue dans Shinobi Genesis, jeune apprenti ninja !")
	fmt.Println("Affronte des monstres, gagne des pièces et améliore ton personnage.")
	fmt.Println("\n--- Création du personnage ---\n")
	fmt.Println("Écrivez votre pseudo :")

	var choix string
	var choix2 string

	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Erreur de lecture de votre pseudo")
	}

	fmt.Println("\n========================================")
	fmt.Println("         SÉLECTION DE LA CLASSE")
	fmt.Println("========================================\n")
	fmt.Println("La classe que tu choisis définit le mode de difficulté du jeu.")
	fmt.Println("")
	fmt.Println("Kage = Tres facile")
	fmt.Println("Jonin = Facile")
	fmt.Println("Genin = Moyen")
	fmt.Println("Ninja = Tres dur")
	fmt.Println("Naruto_Prime = hardcore (une seule mort)")
	fmt.Println("Admin0000(remplace 0000 par le code secret) = mode de triche")
	fmt.Println("")
	for {
		fmt.Println("Écrivez la class de votre choix :")
		_, err2 := fmt.Scanln(&choix2)
		if err2 != nil {
			fmt.Println("Erreur de lecture de votre classe")
			continue
		}

		choix2 = jeu.Capitalize(choix2)

		switch choix2 {
		case "Kage", "Jonin", "Genin", "Ninja", "Naruto_Prime", "Admin4416":
			c1 := jeu.CharacterCreation(choix, choix2)
			if c1.Hardcore {
				fmt.Println("\n!!! MODE HARDCORE ACTIVÉ !!!")
				fmt.Println("Une seule mort est décisive : la partie s'arrêtera définitivement.")
				fmt.Println("C'est ici que le jeu devient vraiment intéressant.")
				fmt.Println("Bonne chance pour terminer le jeu dans cet état !\n")
			}
			fmt.Println("\n========================================")
			fmt.Println("          VOTRE PERSONNAGE")
			fmt.Println("========================================\n")
			personnage.DisplayInfo(c1)
			fmt.Println()
			menu.Menu(c1)
		default:
			fmt.Println("La classe ne correspond à aucune classe existante.")
		}
	}
}
