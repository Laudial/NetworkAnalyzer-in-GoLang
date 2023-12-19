package main

import (
	"fmt"

	"NetworkAnalyzer/controllers"
	"NetworkAnalyzer/utils"
)

func main() {

	for {
		// Nettoyage de l'écran
		utils.ClearScreen()

		// Affichage du menu
		fmt.Println("1. Résolution DNS")
		fmt.Println("2. Scan de ports")
		fmt.Println("0. Quitter")

		// Demande du choix de l'utilisateur
		var choice int
		fmt.Print("Choisissez une option : ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Erreur de saisie :", err)
			return
		}

		// Traitement du choix de l'utilisateur
		switch choice {
		case 0:
			utils.ClearScreen()
			fmt.Println("Au revoir !")
			return
		case 1:
			utils.ClearScreen()
			controllers.DnsResolver()
		case 2:
			utils.ClearScreen()
			controllers.PortScanner()
		default:
			utils.ClearScreen()
			fmt.Println("Option non valide. Veuillez choisir à nouveau.")
		}
		
	}
}