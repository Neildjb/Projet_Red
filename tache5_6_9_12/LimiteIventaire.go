package personnage

func InventaireFull(perso Etudiant) string { 
	if len(perso.Inventaire) >= 10 {
		return "L'item à bien été ajouté!"
	}
	return "Votre inventaire est rempli"
}
