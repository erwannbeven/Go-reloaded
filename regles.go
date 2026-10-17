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
	ligne = normaliserMarqueurs(ligne)
	mots := strings.Fields(ligne)
	mots = appliquerMarqueurs(mots)
	lignes[i] = strings.Join(mots, " ")
}
	return strings.Join(lignes, "\n")
}

// normaliserMarqueurs supprime les espaces après les virgules dans les balises
func normaliserMarqueurs(s string) string {
	for _, m := range []string{"up", "low", "cap"} {
		prefixe := "(" + m + ","
		for {
			idx := strings.Index(s, prefixe)
			if idx == -1 {
				break
			}
			fin := strings.Index(s[idx:], ")")
			if fin == -1 {
				break
			}
			bloc := s[idx : idx+fin+1]
			blocSansEspaces := strings.ReplaceAll(bloc, " ", "")
			s = strings.Replace(s, bloc, blocSansEspaces, 1)
		}
	}
	return s
}

// appliquerMarqueurs crée les marqueur et determine leur effet sur les mots les précedants
func appliquerMarqueurs(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		if strings.HasPrefix(mot, "(up,") || strings.HasPrefix(mot, "(low,") || strings.HasPrefix(mot, "(cap,") {
			if len(resultat) == 0 {
				continue
			}
			debut := strings.Index(mot, ",")
			if debut != -1 && strings.HasSuffix(mot, ")") {
				strNombre := strings.TrimSpace(mot[debut+1 : len(mot)-1])
				n, err := strconv.Atoi(strNombre)
				if err == nil && n > 0 {
					action := mot[1:3]

					for i := 0; i < n; i++ {
						dernier := len(resultat) - 1 - i
						if dernier < 0 {
							break
						}

						if action == "up" { // Ce n'est pas une erreur ne pas toucher (go n'est pas content de la manière dont est écrit le code)
							resultat[dernier] = strings.ToUpper(resultat[dernier])
						} else if action == "lo" {
							resultat[dernier] = strings.ToLower(resultat[dernier])
						} else if action == "ca" {
							resultat[dernier] = capitaliser(resultat[dernier])
						}
					}
				}
			}
			continue
		}

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

		case "(bin)":
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
        default:
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