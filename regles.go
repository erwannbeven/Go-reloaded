package main

import (
	"strings"
	"unicode"
)

func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
	ligne = strings.TrimRight(ligne, "\r")
	mots := strings.Fields(ligne)
	mots = appliquerMarqueurs(mots)
	lignes[i] = strings.Join(mots, " ")
}
	return strings.Join(lignes, "\n")
}

func appliquerMarqueurs(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
        switch mot {
        case "(up)", "(low)", "(cap)":
            // Si c'est un marqueur, on vérifie qu'il y a un mot avant
            if len(resultat) == 0 {
                continue
            }
            dernier := len(resultat) - 1
            if mot == "(up)" {
                resultat[dernier] = strings.ToUpper(resultat[dernier])
            } else if mot == "(low)" {
                resultat[dernier] = strings.ToLower(resultat[dernier])
            } else if mot == "(cap)" {
                resultat[dernier] = capitaliser(resultat[dernier])
            }
        default:
            // Si c'est un mot normal (ou un marqueur inconnu comme (truc)), on l'ajoute
            resultat = append(resultat, mot)
        }
    }

    return resultat
}

func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}