package main

import (
	"regexp"
    "strconv"
	"strings"
	"unicode"
)

// transformer transforme le texte trouver dans le fichier d'entrée en le découpant pour ensuite le recoler en fesant cetraine modification en se servant des autres fonctions
func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")
		
		ligne = corrigerEspacesMarqueurs(ligne)

		mots := strings.Fields(ligne)
		mots = appliquerAAn(mots)
		mots = appliquerMarqueurs(mots)
		
	// Assemblage et corrections finales de ponctuation et guillemets
		ligneTraitee := strings.Join(mots, " ")
		ligneTraitee = corrigerPonctuation(ligneTraitee)
		lignes[i] = corrigerGuillemets(ligneTraitee)
	}
	return strings.Join(lignes, "\n")
}

// corrigerEspacesMarqueurs nettoie proprement sans boucle infinie
func corrigerEspacesMarqueurs(s string) string {
	// On cherche un motif du type (up, 1), (low,  2), etc., avec des espaces optionnels après la virgule
	re := regexp.MustCompile(`\((up|low|cap),\s*([0-9]+)\)`)
	// On remplace proprement par le format compact sans espace (ex: "(up,1)")
	return re.ReplaceAllString(s, "($1,$2)")
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

// appliquerAAn transforme le mot a devient an si le mot suivant commence par une voyelle (a, e, i, o, u) ou un h.
func appliquerAAn(mots []string) []string {
	for i := 0; i < len(mots); i++ {
		if mots[i] == "a" || mots[i] == "A" {
			// On cherche le prochain VRAI mot en sautant les balises (up, low, cap, hex, bin, etc.)
			j := i + 1
			for j < len(mots) && (strings.HasPrefix(mots[j], "(") && strings.HasSuffix(mots[j], ")") || strings.Contains(mots[j], ",")) {
				j++
			}

			// Si on a trouvé un mot après les balises
			if j < len(mots) {
				prochainMot := mots[j]
				if len(prochainMot) > 0 {
					premiereLettre := rune(prochainMot[0])
					lettreInf := unicode.ToLower(premiereLettre)

					// Vérifie si le vrai mot suivant commence par une voyelle ou un 'h'[cite: 2]
					if lettreInf == 'a' || lettreInf == 'e' || lettreInf == 'i' || lettreInf == 'o' || lettreInf == 'u' || lettreInf == 'h' {
						if mots[i] == "A" {
							mots[i] = "An"
						} else {
							mots[i] = "an"
						}
					}
				}
			}
		}
	}
	return mots
}

// corrigerPonctuation corride les ponctuation simple (., ,, !, ?, :, ;) : Collée au mot précédent, suivie d'un espace. Et les groupes de ponctuation (..., !?) : Restent contigus et collés au mot précédent.
func corrigerPonctuation(ligne string) string {
	// 1. S'assurer qu'il y a un espace après un signe de ponctuation (s'il n'y en a pas déjà un)
	// Exemple: "bonjour,tout" -> "bonjour, tout"
	reEspaceApres := regexp.MustCompile(`([.,!?;:])([^\s"'])`)
	ligne = reEspaceApres.ReplaceAllString(ligne, `$1 $2`)

	// 2. Coller le signe de ponctuation au mot qui le précède (en supprimant l'espace avant)
	// Exemple: "bonjour , tout" -> "bonjour, tout"
	reEspaceAvant := regexp.MustCompile(`\s+([.,!?;:])`)
	ligne = reEspaceAvant.ReplaceAllString(ligne, `$1`)

	// 3. Nettoyer les espaces à l'intérieur des groupes de ponctuation (ex: ". . ." -> "...")
	// On peut cibler des cas fréquents comme les points de suspension espacés
	reGroupes := regexp.MustCompile(`\.\s*\.\s*\.`)
	ligne = reGroupes.ReplaceAllString(ligne, "...")

	return ligne
}

// corrigerGuillemets Le texte situé à l'intérieur est collé contre les guillemets sans espaces superflus.
func corrigerGuillemets(ligne string) string {
	// Cette regex cherche un guillemet, des espaces éventuels, du texte, des espaces, et le guillemet fermant
	// Et elle nettoie les espaces intérieurs en gardant les guillemets collés au texte
	re := regexp.MustCompile(`'\s+([^']+?)\s+'`)
	ligne = re.ReplaceAllString(ligne, "'$1'")

	// On gère aussi le cas où il y a un espace d'un seul côté (ex: "' awesome'")
	reGauche := regexp.MustCompile(`'\s+([^']+)`)
	ligne = reGauche.ReplaceAllString(ligne, "'$1")

	reDroite := regexp.MustCompile(`([^']+)\s+'`)
	ligne = reDroite.ReplaceAllString(ligne, "$1'")

	return ligne
}