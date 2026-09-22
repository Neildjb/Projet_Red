package tache14

import "fmt"

func Marchand(){
	fmt.Println("Bienvenue chez le marchand de Konoha, les temps sont durs mais voilà ce que je te propose")
	
}

type Item struct{
	Nom 	string
	Prix	 int
}
var boutique = []Item{
	{Nom: "Kunaï",Prix: 5},
	{Nom: "Shuriken",Prix: 3},
	{Nom: "Rasengan",Prix: 10},
	{Nom: "potion de vie",Prix: 0},
}