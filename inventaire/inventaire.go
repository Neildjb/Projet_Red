package inventaire

import (
	"Projet_Red/UpgradeInv"
	"Projet_Red/forgeron"
	"Projet_Red/personnage"
	spellbook "Projet_Red/spellBook"
	potions "Projet_Red/tache5_6_9_12"
	"Projet_Red/tache7101111suite"
	"fmt"
	"strconv"
	"strings"
)

func AccessInventory(e *personnage.Etudiant) {
	for {
		fmt.Println("Inventaire :", strings.Join(e.Inventaire, ", "))
		fmt.Println("Capacité :", len(e.Inventaire), "/", e.CapaciteInventaire)
		fmt.Println("1 - Supprimer un objet de l'inventaire ")
		fmt.Println("2 - Utiliser un objet")
		fmt.Println("3 - Retour")
		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			if len(e.Inventaire) == 0 {
				fmt.Println("Votre inventaire est vide.")
				continue
			}
			fmt.Println("Que souhaites-tu retirer, tapes le chiffre coreespondant à la place de l'item.")
			for i, r := range e.Inventaire {
				fmt.Println(strconv.Itoa(i+1) + ": " + r)
			}

			for {
				var choix_sup string
				fmt.Scanln(&choix_sup)
				n, err := strconv.Atoi(choix_sup)
				if err != nil || n < 1 || n > len(e.Inventaire) {
					fmt.Println("Ce n'est pas au menu, gamin.")
					continue

				} else {

					fmt.Println("Je vais retirer " + e.Inventaire[n-1] + " de l'inventaire")
					tache7101111suite.RemoveInventory(e, e.Inventaire[n-1])
					fmt.Println(e.Inventaire)
					return
				}

			}

		case "2":
			utiliserObjet(e)
		case "3":
			return

		}
	}

}

func utiliserObjet(e *personnage.Etudiant) {
	if len(e.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	fmt.Println("Quel objet veux-tu utiliser ?")
	for i, objet := range e.Inventaire {
		fmt.Printf("%d - %s\n", i+1, objet)
	}

	var choix string
	fmt.Scanln(&choix)
	n, err := strconv.Atoi(choix)
	if err != nil || n < 1 || n > len(e.Inventaire) {
		fmt.Println("Choix invalide.")
		return
	}

	objet := e.Inventaire[n-1]
	switch objet {
	case "Potion de soin":
		*e = potions.TakePot(*e)
	case "Potion de poison":
		*e = potions.PoisonPot(*e)
	case "Potion de PV total":
		*e = potions.FullHealthPot(*e)
	case "Potion de guérison du poison":
		fmt.Println("Cette potion est utilisable uniquement pendant un combat.")
	case "bandeau frontale ninja", "Manteau Akatsuki", "Bottes de Shinobi":
		forgeron.Equiper(e, objet)
	case "upgrade1", "upgrade2":
		ancienneCapacite := e.CapaciteInventaire
		fmt.Println(UpgradeInv.UpgradeInventorySlot(e))
		if e.CapaciteInventaire > ancienneCapacite {
			tache7101111suite.RemoveInventory(e, objet)
		}
	case "Kunaï":
		spellbook.LearnSpell(e, objet, "Kunaï")
	case "Rasengan":
		spellbook.LearnSpell(e, objet, "Rasengan")
	case "Sharingan":
		spellbook.LearnSpell(e, objet, "Sharingan")
	default:
		fmt.Println("Cet objet est un matériau du forgeron et ne peut pas être utilisé.")
	}

}
