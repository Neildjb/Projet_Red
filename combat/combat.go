package combat

import (
	"Projet_Red/personnage"
	spellbook "Projet_Red/spellBook"
	potions "Projet_Red/tache5_6_9_12"
	"Projet_Red/tache8"
	"fmt"
	"strings"
)

func degatsAttaque(attaque string) int {
	switch strings.ToLower(strings.TrimSpace(attaque)) {
	case "coup de poing":
		return 10
	case "sort du kunaï":
		return 20
	case "attaque de destruction":
		return 100
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

func collectMonsterTrophy(monstre *personnage.Monster, joueur *personnage.Etudiant) {
	if monstre.Trophee == "" {
		return
	}

	nombreObjets := len(joueur.Inventaire)
	personnage.AjouterItem(joueur, monstre.Trophee)
	if len(joueur.Inventaire) > nombreObjets {
		fmt.Println("Vous récupérez le", monstre.Trophee, ".")
	} else {
		fmt.Println("Le trophée ne peut pas être ajouté : inventaire plein.")
	}
	fmt.Println("Félicitations ! Vous avez remporté le trophée final !")
}

func collectMonsterMoney(monstre *personnage.Monster, joueur *personnage.Etudiant) {
	if monstre.ArgentDrop <= 0 {
		return
	}

	joueur.Argent += monstre.ArgentDrop
	fmt.Println("Vous gagnez", monstre.ArgentDrop, "pièces d'or.")
	fmt.Println("Argent total :", joueur.Argent, "pièces d'or")
}

func collectCombatExperience(monstre *personnage.Monster, joueur *personnage.Etudiant) {
	joueur.ExperienceCombat += monstre.ExperienceDrop
	fmt.Println("Expérience de combat : +", monstre.ExperienceDrop)
}

// Combat lance un combat au tour par tour et renvoie true si le joueur gagne.
func Combat(monstre *personnage.Monster, joueur *personnage.Etudiant) bool {

	fmt.Println("Un", monstre.Nom, "apparaît !")

	for joueur.Vie > 0 && monstre.Vie > 0 {
		displayCombatStatus(monstre, joueur)
		if !characterTurn(joueur, monstre) {
			fmt.Println(monstre.Nom, "est vaincu !")
			collectMonsterDrop(monstre, joueur)
			collectMonsterTrophy(monstre, joueur)
			collectMonsterMoney(monstre, joueur)
			collectCombatExperience(monstre, joueur)
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
	case "Potion de soin", "potion de vie":
		*joueur = potions.TakePot(*joueur)
		fmt.Println("Vie de", joueur.Nom, ":", joueur.Vie, "/", joueur.MaxVie)
		return true
	case "Potion de poison":
		*joueur = potions.PoisonPot(*joueur)
		fmt.Println("Vie de", joueur.Nom, ":", joueur.Vie, "/", joueur.MaxVie)
		return true
	case "Kunaï":
		return spellbook.UseSpellBook(joueur)
	case "Shuriken":
		return spellbook.UseShuriken(joueur)
	default:
		fmt.Println("L'objet", objet, "ne peut pas être utilisé pendant le combat.")
		return false
	}

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
			collectMonsterMoney(&monstre, joueur)
			collectCombatExperience(&monstre, joueur)
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
