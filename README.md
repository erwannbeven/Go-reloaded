# textfix

`textfix` est un outil en Go conçu pour nettoyer et reformatter automatiquement un texte à l'aide de marqueurs spécifiques et de règles de ponctuation/grammaire.
Exemple : (up), (low) et (cap)
---

## Installation

Clonez le projet ou téléchargez-le, puis placez-vous dans le dossier racine du projet :

```bash
cd textfix
```
## Utilisation
Pour démarrer le programme, spécifiez le fichier d'entrée et le fichier de sortie dans le terminal :
```bash
go run . entree.txt sortie.txt
```
Exemple :

entrée.txt:
It was 10 (bin) times better than 42 (hex) , right?
A (up) amazing adventure is starting here...
I can't believe it, she said ' yes ' !
This is so exciting (up, 2) and a amazing day.

sortie.txt:
It was 2 times better than 66, right?
AN amazing adventure is starting here...
I can't believe it, she said 'yes'!
This is SO EXCITING and an amazing day.

## Les règles
TOUJOUR mettre un espace après un marqueur si une virgule le suit
exemple : mauvais
entrée.txt:It was 10 (bin) times better than 42 (hex), right?
sortie.txt:It was 2 times better than 42 (hex), right?
exemple : correcte
entrée.txt:It was 10 (bin) times better than 42 (hex) , right?
sortie.txt:It was 2 times better than 66, right?


| Marqueur | Effet | Exemple | Résultat |
|---|---|---|---|
1. Marqueurs de transformation de casse
| `(up)` | met le mot precedant en entier en majuscule| bonjour (up)| BONJOUR|
|`(up, N)`|Met les N mots précédents en majuscules|so exciting (up, 2)|SO EXCITING|
| `(low)` | met le mot precedant en entier en minuscule| BONJOUR (low)| bonjour|
|`(low, N)`|Met les N mots précédents en minuscules|TURN DOWN (low, 2)|turn down|
| `(cap)` | met la premiere lettre du mot precedant en majuscule et le reste en minuscule| boNJOur (cap)| Bonjour|
|`(cap, N)`|Capitalise les N mots précédents|salut tout le monde (cap, 2)|salut tout Le Monde|
2. Conversions numériques
|`(hex)`|Convertit un nombre hexadécimal en décimal|42 (hex)|66|
|`(bin)`|Convertit un nombre binaire en décimal|10 (bin)|2|
3. Grammaire (a / an)
Le mot a devient an si le mot suivant commence par une voyelle (a, e, i, o, u) ou un h.La règle fonctionne même si des marqueurs (comme (up)) s'interposent entre a et le mot suivant.Exemple : a (up) amazing $\rightarrow$ AN amazing
4. Ponctuation et Guillemets
Ponctuation simple (., ,, !, ?, :, ;) : Collée au mot précédent, suivie d'un espace.

Groupes de ponctuation (..., !?) : Restent contigus et collés au mot précédent.

Guillemets simples (') : Le texte situé à l'intérieur est collé contre les guillemets sans espaces superflus.

## Les messages d'erreur
Dans le cas ou vous n'auriez pas mis de nom de fichier il affichera :Erreur : il faut deux noms de fichiers.
Dans le cas ou vous auriez mis le meme nom de fichier a l'entre et a la sortie il affichera : Erreur : les noms des fichiers d'entrée et de sortie doivent être différents.
Dans le cas ou le programme a un probleme avec le fichier d'entrée il affichera l'erreur du systeme
Dans la meme logique que pour le fichier d'entrée si il a un probleme avec le fichier de sortie il affichera l'erreur du systeme

## Les tests
Pour lancer un teste de faire comme pour lancer avec le texte dans entrée.txt mais a la place de entrée on choisie un des txt dans le sous-fichier tests comme 01-majuscules.txt et de le mettre a la place de entrée dans le terminal tout en gardant sortie.txt 
exemple : go run . tests/01-majuscules.txt sortie.txt

Et pour comparer toujour dans le terminal il suffit de taper diff sortie.txt suivi du fichier txt des resultat attendu du systeme en lien avec le fichier teste comme par exemple : 01-majuscules.txt et 01-majuscules.attendu.txt
donc la ligne final resenble a ça : diff sortie.txt tests/01-majuscules.attendu.txt

## Organisation du code
main.go : Vérifier les arguments, enchaîner les étapes, afficher les messages
fichiers.go : Lire et écrire les fichiers
regles.go : Logique de transformation du texte (marqueurs, conversions de base, règle a/an, ponctuation et guillemets).

## Limites
Le programme nécessite des espaces autour des marqueurs pour les détecter correctement (ex: mot (up) et non mot(up)).

Les apostrophes de contraction (ex: don't, can't) ne sont pas confondues avec les guillemets de citation.
Mais le guillemts de citation seul sont laissés tel quel.

## Auteur
Erwann BEVEN