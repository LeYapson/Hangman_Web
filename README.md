# Hangman Web - Documentation

Ce projet Hangman Web a été entretenu par Théau YAPI et Yanisse ISMAILI

## Sommaire

- [I. A propos du projet](#i-a-propos-du-projet)
- [II. Utilisation de Hangman Web](#ii-utilisation-de-hangman-web)
- [III. Réalisation des consignes & bonus](#iii-réalisation-des-consignes-bonus)
    - [A. Consignes requises](#a-consignes-requises)
    - [B. Partie 1 : level](#b-Partie-1-level)
    - [C. Partie 2 : gamifier-l'interface](#c-Partie-2-gamifier-l'interface)
    - [D. Partie 3 : bring-to-death](#d-Partie-3-bring-to-death)
    - [E. Partie 4 : Bonus](#e-Partie-4-Bonus)
- [IV. Crédits](#iv-crédits)


## I. A propos du projet

Le projet Hangman Web est un projet dont le but est de recréer le jeu du pendu sur une interface web en utilisant un langage imposé : le [Golang](https://go.dev/).

Nous avons choisi de suivre la charte graphique de Ytrack, la plateforme de gestion de projets d'Ynov, mêlée à une idée de terminal.

## II. Utilisation de Hangman Web

Pour lancer le projet, vous avez seulement besoin de cloner ou télécharger ce repository, ouvrir un terminal et vous diriger dans le dossier contenant le projet.
Ici, vous n'avez plus qu'à exécuter la commande `go run main.go`, et utiliser votre navigateur pour vous rendre sur le lien suivant : [http://localhost:8080](http://localhost:8080).

## III. Réalisation des consignes & bonus

### A. Consignes requises

La consigne principale était de créer une interface web proposant de jouer au jeu du pendu avec le langage de programmation Golang.

---

### B. Partie 1 : level

La partie "level" consiste dans l'ajout d'un choix de difficulté avant de commencer une partie : nous avons implémenter une page dédiée au début de la partie qui permet à l'utilisateur de choisir entre trois difficultés (Classic et ! OHIO !).

---

### C. Partie 2 : gamifier l'interface

cette partie consiste à faire en sorte que l'interface corresponde mieux à un jeu vidéo.

Nous avons donc décidés de créer de nouvelles interfaces telles que : une page de connexion (pour choisir son pseudo), un écran de fin de partie (gagné / perdu), un bouton pour jouer de nouveau, une liste des lettes déjà utilisées durant la partie, ...

---

### D. Partie 3 : bring-to-death

La Partie "bring-to-death" consiste en la création d'un pendu qui s'actualise au fur-et-à-mesure du jeu, afin de rendre plus visuel le nombre d'essais restants.

Chaque fois que l'on émet une suggestion mauvaise (mauvaise lettre / mauvais mot), le pendu apparaît de plus en plus.

---

### E. Partie 4 : Bonus

Cette partie consiste en l'ajout de quelques easter-eggs partout sur le projet. tous les easters-eggs présent sont des boutons visibles sur les pages. à vous de découvrir vers où ils mènent.

---

## IV. Crédits

Le projet Hangman Web a été conçu par Théau YAPI et Yanisse ISMAILI (avec l'aide précieuse de notre mentor Abdelhadi Hasnaoui).
