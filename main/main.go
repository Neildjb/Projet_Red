package main

import (
	"Projet_Red/personnage"
	"fmt"
)

func main() {
	c1 := personnage.InitCharacter("Kael", "Etudiant", 1, 10, 5, []string{})
	fmt.Print(c1)
}
