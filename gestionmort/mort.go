package gestionmort

import (
	"Projet_Red/personnage"
	"fmt"
)

func Isdead(e *personnage.Etudiant) bool {
	if e.Vie != 0 {
		return false
	}
	if e.Hardcore {
		e.GameOver = true
		fmt.Println("Mode hardcore : une seule mort est décisive. La partie est terminée.")
		return true
	}
	fmt.Println("Tu es mort, mais tu as droit à une seconde chance.")
	e.Vie = e.MaxVie
	fmt.Println("Tu revis avec tous tes PV.")
	return true
}
