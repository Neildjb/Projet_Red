package spellbook

import (
	"Projet_Red/personnage"
	"fmt"
)

const spellbookName = "Kunaï"
const spellbookSpell = "Sort du Kunaï"

func UseSpellBook(joueur *personnage.Etudiant) bool {
	position := -1
	for i, objet := range joueur.Inventaire {
		if objet == spellbookName {
			position = i
			break
		}
	}
	if position == -1 {
		fmt.Println("Vous n'avez pas de Kunaï dans votre inventaire.")
		return false
	}

	for _, sort := range joueur.Skills {
		if sort == spellbookSpell {
			fmt.Println("Vous connaissez déjà le sort du Kunaï.")
			return false
		}
	}

	joueur.Inventaire = append(joueur.Inventaire[:position], joueur.Inventaire[position+1:]...)
	joueur.Skills = append(joueur.Skills, spellbookSpell)
	joueur.Vie -= 5
	if joueur.Vie < 0 {
		joueur.Vie = 0
	}
	fmt.Println("Vous apprenez", spellbookSpell, "et perdez 5 PV.")
	return true
}

func UseShuriken(joueur *personnage.Etudiant) bool {
	position := -1
	for i, objet := range joueur.Inventaire {
		if objet == "Shuriken" {
			position = i
			break
		}
	}
	if position == -1 {
		fmt.Println("Vous n'avez pas de Shuriken dans votre inventaire.")
		return false
	}

	const sortDestruction = "Attaque de destruction"
	for _, sort := range joueur.Skills {
		if sort == sortDestruction {
			fmt.Println("Vous connaissez déjà", sortDestruction, ".")
			return false
		}
	}

	joueur.Inventaire = append(joueur.Inventaire[:position], joueur.Inventaire[position+1:]...)
	joueur.Skills = append(joueur.Skills, sortDestruction)
	fmt.Println("Vous découvrez", sortDestruction, ": 100 dégâts.")
	return true
}
