package inventaire

import (
	npc "Projet_Red/NPC"
	"Projet_Red/amelioration"
	"Projet_Red/personnage"
	"Projet_Red/potions"
	"Projet_Red/sorts"
	"fmt"
	"strconv"
	"strings"
)

func AccessInventory(e *personnage.Etudiant) {
	fmt.Println("Inventaire :", strings.Join(e.Inventaire, ", "))
	fmt.Println("Capacité :", len(e.Inventaire), "/", e.CapaciteInventaire)
	fmt.Println("1 - Supprimer un objet de l'inventaire ")
	fmt.Println("2 - Utiliser un objet")
	fmt.Println("0 - Retour")
	for {

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "0":
			return
		case "1":
			if len(e.Inventaire) == 0 {
				fmt.Println("Votre inventaire est vide.")
				continue
			}
			fmt.Println("Que souhaites-tu retirer ? Tape le chiffre correspondant à la place de l'item.")
			for i, r := range e.Inventaire {
				fmt.Println(strconv.Itoa(i+1) + ": " + r)
			}
			fmt.Println("0 - Retour")

			for {
				var choix_sup string
				fmt.Scanln(&choix_sup)
				if choix_sup == "0" {
					return
				}
				n, err := strconv.Atoi(choix_sup)
				if err != nil || n < 1 || n > len(e.Inventaire) {
					fmt.Println("Ce n'est pas au menu, gamin.")
					continue

				} else {

					fmt.Println("Je vais retirer " + e.Inventaire[n-1] + " de l'inventaire")
					npc.RemoveInventory(e, e.Inventaire[n-1])
					fmt.Println(e.Inventaire)
					return
				}

			}

		case "2":
			utiliserObjet(e)

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
	fmt.Println("0 - Retour")

	var choix string
	fmt.Scanln(&choix)
	if choix == "0" {
		return
	}
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
	case "bandeau frontal ninja", "Manteau Akatsuki", "Pantalon des Six Chemins", "Sandales du Shinobi":
		npc.Equiper(e, objet)
	case "upgradeinventoryslot1", "upgradeinventoryslot2":
		ancienneCapacite := e.CapaciteInventaire
		fmt.Println(amelioration.UpgradeInventorySlot(e))
		if e.CapaciteInventaire > ancienneCapacite {
			npc.RemoveInventory(e, objet)
		}
	case "Lot de Kunaï (nouveau sort)", "Kunaï":
		sorts.LearnSpell(e, objet, "Kunaï")
	case "Rasengan", "Rasengan (nouveau sort)":
		sorts.LearnSpell(e, objet, "Rasengan")
	case "Sharingan":
		sorts.LearnSpell(e, objet, "Sharingan")
	default:
		fmt.Println("Cet objet est un matériau du forgeron et ne peut pas être utilisé.")
	}

}
