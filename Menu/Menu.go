package Menu

import (
	"fmt"
	"Projet_Red/personnage"
	"Projet_Red/inventaire"
)

func menu(c personnage.Etudiant) {
	continuer := true

	for continuer {
		fmt.Println("=======MENU=======")
		fmt.Println("")
		fmt.Println("tapez 1 pour afficher les informations du personnages, tapez 2 pour afficher l'inventaire. cd")

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
