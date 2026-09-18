package main


import (
	"Projet_Red/personnage"
	"fmt"
	"Projet_Red/tache8"
)

func main() {
	c1 := personnage.InitCharacter("Kael", "Etudiant", 1, 10, 5, []string{})
	tache8.isdead()
	fmt.Print(c1)
}
