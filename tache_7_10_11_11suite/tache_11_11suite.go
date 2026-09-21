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

func characterCreation(nom, classe string, )string{
	nom = Capitalize(nom)
	niveau := 1 
	maxVie := 0
	vie := 0
	if classe == "Etudiant" {
		maxVie = 100
		vie = maxVie
	} 
	if classe == "Futur_Etudiant" {
		maxVie = 80
		vie = maxVie
	}
	if classe == "Etudiant_B2" {
		maxVie = 120
		vie = maxVie 
	}
	inventaire := []string{}
	personnage.InitCharacter(nom, classe , niveau, maxVie, vie , inventaire)
	return "votre nouveau personnage a été créé avec succès !"
}
