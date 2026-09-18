package inventaire

import "Projet_Red/personnage"
import "fmt"

func AccessInventory(e personnage.Etudiant) {
	fmt.Print(e.Inventaire)
}
