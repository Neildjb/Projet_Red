# Shinobi Genesis

Jeu de combat au tour par tour réalisé en Go dans le cadre du Projet RED.

## Lancer le jeu

Depuis la racine du projet :

```bash
go run .
```

Lorsque tu lances le jeu, une fenêtre graphique contenant les images des personnages et des monstres s'ouvre en plus du terminal. Pour voir les deux en même temps, partage ton écran en deux : garde le terminal d'un côté et la fenêtre graphique de l'autre.

Les questions et les menus se remplissent dans le terminal : écris le numéro ou le texte demandé, puis appuie sur `Entrée`.

Pour arrêter le jeu, appuie sur `Ctrl+C` dans le terminal. Le menu principal ne propose pas de bouton pour quitter.

Les informations sont affichées les unes sous les autres. Si elles dépassent la hauteur de la fenêtre, remonte avec la molette de la souris pour relire les choix, les statistiques et les messages du combat.

## Comment jouer

1. Choisis un pseudo et une classe.
2. Consulte le menu principal pour afficher ton personnage, ouvrir l'inventaire, acheter des objets ou combattre.
3. Pendant un combat, choisis une attaque ou ouvre ton inventaire. Le choix `0` permet uniquement de revenir au menu de combat depuis la liste des sorts ou des objets ; il ne permet pas de quitter le combat.
4. Termine le tutoriel contre Maxime : il rapporte 20 pièces la première fois, puis 2 pièces lors des fois suivantes.
5. Gagne des combats pour obtenir de l'argent et de l'expérience.
6. Achète des sorts, fabrique et équipe des équipements, puis change de classe à l'école.
7. Utilise `0` pour revenir en arrière dans les menus et sous-menus.
8. Consulte les quêtes pour suivre ta progression jusqu'à la fin du jeu.

### Équipements

Le personnage peut porter quatre équipements en même temps : tête, corps, jambes et pieds. Chaque équipement ajoute des PV lorsqu'il est équipé :

- Bandeau frontale ninja : `+15 PV`
- Manteau Akatsuki : `+20 PV`
- Pantalon des Six Chemins : `+25 PV`
- Sandales du Shinobi : `+30 PV`

Les Sandales du Shinobi se fabriquent chez le forgeron avec 4 `queue de demon`.

### Classes

| Classe | Difficulté | PV de départ | Image du personnage | Ennemis combattus |
| --- | --- | ---: | --- | --- |
| Kage | Très facile | 300 | `assets/background/Naruto_demon.png` | Pain, Itachi, Obito, Madara |
| Jonin | Facile | 150 | `assets/background/sasuke.png` | Deidara, Kisame, Sasori, Orochimaru |
| Genin | Moyenne | 100 | `assets/background/RockLee.png` | Asuma, Gaara, Shino, Kakshi |
| Ninja | Très difficile | 50 | `assets/background/Naruto_Prime.png` | Ninja déserteur, Sai, Yamato, Danzo |
| Naruto_Prime | Hardcore | 30 | `assets/background/Naruto_Prime.png` | Neil, Andy, Mathias, Maxime_boss |
| Admin4416 | Spéciale | 10 000 | `assets/background/Naruto_Prime.png` | Neil, Andy, Mathias, Maxime_boss |

Le mode `Naruto_Prime` est un mode hardcore : une seule mort est décisive et la partie s'arrête définitivement. C'est ici que le jeu devient vraiment intéressant. Bonne chance pour terminer le jeu dans cet état !

## Arborescence

```text
Projet_Red/
├── amelioration/
│   └── inventaire.go    # UpgradeInventorySlot
├── combat/
│   └── combat.go        # Combat, TrainingFight, displayCombatStatus,
│                        # collectMonsterDrop, collectMonsterMoney,
│                        # collectCombatExperience, inventoryTurn,
│                        # characterTurn
├── etat/
│   └── etat.go          # JoueurActuel, MonstreActuel
├── gestionmort/
│   └── mort.go          # Isdead
├── inventaire/
│   └── inventaire.go    # AccessInventory, utiliserObjet
├── jeu/
│   └── creation_personnage.go
│                        # Capitalize, CharacterCreation
├── Menu/
│   └── Menu.go          # Menu, quetesFinJeu, affiche_credit,
│                        # afficherQuete, nombreDeSorts,
│                        # nombreDEquipements, boolVersNombre,
│                        # possedeObjet, ecole, arenaMenu
├── NPC/
│   ├── forgeron.go      # Forgeron, Equiper, retirerUn, compter
│   └── marchand.go      # Marchand, AddInventory, RemoveInventory,
│                        # prixActuel
├── personnage/
│   └── personnage.go    # InitCharacter, DisplayInfo, AjouterItem,
│                        # bonusEquipement,
│                        # Init_Maxime, Init_Ninjas_déserteurs,
│                        # Init_Golems_de_chakra, Init_Demon_a_queue,
│                        # Init_Madara
├── potions/
│   └── potions.go       # TakePot, PoisonPot, FullHealthPot
├── sorts/
│   └── sorts.go         # LearnSpell, containsSpell, SpellDamage
├── assets/              # Images des personnages et des monstres
├── jeu_graphique.go     # chargerImage, chargerSprites, spriteMonstre,
│                        # Update, Draw, Layout, dessinerBarreDeVie
├── main.go              # main, lancerJeuTerminal
├── go.mod               # Configuration du module Go
└── README.md            # Documentation du projet
```

## Vérifier le projet

Depuis la racine :

```bash
go test ./...
```

## Crédits

Projet réalisé au Campus Ynov Bordeaux, Bachelor 1, 2026.

### Équipe de développement

- Mathias Fontagne : développement, gameplay et conception
- Neil Djebali : développement, gameplay et conception
- Andy Abbas : développement, gameplay et conception

### Mentors

Merci à Maxime et Sarha pour leur accompagnement, leurs conseils et leur aide tout au long du projet.

Merci d'avoir joué !

© 2026 — Projet RED • Ynov Bordeaux