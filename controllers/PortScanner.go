package controllers

import (
	"fmt"
	"log"
	"strconv"
	"net"

	"github.com/Ullaakut/nmap/v2"
)

func NmapPortScanner(target, startPort, endPort string) {

	// Convertir les ports en chaînes de caractères
	startPortInt, err := strconv.Atoi(startPort)
	if err != nil {
		log.Fatalf("Erreur lors de la conversion du port de départ: %v", err)
	}
	startPortStr := strconv.Itoa(startPortInt)

	endPortInt, err := strconv.Atoi(endPort)
	if err != nil {
		log.Fatalf("Erreur lors de la conversion du port de fin: %v", err)
	}
	endPortStr := strconv.Itoa(endPortInt)

	// Créer un nouveau scanner Nmap
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(target),
		nmap.WithPorts(startPortStr, "-", endPortStr),
	)
	if err != nil {
		log.Fatalf("Erreur lors de la création du scanner: %v", err)
	}

	// Lancer le scan
	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("Erreur lors du scan: %v", err)
	}

	// Afficher les avertissements éventuels
	if warnings != nil {
		fmt.Printf("Avertissements: %v\n", warnings)
	}

	// Afficher les résultats du scan
	fmt.Printf("Host %s\n", result.Hosts[0].Addresses[0])
	for _, port := range result.Hosts[0].Ports {
		fmt.Printf("Port %d : %s\n", port.ID, port.State.String())
	}
}

func GoPortScanner(target, startPort, endPort string) {
	
	// Convertir les ports en chaînes de caractères
	startPortInt, err := strconv.Atoi(startPort)
	if err != nil {
		log.Fatalf("Erreur lors de la conversion du port de départ: %v", err)
	}
	startPortStr := strconv.Itoa(startPortInt)

	endPortInt, err := strconv.Atoi(endPort)
	if err != nil {
		log.Fatalf("Erreur lors de la conversion du port de fin: %v", err)
	}
	endPortStr := strconv.Itoa(endPortInt)

	// Créer une connexion TCP
	conn, err := net.Dial("tcp", target + ":" + startPortStr)
	if err != nil {
		log.Fatalf("Erreur lors de la connexion au port %s: %v", startPortStr, err)
	}
	defer conn.Close()

	// Afficher le port ouvert
	fmt.Printf("Port %s ouvert\n", startPortStr)

	// Vérifier si le port de fin a été atteint
	if startPortStr == endPortStr {
		return
	}

	// Appeler la fonction récursivement
	GoPortScanner(target, strconv.Itoa(startPortInt + 1), endPortStr)
}

func PortScanner() {
	var (
		target string
		startPort string
		endPort string
	)

	fmt.Print("Adresse IP : ")
	_, err := fmt.Scan(&target)
	if err != nil {
		fmt.Println("Erreur de saisie :", err)
		return
	}

	fmt.Print("Port de départ : ")
	_, err = fmt.Scan(&startPort)
	if err != nil {
		fmt.Println("Erreur de saisie :", err)
		return
	}

	fmt.Print("Port de fin : ")
	_, err = fmt.Scan(&endPort)
	if err != nil {
		fmt.Println("Erreur de saisie :", err)
		return
	}

	// Affichage du menu
	fmt.Println("\n1. Utiliser Nmap ")
	fmt.Println("2. Utiliser Go ")

	// Demande du choix de l'utilisateur
	var choice int
	fmt.Print("\nChoisissez une option : ")
	_, err = fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Erreur de saisie :", err)
		return
	}

	switch choice {
		case 1:
			fmt.Print("Veuillez patienter... \n")
			NmapPortScanner(target, startPort, endPort)
		case 2:
			fmt.Print("Veuillez patienter... \n")
			GoPortScanner(target, startPort, endPort)
		default:
			fmt.Println("Option non valide. Veuillez choisir à nouveau.")
	}

}