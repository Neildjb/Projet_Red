package main

import (
	"Projet_Red/etat"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Jeu struct {
	background      *ebiten.Image
	spritesJoueurs  map[string]*ebiten.Image
	spritesMonstres map[string]*ebiten.Image
}

func chargerImage(chemin string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(chemin)
	if err != nil {
		log.Println("image introuvable :", chemin, "-", err)
		return nil
	}
	return img
}

func (j *Jeu) chargerSprites() {
	// Pas de fond pour l'instant. Une fois que tu as une image, mets-la dans
	// assets/background/ et décommente la ligne suivante avec le bon nom :
	// j.background = chargerImage("assets/background/fond.png")

	nomsFichiersJoueurs := map[string]string{
		"Kage":         "background/Naruto_demon.png",
		"Jonin":        "background/sasuke.png",
		"Genin":        "background/RockLee.png",
		"Ninja":        "background/Naruto_Prime.png",
		"Naruto_Prime": "background/Naruto_Prime.png",
		"Admin4416":    "background/Naruto_Prime.png",
	}
	j.spritesJoueurs = make(map[string]*ebiten.Image)
	for classe, fichier := range nomsFichiersJoueurs {
		j.spritesJoueurs[classe] = chargerImage("assets/" + fichier)
	}

	nomsFichiersMonstres := map[string]string{
		"Maxime_boss|Naruto_Prime": "Maxime_boss.png",
		"Mathias|Naruto_Prime":     "Mathias.png",
		"Andy|Naruto_Prime":        "Andy.png",
		"Neil|Naruto_Prime":        "Neil.png",

		"Neil|Admin4416":        "Neil.png",
		"Andy|Admin4416":        "Andy.png",
		"Mathias|Admin4416":     "Mathias.png",
		"Maxime_boss|Admin4416": "Maxime_boss.png",

		"Pain|Kage":   "Pain.png",
		"Itachi|Kage": "Itachi.png",
		"Obito|Kage":  "Obito.png",
		"Madara|Kage": "Madara.png",

		"Deidara|Jonin":    "Deidara.png",
		"Kisame|Jonin":     "Kisame.png",
		"Sasori|Jonin":     "Sasori.png",
		"Orochimaru|Jonin": "Orochimaru.png",

		"Asuma|Genin":   "Asuma.png",
		"Gaara|Genin":   "Gaara.png",
		"Shino|Genin":   "Shino.png",
		"Kakashi|Genin": "Kakshi.png",

		"Maxime":                "maxime.png",
		"Ninja déserteur|Ninja": "NINJA_deserteur.png",
		"Sai|Ninja":             "SAI.png",
		"Yamato|Ninja":          "Yamato.png",
		"Danzô|Ninja":           "Danzo.png",
	}
	j.spritesMonstres = make(map[string]*ebiten.Image)
	for nom, fichier := range nomsFichiersMonstres {
		j.spritesMonstres[nom] = chargerImage("assets/monstres/" + fichier)
	}
}
func (j *Jeu) spriteMonstre(nom string, classeJoueur string) *ebiten.Image {
	if sprite, ok := j.spritesMonstres[nom+"|"+classeJoueur]; ok && sprite != nil {
		return sprite
	}
	return j.spritesMonstres[nom]
}
func (j *Jeu) Update() error {
	return nil
}

func (j *Jeu) Draw(ecran *ebiten.Image) {
	if j.background != nil {
		ecran.DrawImage(j.background, nil)
	}

	if etat.JoueurActuel == nil {
		ebitenutil.DebugPrint(ecran, "Création du personnage en cours dans le terminal...")
		return
	}

	if sprite := j.spritesJoueurs[etat.JoueurActuel.Classe]; sprite != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(0.3, 0.3)
		opts.GeoM.Translate(80, 150)
		ecran.DrawImage(sprite, opts)
	}
	dessinerBarreDeVie(ecran, 50, 50, etat.JoueurActuel.Vie, etat.JoueurActuel.MaxVie)
	ebitenutil.DebugPrintAt(ecran, fmt.Sprintf("%s (%s)", etat.JoueurActuel.Nom, etat.JoueurActuel.Classe), 50, 30)

	if etat.MonstreActuel != nil {
		sprite := j.spriteMonstre(etat.MonstreActuel.Nom, etat.JoueurActuel.Classe)
		if sprite != nil {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Scale(0.3, 0.3)
			opts.GeoM.Translate(400, 150)
			ecran.DrawImage(sprite, opts)
		}
		dessinerBarreDeVie(ecran, 400, 50, etat.MonstreActuel.Vie, etat.MonstreActuel.Max_vie)
		ebitenutil.DebugPrintAt(ecran, etat.MonstreActuel.Nom, 400, 30)
	}
}

func (j *Jeu) Layout(largeurEcran, hauteurEcran int) (int, int) {
	return 640, 480
}

func dessinerBarreDeVie(ecran *ebiten.Image, x, y, vie, vieMax int) {
	if vieMax <= 0 {
		return
	}
	largeurMax := 200
	largeurActuelle := int(float64(vie) / float64(vieMax) * float64(largeurMax))
	if largeurActuelle < 0 {
		largeurActuelle = 0
	}

	ebitenutil.DrawRect(ecran, float64(x), float64(y), float64(largeurMax), 20, color.RGBA{80, 80, 80, 255})
	ebitenutil.DrawRect(ecran, float64(x), float64(y), float64(largeurActuelle), 20, color.RGBA{0, 200, 0, 255})
}
