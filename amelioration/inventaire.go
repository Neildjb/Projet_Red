package amelioration

import (
	"Projet_Red/personnage"
)

const capaciteMaxInventaire = 20

func UpgradeInventorySlot(c *personnage.Etudiant) string {
	if c.CapaciteInventaire >= capaciteMaxInventaire {
		return "Erreur : vous avez déjà utilisé vos 2 augmentations d'inventaire."
	}
	c.CapaciteInventaire += 5
	return "Inventaire amélioré. ( + 5 inventaire )"
}
