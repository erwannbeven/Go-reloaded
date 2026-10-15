package main

import (
    "strconv"
	"strings"
	"unicode"
)

// transformer transforme le texte trouver dans le fichier d'entrée en le découpant pour ensuite le recoler en fesant cetraine modification en se servant des autres fonctions
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

// appliquerMarqueurs crée les marqueur et determine leur effet sur les mots les précedants
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
            if mot == "(up)" { // Ce n'est pas une erreur ne pas toucher (go n'est pas content de la manière dont est écrit le code)
                resultat[dernier] = strings.ToUpper(resultat[dernier])
            } else if mot == "(low)" {
                resultat[dernier] = strings.ToLower(resultat[dernier])
            } else if mot == "(cap)" {
                resultat[dernier] = capitaliser(resultat[dernier])
            }
            case "(hex)":
			// On vérifie qu'il y a un mot avant
			if len(resultat) == 0 {
				continue
			}
			dernier := len(resultat) - 1
			// Conversion de la base 16 (hexadécimal) vers un entier (int64)
			decimalVal, err := strconv.ParseInt(resultat[dernier], 16, 64)
			if err == nil {
				// Si c'est un nombre hex valide, on le transforme en texte décimal (base 10)
				resultat[dernier] = strconv.FormatInt(decimalVal, 10)
			}
			// Note : On n'ajoute PAS "(hex)" dans resultat, il disparaît du texte final.

		case "(bin)":
			// On vérifie qu'il y a un mot avant
			if len(resultat) == 0 {
				continue
			}
			dernier := len(resultat) - 1
			// Conversion de la base 2 (binaire) vers un entier (int64)
			decimalVal, err := strconv.ParseInt(resultat[dernier], 2, 64)
			if err == nil {
				// Si c'est un nombre binaire valide, on le transforme en texte décimal (base 10)
				resultat[dernier] = strconv.FormatInt(decimalVal, 10)
			}
			// Note : On n'ajoute PAS "(bin)" dans resultat, il disparaît également.
        default:
            // Si c'est un mot normal (ou un marqueur inconnu comme (truc)), on l'ajoute
            resultat = append(resultat, mot)
        }
    }

    return resultat
}

// capitaliser sépare le mot devant (cap) en caractere pour tous les mettres en minuscule et ensuite mettre la premiere letre en majuscule
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}