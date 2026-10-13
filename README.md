# textfix
ce programme sert a modifier un texte en utilisent des marqueurs
Exemple : (up), (low) et (cap)

## Installation
Clonez le projet ou téléchargez-le.
Placez-vous dans le dossier racine du projet :
   ```bash
   cd textfix
   ```

## Utilisation
pour démarer le programme il suffit d'entrée dans le terminal bash :
```bash
go run . entree.txt sortie.txt
```
resultat:

entrée.txt:
bonjour (cap) il (up) fait   nuit

sortie.txt:
Bonjour IL fait nuit

## Les règles

| Marqueur | Effet | Exemple | Résultat |
|---|---|---|---|
| `(up)` | met le mot precedant en entier en majuscule| bonjour (up)| BONJOUR|
| `(low)` | met le mot precedant en entier en minuscule| BONJOUR (low)| bonjour|
| `(cap)` | met la premiere lettre du mot precedant en majuscule et le reste en minuscule| boNJOur (cap)| Bonjour|

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
regles.go : Transformer le texte

## Limites
Le programme ne gère pas les marqueurs collés directement au mot sans espace 
exemple : mot(up) ne sera pas transformé correctement

## Auteur
Erwann BEVEN