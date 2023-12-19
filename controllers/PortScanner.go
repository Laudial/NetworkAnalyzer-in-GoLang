package controllers

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
	"log"

	"NetworkAnalyzer/utils"

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
	targetAddr := net.ParseIP(target)
	maxGoroutines := 5000 // Nombre maximal de goroutines autorisées
	semaphore := make(chan struct{}, maxGoroutines)

	// Vérifier l'adresse IP
	if targetAddr == nil {
		fmt.Println("Adresse IP invalide")
		return
	}

	// Convertir l'adresse IP en entier 32 bits
	targetIPInt := ipToInt(targetAddr)

	// Démarrer le groupe d'attente
	var wg sync.WaitGroup

	// Scan
	for ip := targetIPInt; ip <= targetIPInt; ip++ {
		ipStr := intToIP(ip).String() // IP en chaîne de caractères
		startPortInt, err := strconv.Atoi(startPort)
		if err != nil {
			log.Fatalf("Erreur lors de la conversion du port de départ: %v", err)
		}
		endPortInt, err := strconv.Atoi(endPort)
		if err != nil {
			log.Fatalf("Erreur lors de la conversion du port de fin: %v", err)
		}

		var wg sync.WaitGroup

		// Vérifier les ports
		for port := startPortInt; port <= endPortInt; port++ {
			wg.Add(1)
			semaphore <- struct{}{}
			go func(ip, portStr string) {
				defer func() {
					<-semaphore
					wg.Done()
				}()
				if isPortOpen(ip, portStr) {
					result := fmt.Sprintf("%s:%s\n", ip, portStr)
					fmt.Print(result)
				}
			}(ipStr, strconv.Itoa(port))
		}
	}

	// Attendre la fin des goroutines
	go func() {
		wg.Wait()
		close(semaphore) // Fermer le canal une fois que toutes les goroutines sont terminées
	}()

	<-semaphore // Attendre que le canal soit fermé
	utils.PauseScreen("Scan terminé")
}

// Convertir l'adresse IP en entier 32 bits
func ipToInt(ip net.IP) int {
	ip = ip.To4()
	if ip == nil {
		// Adresse IP invalide
		return 0
	}
	return int(ip[0])<<24 + int(ip[1])<<16 + int(ip[2])<<8 + int(ip[3])
}

// Convertir l'entier 32 bits en adresse IP
func intToIP(ip int) net.IP {
	return net.IPv4(byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

// Vérifier si le port est ouvert
func isPortOpen(ip, port string) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ip, port), 10*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
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
			utils.ClearScreen()
			fmt.Print("Scan en cours... \n")
			NmapPortScanner(target, startPort, endPort)
		case 2:
			utils.ClearScreen()
			fmt.Print("Scan en cours... \n")
			GoPortScanner(target, startPort, endPort)
		default:
			fmt.Println("Option non valide. Veuillez choisir à nouveau.")
	}

}