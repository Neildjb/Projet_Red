package personnage
import (
	"Projet_Red/personnage"
)
func InventaireFull(perso personnage.Etudiant) string { 
	if len(perso.Inventaire) >= 10 {
		return "L'item à bien été ajouté!"
	}
	return "Votre inventaire est rempli"
}
