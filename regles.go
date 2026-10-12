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

func appliquerMarqueurs(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		if mot == "(up)" {
			dernier := len(resultat) - 1
			resultat[dernier] = strings.ToUpper(resultat[dernier])
		} else {
			resultat = append(resultat, mot)
		}
	}

	return resultat
}