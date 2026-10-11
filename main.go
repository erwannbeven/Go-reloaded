package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Erreur : il faut deux noms de fichiers.")
		os.Exit(1)
	}
	contenu, err := lireFichier(os.Args[1])
	if err != nil {
    fmt.Println(err)
	os.Exit(1)
	}
	fmt.Println(contenu)
}