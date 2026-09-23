# Shinobi Genesis

Jeu de combat au tour par tour réalisé en Go dans le cadre du Projet RED.

## Lancer le jeu

Depuis la racine du projet :

```text
cd main
go run .
```

Le jeu se joue entièrement dans le terminal. Lorsque le jeu affiche une question ou un menu, écris le numéro ou le texte demandé, puis appuie sur `Entrée`.

Les informations sont affichées les unes sous les autres. Si elles dépassent la hauteur de la fenêtre, remonte avec la molette de la souris pour relire les choix, les statistiques et les messages du combat.

## Comment jouer

1. Choisis un pseudo et une classe.
2. Consulte le menu principal pour afficher ton personnage, ouvrir l'inventaire, acheter des objets ou combattre.
3. Pendant un combat, choisis une attaque ou ouvre ton inventaire.
4. Gagne des combats pour obtenir de l'argent et de l'expérience.
5. Achète des sorts, fabrique et équipe des équipements, puis change de classe à l'école.
6. Consulte les quêtes pour suivre ta progression jusqu'à la fin du jeu.

### Classes

| Classe | Difficulté | PV de départ |
| --- | --- | ---: |
| Kage | Très facile | 300 |
| Jonin | Facile | 150 |
| Genin | Moyenne | 100 |
| Ninja | Très difficile | 50 |
| Naruto_Prime | Hardcore | 30 |

Le mode `Naruto_Prime` est un mode hardcore : une seule mort est décisive et la partie s'arrête définitivement. C'est ici que le jeu devient vraiment intéressant. Bonne chance pour terminer le jeu dans cet état !

## Arborescence

```text
Projet_Red/
├── amelioration/        # Améliorations de l'inventaire
├── combat/              # Combats, attaques et récompenses
├── gestionmort/         # Gestion de la mort et du mode hardcore
├── inventaire/          # Consultation et utilisation de l'inventaire
├── jeu/                 # Création et préparation du personnage
├── main/                # Point d'entrée du jeu
├── Menu/                # Menu principal, arène, école et quêtes
├── NPC/                 # Marchand et forgeron
├── personnage/          # Structures du personnage et des monstres
├── potions/             # Effets des potions
├── sorts/               # Sorts disponibles et dégâts
├── go.mod               # Configuration du module Go
└── README.md            # Documentation du projet
```

## Vérifier le projet

Depuis la racine :

```text
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