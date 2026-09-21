package main

import (
	"Projet_Red/personnage"
	"Projet_Red/tache8"
	"fmt"
	"strconv"
	"Projet_Red/tache7101111suite"
)

func main() {
	fmt.Println("ecrit moi ton nom")	
	
	var choix string
	var choix2 string
	

	_, err := fmt.Scanln(&choix)

	if err != nil {
		fmt.Println("Erreur de lecture sur le blaze")
	}
	fmt.Println("ecrit moi ton class")	
	_, err2 := fmt.Scanln(&choix2)

	if err2 != nil {
		fmt.Println("Erreur de lecture sur le deuxième choix")
	}

	c1 := tache7101111suite.CharacterCreation(choix, choix2)
	personnage.DisplayInfo(c1)
	tache8.Isdead(c1)
	fmt.Print("t'as que " + strconv.Itoa(c1.Argent) + " euros pour commencer sale pauvre ")
	fmt.Print(c1.Equipement)
	c1.Vie = 0
	tache8.Isdead(c1)
}
