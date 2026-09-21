package personnage

import "fmt"
import (
	"Projet_Red/personnage"
)
func takePot(perso personnage.Etudiant) personnage.Etudiant {
	trouve := false
	position := 0

	for i, item := range perso.Inventaire {
		if item == "Potion de soin" {
			trouve = true
			position = i
		}
	}

	if trouve == false {
		fmt.Println("Pas de potion de soin dans l'inventaire !")
		return perso
	}

	inv := perso.Inventaire[:len(perso.Inventaire)-1]
	j := 0

	for i, item := range perso.Inventaire {
		if i != position {
			inv[j] = item
			j = j + 1
		}
	}

	perso.Inventaire = inv

	perso.Vie = perso.Vie + 50
	if perso.Vie > perso.MaxVie {
		perso.Vie = perso.MaxVie
	}

	fmt.Print("a utilisé une potion de soin. Ta vie actuelle est de : ", perso.Vie, perso.MaxVie) 

	return perso
}

func poisonPot(perso Etudiant) Etudiant {
	trouve := false
	position := 0

	for i, item := range perso.Inventaire {
		if item == "Potion de poison" {
			trouve = true
			position = i
		}
	}

	if trouve == false {
		fmt.Print("Pas de potion de poison dans l'inventaire !")
		return perso
	}

	inv := perso.Inventaire[:len(perso.Inventaire)-1]
	j := 0

	for i, item := range perso.Inventaire {
		if i != position {
			inv[j] = item
			j = j + 1
		}
	}

	perso.Inventaire = inv

	perso.Vie = perso.Vie - 20
	if perso.Vie < 0 {
		perso.Vie = 0
	}

	fmt.Print("a bu une potion de poison ! Ta vie actuelle est de : ", perso.Vie, perso.MaxVie)           

	return perso
}
