package jeu

import "Projet_Red/personnage"

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
	argent := 65
	skills := []string{"Coup de poing"}

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

	if classe == "Ninja" {
		maxVie = 50
		vie = maxVie
	}

	if classe == "Naruto_Prime" {
		maxVie = 30
		vie = maxVie
	}

	if classe == "Admin4416" {
		return personnage.InitCharacter(
			nom,
			classe,
			100,
			10000,
			10000,
			1000000,
			[]string{"Potion de soin", "Potion de poison", "Potion de guérison du poison", "Potion de PV total", "upgrade1", "upgrade2", "Trophee en or"},
			personnage.Equipment{
				Headgear:  "bandeau frontale ninja",
				BodyArmor: "Manteau Akatsuki",
				FeetArmor: "Bottes de Shinobi",
			},
			[]string{"Coup de poing", "Kunaï", "Rasengan", "Sharingan"},
		)
	}

	personnageCree := personnage.InitCharacter(
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
	personnageCree.Hardcore = classe == "Naruto_Prime"
	return personnageCree
}
