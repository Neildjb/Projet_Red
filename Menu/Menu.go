package Menu

import (
	"fmt"
	"Projet_Red/personnage"
	"Projet_Red/inventaire"
	"Projet_Red/inventaire"
)


func menu(c personnage.Etudiant) {
	liste_objets:= ["kunaï","lance","spellbook","potion de vie",]
	
	continuer := true

	for continuer {
		fmt.Println("=======MENU=======")
		fmt.Println("")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder à l'inventaire")
		fmt.Println("3 - Voir ce que vends le Marchand")
		fmt.Println("3 - Quitter")
		fmt.Println("4 - Retour")

		var choix string
		fmt.Scanln(&choix)
		

		switch choix {
		case "1":
			personnage.DisplayInfo(c)
		case "2":
			inventaire.AccessInventory(c)
		case "3":
			fmt.Println("Bienvenue chez le meilleur marchand de tout Konoha, tout a  un prix voici ce que je te propose : ")
			fmt.Println(liste_objets)
			fmt.Scanln()
		case "4":
			continuer = false
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}