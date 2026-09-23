# Shinobi Genesis

Jeu de combat au tour par tour développé en Go dans le cadre du Projet RED.

## Démarrage

Depuis la racine du projet :

```bash
go run .
```

Le jeu utilise deux affichages :

- le terminal pour les menus, les choix et les combats ;
- une fenêtre Ebiten pour afficher les personnages, les monstres et les barres de vie.

Saisis toujours le choix demandé puis appuie sur `Entrée`.

## Dépannage du terminal

Si le jeu semble bloqué après une commande ou si un menu ne répond plus :

1. saisis `0` dans le terminal ;
2. appuie sur `Entrée` ;
3. répète l'opération plusieurs fois si nécessaire, jusqu'à ce que le menu se débloque.

Cette manipulation permet de vider les choix restés dans l'entrée du terminal. Pour arrêter immédiatement le programme, utilise `Ctrl+C`.

## Menu principal

| Choix | Fonction |
| ---: | --- |
| `1` | Afficher les informations du personnage, puis confirmer avec `1` |
| `2` | Ouvrir l'inventaire |
| `3` | Ouvrir le marchand |
| `4` | Ouvrir le forgeron |
| `5` | Lancer le tutoriel de combat |
| `6` | Partir à l'aventure |
| `7` | Accéder à l'école des classes |
| `8` | Consulter les quêtes |
| `9` | Quitter, recommencer ou revenir au menu |

Dans les menus secondaires, `0` sert généralement à revenir en arrière.

## Combat et progression

Les combats se jouent au tour par tour. Le joueur peut attaquer avec ses sorts ou utiliser un objet de combat. Le poison inflige `10 PV` par tour au joueur comme à l'ennemi.

Une victoire rapporte de l'expérience, de l'argent et parfois des objets. Le tutoriel contre Maxime rapporte 20 pièces lors de la première victoire, puis 2 pièces lors des suivantes.

En mode normal, une mort fait perdre le combat, restaure les PV au maximum et conserve l'inventaire, l'équipement, l'argent, l'expérience et la progression. En mode `Naruto_Prime`, une seule mort termine définitivement la partie.

## Classes

| Classe | Difficulté | PV de départ | Image |
| --- | --- | ---: | --- |
| Kage | Très facile | 300 | `Narutodemon` |
| Jonin | Facile | 150 | `sasuke` |
| Genin | Moyenne | 100 | `RockLee` |
| Ninja | Très difficile | 50 | `NarutoPrime` |
| Naruto_Prime | Hardcore | 30 | `NarutoPrime` |
| Admin4416 | Spéciale | 10 000 | `NarutoPrime` |

L'école permet de progresser dans cet ordre : `Naruto_Prime` → `Ninja` → `Genin` → `Jonin` → `Kage`. L'entrée coûte 80 pièces et demande un combat. Les combats de l'école utilisent les lieux affichés dans l'aventure : lac, forêt ou grotte. Si l'expérience minimale n'est pas atteinte, le combat ne démarre pas.

## Marchand et forgeron

Le marchand vend des potions, des sorts et deux améliorations d'inventaire. Les objets uniques suivants ne peuvent être achetés qu'une seule fois :

- `Lot de Kunaï (nouveau sort)` ;
- `Rasengan (nouveau sort)` ;
- `upgradeinventoryslot1` ;
- `upgradeinventoryslot2`.

Chaque équipement du forgeron ne peut également être fabriqué qu'une seule fois. Les équipements sont conservés lorsqu'ils sont équipés ou retirés de l'inventaire.

| Équipement | Bonus de vie |
| --- | ---: |
| Bandeau frontal ninja | `+15 PV` |
| Manteau Akatsuki | `+20 PV` |
| Pantalon des Six Chemins | `+25 PV` |
| Sandales du Shinobi | `+30 PV` |

Ressources disponibles : `Acier`, `Fil d'Akatsuki`, `Tissu déchiré` et `bois millénaire`.

## Quêtes

Le menu des quêtes suit la capacité d'inventaire, l'argent, l'expérience, les sorts, le trophée final et les équipements. Une quête supplémentaire demande d'atteindre `Kage` ou de jouer en mode hardcore. `Admin4416` la valide aussi en interne, sans être mentionné dans le texte affiché.

Une quête secrète existe dans le jeu. Son accès n'est pas indiqué dans le menu. Elle est visible uniquement pour un personnage en mode hardcore et demande de devenir `Kage` en partant de `Naruto_Prime`. Une fois réussie, un message de félicitations s'affiche et le choix `1` permet de retourner au menu.

## Structure du projet

```text
Projet_Red/
├── assets/              # Images du jeu
├── combat/              # Tours, attaques, poison et récompenses
├── etat/                # État du joueur et du monstre affiché
├── gestionmort/         # Gestion des morts normales et hardcore
├── inventaire/          # Utilisation et gestion des objets
├── jeu/                 # Création du personnage
├── Menu/                # Menus, quêtes, école et aventure
├── NPC/                 # Marchand et forgeron
├── personnage/          # Personnage, monstres et ressources
├── potions/             # Effets des potions
├── sorts/               # Sorts et dégâts
├── jeu_graphique.go     # Fenêtre graphique Ebiten
├── main.go              # Lancement du jeu
├── go.mod               # Module Go
└── README.md            # Documentation
```

## Vérification

Depuis la racine du projet :

```bash
go test ./...
```

## Crédits

Projet réalisé au Campus Ynov Bordeaux, Bachelor 1, 2026.

- Mathias Fontagne
- Neil Djebali
- Andy Abbas

Merci à Maxime et Sarha pour leur accompagnement.

© 2026 — Projet RED • Ynov Bordeaux