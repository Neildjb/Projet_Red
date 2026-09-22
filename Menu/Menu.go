package Menu

import (
	"Projet_Red/combat"
	"Projet_Red/forgeron"
	"Projet_Red/inventaire"
	"Projet_Red/personnage"
	"Projet_Red/tache7101111suite"
	"fmt"
)

func Menu(c personnage.Etudiant) {
	continuer := true

	for continuer {
		fmt.Println("=======MENU=======")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder à l'inventaire")
		fmt.Println("3 - Voir ce que vend le Marchand")
		fmt.Println("4 - Voir ce que vend le Forgeron")
		fmt.Println("5 - Tuto combat")
		fmt.Println("6 - Quitter")

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
			forgeron.Forgeron(&c)
		case "5":
			var nomObjet string
			fmt.Println("Quel objet veux-tu équiper ?")
			fmt.Scanln(&nomObjet)
			inventaire.AccessInventory(c)
			forgeron.Equiper(&c, nomObjet)
		case "6":
			combat.TrainingFight(&c)
		case "7":
			continuer = false
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
