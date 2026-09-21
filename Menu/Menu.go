package Menu

import (
	"Projet_Red/inventaire"
	"Projet_Red/inventaire"
	"Projet_Red/personnage"
	"fmt"

)


func menu(c personnage.Etudiant) {
	liste_objets:= ["kunaï","lance","spellbook","potion de vie",]
	
	continuer := true

	for continuer {
		fmt.Println("=======MENU=======")
		fmt.Println("")
		fmt.Println("tapez 1 pour afficher les informations du personnages, tapez 2 pour afficher l'inventaire. cd")

		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder à l'inventaire")
		fmt.Println("3 - Voir ce que vends le Marchand")
		fmt.Println("3 - Quitter")
		fmt.Println("4 - Retour")

		var choix string
		fmt.Scanln(&choix)
		if choix == "1" {
			personnage.DisplayInfo(c)
		} else if choix == "2" {
			inventaire.AccessInventory(c)
		} else if choix == "3" {
			continuer = false
		} else {
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
