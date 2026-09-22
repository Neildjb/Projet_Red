package inventaire

import (
	"Projet_Red/personnage"
	"Projet_Red/tache7101111suite"
	"fmt"
	"strconv"
)

func AccessInventory(e personnage.Etudiant) {
	fmt.Println("Inventaire :", e.Inventaire)
	fmt.Println("Capacité :", len(e.Inventaire), "/", e.CapaciteInventaire)
	fmt.Println("1 - Supprimer un objet de l'inventaire ")
	fmt.Println("2 - Retour")
	var choix string
	fmt.Scanln(&choix)

	switch choix {
	case "1":
		fmt.Println("Que souhaites-tu retirer, tapes le chiffre coreespondant à la place de l'item.")
		for i, r := range e.Inventaire {
			println(strconv.Itoa(i+1) + ": " + r)
		}
		for {
			var choix_sup string
			fmt.Scanln(choix_sup)
			n, err := strconv.Atoi(choix_sup)
			if err != nil || n < 0 || n > len(choix_sup) {
				fmt.Println("Ce n'est pas au menu, gamin.")
				continue

			} else {

				for i := 0; i < len(e.Inventaire); i++ {
					if e.Inventaire[i] == e.Inventaire[n-1] {
						tache7101111suite.RemoveInventory(&e, e.Inventaire[i])

					}
				}

			}

		}

	case "2":
		return

	}

}
