package main

import (
	"fmt"
	"os"
)

//lireFichier lie le fichier dont on a donner le nom et si il y a un probleme le signal
func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("impossible de lire le fichier %s", chemin)
	}
	return string(contenu), nil
}

//ecrireFichier ecrit un fichier de sortie du texte dont on a donner le nom au moment de la commande (comme pour lireFichier) et crée le fichier
func ecrireFichier(chemin string, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("impossible d'écrire le fichier %s", chemin)
	}
	return nil
}