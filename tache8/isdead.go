package tache8

import (
	"Projet_Red/personnage"
	"fmt"
)

func Isdead(e *personnage.Etudiant) {
	if e.Vie == 0 {
		if e.Hardcore {
			e.GameOver = true
			fmt.Println("Mode hardcore : une seule mort est décisive. La partie est terminée.")
			return
		}
		fmt.Println("tu es mort mais tout le monde a le droit à une seconde chance")
		e.Vie = 50
		fmt.Println("tu es  revivant")
	}
}
