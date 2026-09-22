package combat

import (
	"Projet_Red/personnage"
	"Projet_Red/tache8"
	"fmt"
	"strings"
)

func degatsAttaque(attaque string) int {
	switch strings.ToLower(strings.TrimSpace(attaque)) {
	case "coup de poing":
		return 10
	default:
		return 0
	}
}

func displayCombatStatus(monstre *personnage.Monster, joueur *personnage.Etudiant) {
	fmt.Println("Vie de", joueur.Nom, ":", joueur.Vie, "/", joueur.MaxVie)
	fmt.Println("Vie de", monstre.Nom, ":", monstre.Vie, "/", monstre.Max_vie)
}

func collectMonsterDrop(monstre *personnage.Monster, joueur *personnage.Etudiant) {
	if monstre.Drop == "" {
		return
	}

	nombreObjets := len(joueur.Inventaire)
	personnage.AjouterItem(joueur, monstre.Drop)
	if len(joueur.Inventaire) > nombreObjets {
		fmt.Println("Vous récupérez :", monstre.Drop)
	}
}

// Combat lance un combat au tour par tour et renvoie true si le joueur gagne.
func Combat(monstre *personnage.Monster, joueur *personnage.Etudiant) bool {

	fmt.Println("Un", monstre.Nom, "apparaît !")

	for joueur.Vie > 0 && monstre.Vie > 0 {
		displayCombatStatus(monstre, joueur)
		fmt.Println("\nChoisissez une attaque :")
		for numero, skill := range joueur.Skills {
			fmt.Printf("%d - %s (%d dégâts)\n", numero+1, skill, degatsAttaque(skill))
		}

		if len(joueur.Skills) == 0 {
			fmt.Println("Vous n'avez aucune attaque.")
			return false
		}

		var choix int
		for choix < 1 || choix > len(joueur.Skills) {
			fmt.Scan(&choix)
			if choix < 1 || choix > len(joueur.Skills) {
				fmt.Println("Choix invalide, choisissez un numéro dans la liste.")
			}
		}

		attaque := joueur.Skills[choix-1]
		degats := degatsAttaque(attaque)
		monstre.Vie -= degats
		if monstre.Vie < 0 {
			monstre.Vie = 0
		}
		fmt.Println(joueur.Nom, "utilise", attaque, "et inflige", degats, "dégâts à", monstre.Nom)

		if monstre.Vie == 0 {
			fmt.Println(monstre.Nom, "est vaincu !")
			collectMonsterDrop(monstre, joueur)
			return true
		}

		joueur.Vie -= monstre.Points_attaque
		if joueur.Vie < 0 {
			joueur.Vie = 0
		}
		fmt.Println(monstre.Nom, "inflige", monstre.Points_attaque, "dégâts à", joueur.Nom)
		tache8.Isdead(joueur)
	}

	fmt.Println(joueur.Nom, "a perdu le combat.")
	return false
}

func initTrainingMonster() personnage.Monster {
	return personnage.Init_Maxime()
}

func monsterPattern(monstre *personnage.Monster, joueur *personnage.Etudiant, tour int) {
	degats := monstre.Points_attaque
	if tour%3 == 0 {
		degats *= 2
	}

	joueur.Vie -= degats
	if joueur.Vie < 0 {
		joueur.Vie = 0
	}

	fmt.Println(monstre.Nom, "inflige à", joueur.Nom, degats, "de dégâts")
	fmt.Println("Vie de", joueur.Nom, ":", joueur.Vie, "/", joueur.MaxVie)
}

func inventoryTurn(joueur *personnage.Etudiant) bool {
	if len(joueur.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return false
	}

	fmt.Println("\nInventaire :", len(joueur.Inventaire), "/", joueur.CapaciteInventaire)
	for numero, objet := range joueur.Inventaire {
		fmt.Printf("%d - %s\n", numero+1, objet)
	}

	var choixObjet int
	for choixObjet < 1 || choixObjet > len(joueur.Inventaire) {
		fmt.Print("Choisissez un objet : ")
		fmt.Scan(&choixObjet)
		if choixObjet < 1 || choixObjet > len(joueur.Inventaire) {
			fmt.Println("Choix invalide, choisissez un objet dans la liste.")
		}
	}

	index := choixObjet - 1
	objet := joueur.Inventaire[index]
	switch objet {
	case "Potion de soin":
		joueur.Vie += 50
		if joueur.Vie > joueur.MaxVie {
			joueur.Vie = joueur.MaxVie
		}
		fmt.Println(joueur.Nom, "utilise", objet, "et récupère 50 points de vie.")
	case "Potion de poison":
		joueur.Vie -= 20
		if joueur.Vie < 0 {
			joueur.Vie = 0
		}
		fmt.Println(joueur.Nom, "utilise", objet, "et perd 20 points de vie.")
	default:
		fmt.Println("L'objet", objet, "ne peut pas être utilisé pendant le combat.")
		return false
	}

	joueur.Inventaire = append(joueur.Inventaire[:index], joueur.Inventaire[index+1:]...)
	fmt.Println("Vie de", joueur.Nom, ":", joueur.Vie, "/", joueur.MaxVie)
	return true
}

func characterTurn(joueur *personnage.Etudiant, monstre *personnage.Monster) bool {
	for {
		fmt.Println("\nMenu")
		fmt.Println("1 - Attaquer")
		fmt.Println("2 - Inventaire")

		var choix string
		fmt.Scan(&choix)

		switch choix {
		case "1":
			if len(joueur.Skills) == 0 {
				fmt.Println("Vous n'avez aucun sort.")
				continue
			}

			fmt.Println("\nSorts disponibles :")
			for numero, sort := range joueur.Skills {
				fmt.Printf("%d - %-20s %d dégâts\n", numero+1, sort, degatsAttaque(sort))
			}

			var choixSort int
			for choixSort < 1 || choixSort > len(joueur.Skills) {
				fmt.Scan(&choixSort)
				if choixSort < 1 || choixSort > len(joueur.Skills) {
					fmt.Println("Choix invalide, choisissez un sort dans la liste.")
				}
			}

			sort := joueur.Skills[choixSort-1]
			degats := degatsAttaque(sort)
			monstre.Vie -= degats
			if monstre.Vie < 0 {
				monstre.Vie = 0
			}
			fmt.Println(joueur.Nom, "utilise", sort, "et inflige", degats, "dégâts à", monstre.Nom)
			fmt.Println("Vie de", monstre.Nom, ":", monstre.Vie, "/", monstre.Max_vie)
			return monstre.Vie > 0
		case "2":
			if inventoryTurn(joueur) {
				return true
			}
		default:
			fmt.Println("Choix invalide, choisissez Attaquer ou Inventaire.")
		}
	}
}

func trainingFight(joueur *personnage.Etudiant) bool {
	monstre := initTrainingMonster()
	tour := 1

	fmt.Println("Un", monstre.Nom, "apparaît !")
	for joueur.Vie > 0 && monstre.Vie > 0 {
		fmt.Println("\nTour", tour)
		displayCombatStatus(&monstre, joueur)
		if !characterTurn(joueur, &monstre) {
			fmt.Println(monstre.Nom, "est vaincu !")
			collectMonsterDrop(&monstre, joueur)
			return true
		}

		monsterPattern(&monstre, joueur, tour)
		tour++
	}

	fmt.Println(joueur.Nom, "a perdu le combat.")
	return false
}

func TrainingFight(joueur *personnage.Etudiant) bool {
	return trainingFight(joueur)
}
