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
		fmt.Println("5 - s'équiper ")
		fmt.Println("6 - Tuto combat")
		fmt.Println("7 - Arène de combat")
		fmt.Println("8 - Quitter")

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
			forgeron.MenuEquiper(&c)
		case "6":
			combat.TrainingFight(&c)
		case "7":
			arenaMenu(&c)
		case "8":
			continuer = false
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}

func arenaMenu(joueur *personnage.Etudiant) {
	for {
		fmt.Println("\n=== Arène de combat ===")
		fmt.Println("1 - Ninjas déserteurs (expérience >= 1)")
		fmt.Println("2 - Golems de chakra (expérience >= 3)")
		fmt.Println("3 - Démon à queue (expérience >= 5)")
		fmt.Println("4 - Madara (expérience >= 10)")
		fmt.Println("5 - Retour au menu")

		var choix string
		fmt.Scanln(&choix)

		var monstre personnage.Monster
		seuilExperience := 0
		switch choix {
		case "1":
			monstre = personnage.Init_Ninjas_déserteurs()
			seuilExperience = 1
		case "2":
			monstre = personnage.Init_Golems_de_chakra()
			seuilExperience = 3
		case "3":
			monstre = personnage.Init_Demon_a_queue()
			seuilExperience = 5
		case "4":
			monstre = personnage.Init_Madara()
			seuilExperience = 10
		case "5":
			return
		default:
			fmt.Println("Choix invalide, réessaie.")
			continue
		}

		if joueur.ExperienceCombat < seuilExperience {
			fmt.Println("Accès refusé. Il faut au moins", seuilExperience, "d'expérience de combat.")
			continue
		}

		combat.Combat(&monstre, joueur)
		return
	}
}
