package inventaire

import "Projet_Red/personnage"
import "fmt"

func AccessInventory(e personnage.Etudiant) {
	fmt.Println("Inventaire :", e.Inventaire)
	fmt.Println("Capacité :", len(e.Inventaire), "/", e.CapaciteInventaire)
}
