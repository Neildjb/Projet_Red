package main

import "Projet_Red/personnage"
import "fmt"

func accessInventory(e personnage.Etudiant) {
	fmt.Print(e.Inventaire)
}
