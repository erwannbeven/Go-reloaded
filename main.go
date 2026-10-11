package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Erreur : il faut deux noms de fichiers.")
		fmt.Fprintln(os.Stderr, "Utilisation : textfix <fichier_entree> <fichier_sortie>")
		os.Exit(1)
	}
	if os.Args[1] == os.Args[2]{
		fmt.Fprintln(os.Stderr, "Erreur : les noms des fichiers d'entrée et de sortie doivent être différents.")
		os.Exit(1)
	}
	contenu, err := lireFichier(os.Args[1])
	if err != nil {
    fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
	}
	contenu = transformer(contenu)
	err = ecrireFichier(os.Args[2], contenu)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Terminé :", os.Args[2], "a été créé.")
}