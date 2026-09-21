package UpgradeInv

import (
	"Projet_Red/personnage"
	"fmt"
)

var maxUpgrades = 2

func AccessInventory(e personnage.Etudiant) {
	fmt.Print(e.Inventaire)
}

func UpgradeInventorySlot(c *personnage.Etudiant) string {
	if c.NbUpgrades >= maxUpgrades {
		return "Erreur : vous avez déjà utilisé vos 2 augmentations d'inventaire."
	}
	c.CapaciteInventaire += 5
	c.NbUpgrades++
	return "Inventaire amélioré."
}
