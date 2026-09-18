package main

import (
	"Projet_Red/personnage"
	"fmt"
)

func Isdead(e personnage.Etudiant) {
	if e.Vie == 0 {
		fmt.Println("tu es mort mais tout le monde a le droit à une seconde chance")
		e.Vie = 50
		fmt.Println("tu es  revivant")
	}
}


