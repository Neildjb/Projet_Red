package personnage

import (
	"Projet_Red/personnage"
)

func InventaireFull(perso personnage.Etudiant) bool {
	return len(perso.Inventaire) >= perso.CapaciteInventaire
}
