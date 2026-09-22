package tache7101111suite

import "Projet_Red/personnage"

// TACHE 11 : Amélioration de la création de personnage
// TACHE 11 - suite

func Capitalize(s string) string {
	tab := []rune(s)
	debut := true

	for i := 0; i < len(tab); i++ {
		if ('a' <= tab[i] && tab[i] <= 'z') || ('A' <= tab[i] && tab[i] <= 'Z') || ('0' <= tab[i] && tab[i] <= '9') {
			if debut {
				if 'a' <= tab[i] && tab[i] <= 'z' {
					tab[i] = tab[i] - 32
				}
				debut = false
			} else {
				if 'A' <= tab[i] && tab[i] <= 'Z' {
					tab[i] = tab[i] + 32
				}
			}
		} else {
			debut = true
		}
	}

	return string(tab)
}

func CharacterCreation(nom, classe string) personnage.Etudiant {

	nom = Capitalize(nom)

	experienceCombat := 0
	maxVie := 0
	vie := 0
	inventaire := []string{}
	equipement := personnage.Equipment{
		Headgear:  "rien",
		BodyArmor: "rien",
		FeetArmor: "rien",
	}
	argent := 100

	if classe == "Kage" {
		maxVie = 300
		vie = maxVie
	}

	if classe == "Jonin" {
		maxVie = 150
		vie = maxVie
	}

	if classe == "Genin" {
		maxVie = 100
		vie = maxVie
	}

	if classe == "ninja" {
		maxVie = 50
		vie = maxVie
	}

	skills := []string{"Coup de poing"}
	c := personnage.InitCharacter(
		nom,
		classe,
		experienceCombat,
		maxVie,
		vie,
		argent,
		inventaire,
		equipement,
		skills,
	)

	return c
}
