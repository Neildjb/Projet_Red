package combat

import (
	"fmt"
	"strings"
	"Projet_Red/personnage"
	"Projet_Red/tache8"
)

func degatsAttaque(attaque string) int {
	switch strings.ToLower(strings.TrimSpace(attaque)) {
	case "coup de poing":
		return 10
	default:
		return 0
	}
}

// Combat lance un combat au tour par tour et renvoie true si le joueur gagne.
func Combat(monstre *personnage.Monster, joueur *personnage.Etudiant) bool {

	fmt.Println("Un", monstre.Nom, "apparaît !")

	for joueur.Vie > 0 && monstre.Vie > 0 {
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
