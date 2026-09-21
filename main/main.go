package main

import (
	"Projet_Red/personnage"
	"Projet_Red/tache8"
	"fmt"
	"strconv"
)

func main() {
	c1 := personnage.InitCharacter("Kael", "Etudiant", 1, 10, 5, []string{}, personnage.Equipment{})
	tache8.Isdead(c1)
	fmt.Print("t'as que " + strconv.Itoa(c1.Argent) + " euros pour commencer sale pauvre ")
	fmt.Print(c1.Equipement)
	c1.Vie = 0
	tache8.Isdead(c1)
}
