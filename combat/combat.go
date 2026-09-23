package combat

import (
	"Projet_Red/etat"
	"Projet_Red/gestionmort"
	"Projet_Red/personnage"
	"Projet_Red/potions"
	"Projet_Red/sorts"
	"fmt"
)

func displayCombatStatus(monstre *personnage.Monster, joueur *personnage.Etudiant) {
	fmt.Println("Vie de", joueur.Nom, ":", joueur.Vie, "/", joueur.MaxVie)
	fmt.Println("Vie de", monstre.Nom, ":", monstre.Vie, "/", monstre.Max_vie)
}

func collectMonsterDrop(monstre *personnage.Monster, joueur *personnage.Etudiant) {
	for _, drop := range monstre.Drop {
		nombreObjets := len(joueur.Inventaire)
		personnage.AjouterItem(joueur, drop)
		if len(joueur.Inventaire) > nombreObjets {
			fmt.Println("Vous récupérez :", drop)
		}
	}
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

func Combat(monstre *personnage.Monster, joueur *personnage.Etudiant) bool {
	etat.MonstreActuel = monstre
	defer func() {
		etat.MonstreActuel = nil
		joueur.Poison = false
	}()

	tour := 1

	fmt.Println("Un", monstre.Nom, "apparaît !")
	if monstre.Nom == "Demon a queue" {
		joueur.Poison = true
		fmt.Println(monstre.Nom, "empoisonne", joueur.Nom, "dès le début du combat.")
	}

	for joueur.Vie > 0 && monstre.Vie > 0 {
		if monstre.Nom == "Madara" && tour%3 == 0 {
			joueur.Poison = true
			fmt.Println(monstre.Nom, "empoisonne", joueur.Nom, "au tour", tour, ".")
		}
		if joueur.Poison {
			joueur.Vie -= 5
			if joueur.Vie < 0 {
				joueur.Vie = 0
			}
			fmt.Println(joueur.Nom, "perd 5 PV à cause du poison.")
			if joueur.Vie == 0 {
				gestionmort.Isdead(joueur)
				break
			}
		}
		if monstre.Poison {
			monstre.Vie -= 5
			if monstre.Vie < 0 {
				monstre.Vie = 0
			}
			fmt.Println(monstre.Nom, "perd 5 PV à cause du poison.")
			if monstre.Vie == 0 {
				fmt.Println(monstre.Nom, "est vaincu !")
				collectMonsterDrop(monstre, joueur)
				collectMonsterMoney(monstre, joueur)
				collectCombatExperience(monstre, joueur)
				return true
			}
		}
		displayCombatStatus(monstre, joueur)
		if !characterTurn(joueur, monstre) {
			fmt.Println(monstre.Nom, "est vaincu !")
			collectMonsterDrop(monstre, joueur)
			collectMonsterMoney(monstre, joueur)
			collectCombatExperience(monstre, joueur)
			return true
		}

		joueur.Vie -= monstre.Points_attaque
		if joueur.Vie < 0 {
			joueur.Vie = 0
		}
		fmt.Println(monstre.Nom, "inflige", monstre.Points_attaque, "dégâts à", joueur.Nom)
		gestionmort.Isdead(joueur)
		tour++
	}

	fmt.Println(joueur.Nom, "a perdu le combat.")
	return false
}

func inventoryTurn(joueur *personnage.Etudiant, monstre *personnage.Monster) bool {
	if len(joueur.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return false
	}

	fmt.Println("\nInventaire :", len(joueur.Inventaire), "/", joueur.CapaciteInventaire)
	for numero, objet := range joueur.Inventaire {
		fmt.Printf("%d - %s\n", numero+1, objet)
	}
	fmt.Println("0 - Retour")

	var choixObjet int
	for {
		fmt.Print("Choisissez un objet : ")
		fmt.Scan(&choixObjet)
		if choixObjet == 0 {
			return false
		}
		if choixObjet >= 1 && choixObjet <= len(joueur.Inventaire) {
			break
		}
		if choixObjet < 1 || choixObjet > len(joueur.Inventaire) {
			fmt.Println("Choix invalide, choisissez un objet dans la liste.")
		}
	}

	index := choixObjet - 1
	objet := joueur.Inventaire[index]
	switch objet {
	case "Potion de soin":
		*joueur = potions.TakePot(*joueur)
		fmt.Println("Vie de", joueur.Nom, ":", joueur.Vie, "/", joueur.MaxVie)
		return true
	case "Potion de poison":
		joueur.Inventaire = append(joueur.Inventaire[:index], joueur.Inventaire[index+1:]...)
		monstre.Poison = true
		fmt.Println(joueur.Nom, "lance une potion de poison sur", monstre.Nom, ".")
		return true
	case "Potion de guérison du poison":
		if !joueur.Poison {
			fmt.Println(joueur.Nom, "n'est pas empoisonné.")
			return false
		}
		joueur.Inventaire = append(joueur.Inventaire[:index], joueur.Inventaire[index+1:]...)
		joueur.Poison = false
		fmt.Println(joueur.Nom, "ne souffre plus du poison.")
		return true
	case "Potion de PV total":
		*joueur = potions.FullHealthPot(*joueur)
		return true
	default:
		fmt.Println("L'objet", objet, "est visible mais seuls les potions peuvent être utilisées pendant le combat.")
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
				fmt.Printf("%d - %-20s %d dégâts\n", numero+1, sort, sorts.SpellDamage(sort))
			}
			fmt.Println("0 - Retour")

			var choixSort int
			for {
				fmt.Scan(&choixSort)
				if choixSort == 0 {
					break
				}
				if choixSort >= 1 && choixSort <= len(joueur.Skills) {
					break
				}
				if choixSort < 1 || choixSort > len(joueur.Skills) {
					fmt.Println("Choix invalide, choisissez un sort dans la liste.")
				}
			}
			if choixSort == 0 {
				continue
			}

			sort := joueur.Skills[choixSort-1]
			degats := sorts.SpellDamage(sort)
			monstre.Vie -= degats
			if monstre.Vie < 0 {
				monstre.Vie = 0
			}
			fmt.Println(joueur.Nom, "utilise", sort, "et inflige", degats, "dégâts à", monstre.Nom)
			fmt.Println("Vie de", monstre.Nom, ":", monstre.Vie, "/", monstre.Max_vie)
			return monstre.Vie > 0
		case "2":
			if inventoryTurn(joueur, monstre) {
				return true
			}
		default:
			fmt.Println("Choix invalide, choisissez Attaquer ou Inventaire.")
		}
	}
}

func TrainingFight(joueur *personnage.Etudiant) bool {
	monstre := personnage.Init_Maxime()
	if joueur.TutorielFait {
		monstre.ArgentDrop = 2
	} else {
		monstre.ArgentDrop = 20
		joueur.TutorielFait = true
	}
	return Combat(&monstre, joueur)
}
