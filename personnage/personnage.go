package personnage

type Etudiant struct{
	nom string
	classe string
	niveau int 
	max_vie int
	vie int
	inventaire []string
}

func initCharacter(nom string,classe string ,niveau int,max_vie int,vie int,inventaire []string) struct {
	return Etudiant{nom,classe,niveau,max_vie,vie,inventaire}
}


