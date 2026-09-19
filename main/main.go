package main


import (
	"Projet_Red/personnage"
	"fmt"
	"Projet_Red/tache8"
)

func main() {
	c1 := personnage.InitCharacter("Kael", "Etudiant", 1, 10, 5, []string{},"chapeau d'aventurier")
	tache8.Isdead(c1)
	fmt.Print(c1.Argent)
	fmt.Print(c1.Equipement)
	tache8.Isdead(c1)
}
