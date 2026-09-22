package spellbook

import (
	"Projet_Red/personnage"
	"fmt"
)

// Ajoute ici les trois sorts que le joueur pourra apprendre.
var spellbook = []string{"Kunaï", "Rasengan", "Sharingan"}

func LearnSpell(joueur *personnage.Etudiant, objet, sort string) bool {
	if !containsSpell(sort) {
		fmt.Println("Ce sort n'est pas disponible dans le Spellbook.")
		return false
	}

	position := -1
	for i, objetPossede := range joueur.Inventaire {
		if objetPossede == objet {
			position = i
			break
		}
	}
	if position == -1 {
		fmt.Println("Vous n'avez pas de", objet, "dans votre inventaire.")
		return false
	}

	joueur.Inventaire = append(joueur.Inventaire[:position], joueur.Inventaire[position+1:]...)
	joueur.Skills = append(joueur.Skills, sort)
	fmt.Println("Vous apprenez", sort, ".")
	return true
}

func containsSpell(sort string) bool {
	for _, sortDisponible := range spellbook {
		if sortDisponible == sort {
			return true
		}
	}
	return false
}

func SpellDamage(sort string) int {
	switch sort {
	case "Kunaï":
		return 20
	case "Rasengan":
		return 35
	case "Sharingan":
		return 100
	default:
		return 0
	}
}
