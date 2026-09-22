package main

import (
	"Projet_Red/jeu"
	"Projet_Red/menu"
	"Projet_Red/personnage"
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
	fmt.Println("Kage = Tres facile")
	fmt.Println("Jonin = Facile")
	fmt.Println("Genin = Complexe")
	fmt.Println("Ninja = Tres complexe")
	fmt.Println("Naruto Prime = hardcore (une seule mort, 30 PV)")
	fmt.Println("Admin0000(remplace 0000 par le code secret) = mode de triche")
	fmt.Println("")
	for {
		fmt.Println("Écrivez la class de votre choix :")
		_, err2 := fmt.Scanln(&choix2)
		if err2 != nil {
			fmt.Println("Erreur de lecture de votre classe")
			continue
		}

		switch choix2 {
		case "Kage", "Jonin", "Genin", "Ninja", "Naruto Prime", "Admin4416":
			goto classeValide
		default:
			fmt.Println("La classe ne correspond à aucune classe existante.")
		}
	}

classeValide:
	c1 := jeu.CharacterCreation(choix, choix2)
	if c1.Hardcore {
		fmt.Println("Mode hardcore activé : une seule mort est décisive, la partie s'arrêtera définitivement.")
	}

	fmt.Println("================Character================")
	fmt.Println("")
	fmt.Println("Voici les données de votre personnage :")
	fmt.Println("")
	personnage.DisplayInfo(c1)
	menu.Menu(c1)
}
