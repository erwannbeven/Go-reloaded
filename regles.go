package main

import "strings"

func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")
		mots := strings.Fields(ligne)
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}