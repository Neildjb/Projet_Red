package Menu

import (
	"Projet_Red/tache7101111suite"
	"Projet_Red/inventaire"
	"Projet_Red/personnage"
	"fmt"

)


func Menu(c personnage.Etudiant) {
	continuer := true

	for continuer {
		fmt.Println("=======MENU=======")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder à l'inventaire")
		fmt.Println("3 - Voir ce que vend le Marchand")
		fmt.Println("4 - Quitter")

		var choix string
		fmt.Scanln(&choix)
		

		switch choix {
		case "1":
			personnage.DisplayInfo(c)
		case "2":
			inventaire.AccessInventory(c)
		case "3":
			tache7101111suite.Marchand(&c)
		case "4":
			continuer = false
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}