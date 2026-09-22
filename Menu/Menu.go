package menu

import (
	"Projet_Red/NPC"
	"Projet_Red/combat"
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
		fmt.Println("4 - Voir ce que vend le Forgeron")
		fmt.Println("5 - Tuto combat")
		fmt.Println("6 - Arène de combat")
		fmt.Println("7 - École des classes")
		fmt.Println("8 - Quêtes")
		fmt.Println("9 - Quitter")

		var choix string
		fmt.Scan(&choix)

		switch choix {
		case "1":
			personnage.DisplayInfo(c)
		case "2":
			inventaire.AccessInventory(&c)
		case "3":
			npc.Marchand(&c)
		case "4":
			npc.Forgeron(&c)
		case "5":
			combat.TrainingFight(&c)
			if c.GameOver {
				return
			}
		case "6":
			arenaMenu(&c)
			if c.GameOver {
				return
			}
		case "7":
			ecole(&c)
			if c.GameOver {
				return
			}
		case "8":
			quetesFinJeu(c)
		case "9":
			continuer = false
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}

func quetesFinJeu(joueur personnage.Etudiant) {
	capaciteComplete := joueur.CapaciteInventaire >= 20
	argentComplete := joueur.Argent >= 300
	experienceComplete := joueur.ExperienceCombat > 30
	nombreSorts := nombreDeSorts(joueur)
	sortsComplets := nombreSorts == 4
	tropheeObtenu := possedeObjet(joueur, "Trophee en or")
	nombreEquipements := nombreDEquipements(joueur)
	equipementComplet := nombreEquipements == 3

	fmt.Println("\n=== Quêtes ===")
	fmt.Println("Ces quêtes représentent tout ce qu'il faut découvrir pour terminer le jeu.")
	afficherQuete("Capacité d'inventaire", fmt.Sprintf("%d / 20", joueur.CapaciteInventaire), capaciteComplete)
	afficherQuete("Pièces d'or", fmt.Sprintf("%d / 300", joueur.Argent), argentComplete)
	afficherQuete("Expérience de combat", fmt.Sprintf("%d / 30 minimum", joueur.ExperienceCombat), experienceComplete)
	afficherQuete("Sorts possédés", fmt.Sprintf("%d / 4", nombreSorts), sortsComplets)
	afficherQuete("Trophée final", fmt.Sprintf("%d / 1", boolVersNombre(tropheeObtenu)), tropheeObtenu)
	afficherQuete("Équipements équipés", fmt.Sprintf("%d / 3", nombreEquipements), equipementComplet)

	if capaciteComplete && argentComplete && experienceComplete && sortsComplets && tropheeObtenu && equipementComplet {
		fmt.Println("\nFélicitations ! Tu as terminé le jeu et tout découvert.")
		fmt.Println("1 - Afficher les crédits")
		fmt.Println("2 - Retour")
		var choix string
		fmt.Scan(&choix)
		if choix == "1" {
			affiche_credit()
		}
	} else {
		fmt.Println("\nToutes les quêtes ne sont pas encore terminées.")
	}
}

func affiche_credit() {
	fmt.Println(`
==================== CRÉDITS ====================

Un projet réalisé dans le cadre du Projet RED

------------------ Campus Ynov Bordeaux ------------------
Bachelor 1 — 2026

---------------- Équipe de développement ----------------
Mathias Fontagne
Développement • Gameplay • Conception

[Nom de l'élève 2]
Développement • Gameplay • Conception

[Nom de l'élève 3]
Développement • Gameplay • Conception

----------------------- Mentors --------------------------
Un grand merci à nos deux mentors pour leur accompagnement,
leurs conseils et leur aide tout au long du projet :

Maxime & Sarha
Mentor du projet

------------------ À propos du projet --------------------
Ce jeu a été conçu, développé et finalisé en une semaine,
dans le cadre du Projet RED au sein du Campus Ynov Bordeaux.

Un projet réalisé par trois étudiants de Bachelor 1, avec pour
objectif de concevoir un jeu vidéo complet dans un temps limité,
de l'idée initiale jusqu'à sa réalisation finale.

------------------------ Merci ----------------------------
Merci à toutes les personnes ayant contribué, directement ou
indirectement, à la réalisation de ce projet.

Merci d'avoir joué !

© 2026 — Projet RED • Ynov Bordeaux`)
}

func afficherQuete(description string, progression string, terminee bool) {
	statut := "En cours"
	if terminee {
		statut = "Réalisée"
	}
	fmt.Println("-", description, ":", progression, "(", statut+")")
}

func nombreDeSorts(joueur personnage.Etudiant) int {
	sortsRequis := []string{"Coup de poing", "Kunaï", "Rasengan", "Sharingan"}
	nombre := 0
	for _, sortRequis := range sortsRequis {
		for _, sort := range joueur.Skills {
			if sort == sortRequis {
				nombre++
				break
			}
		}
	}
	return nombre
}

func nombreDEquipements(joueur personnage.Etudiant) int {
	nombre := 0
	if joueur.Equipement.Headgear != "" && joueur.Equipement.Headgear != "rien" {
		nombre++
	}
	if joueur.Equipement.BodyArmor != "" && joueur.Equipement.BodyArmor != "rien" {
		nombre++
	}
	if joueur.Equipement.FeetArmor != "" && joueur.Equipement.FeetArmor != "rien" {
		nombre++
	}
	return nombre
}

func boolVersNombre(valeur bool) int {
	if valeur {
		return 1
	}
	return 0
}

func possedeObjet(joueur personnage.Etudiant, objetRecherche string) bool {
	for _, objet := range joueur.Inventaire {
		if objet == objetRecherche {
			return true
		}
	}
	return false
}

func ecole(joueur *personnage.Etudiant) {
	const prixEntree = 100

	fmt.Println("\n=== École des classes ===")
	fmt.Println("Si tu te trouves trop faible, tu peux tenter de changer pour la classe suivante.")
	fmt.Println("L'entrée coûte", prixEntree, "pièces d'or et un combat sera lancé.")

	var nouvelleClasse string
	var monstre personnage.Monster
	var maxVie int

	switch joueur.Classe {
	case "Naruto Prime":
		nouvelleClasse = "Ninja"
		monstre = personnage.Init_Maxime()
		maxVie = 50
	case "Ninja":
		nouvelleClasse = "Genin"
		monstre = personnage.Init_Ninjas_déserteurs()
		maxVie = 100
	case "Genin":
		nouvelleClasse = "Jonin"
		monstre = personnage.Init_Golems_de_chakra()
		maxVie = 150
	case "Jonin":
		nouvelleClasse = "Kage"
		monstre = personnage.Init_Demon_a_queue()
		maxVie = 300
	case "Kage":
		fmt.Println("Tu as déjà atteint la classe la plus élevée.")
		return
	case "Admin4416":
		fmt.Println("Tu es déjà dans la classe maximale, aucune classe supérieure n'est accessible.")
		return
	default:
		fmt.Println("Ta classe actuelle ne permet pas de changer de classe à l'école.")
		return
	}

	fmt.Println("Tu peux tenter de devenir", nouvelleClasse, "en affrontant", monstre.Nom+".")
	if joueur.Argent < prixEntree {
		fmt.Println("Tu n'as pas assez d'argent pour entrer dans l'école.")
		return
	}

	var choix string
	fmt.Println("1 - Entrer dans l'école")
	fmt.Println("2 - Retour")
	fmt.Scanln(&choix)
	if choix != "1" {
		return
	}

	joueur.Argent -= prixEntree
	fmt.Println("Les", prixEntree, "pièces d'or ont été dépensées.")
	if !combat.Combat(&monstre, joueur) {
		fmt.Println("Tu gardes ta classe", joueur.Classe, ".")
		return
	}

	joueur.Classe = nouvelleClasse
	joueur.MaxVie = maxVie
	joueur.Vie = maxVie
	fmt.Println("Félicitations ! Tu es maintenant", joueur.Classe, "avec", joueur.Vie, "PV.")
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
