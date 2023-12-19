package controllers

import (
	"fmt"
	"net"
	"strings"

	"NetworkAnalyzer/utils"
)

func DnsResolver() {
	var host string

	fmt.Print("Entrez le nom de domaine (example.com): ")
	_, err := fmt.Scan(&host)
	if err != nil {
		fmt.Println("Erreur de saisie :", err)
		return
	}

	if !strings.HasPrefix(host, "www") {
		host = "www." + host
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("Erreur lors de la résolution DNS :", err)
		return
	}

	utils.ClearScreen()

	var ipStrings []string
	for _, ip := range ips {
		ipStrings = append(ipStrings, ip.String())
	}

	msg := "Adresse(s) IP pour " + host + " : " + strings.Join(ipStrings, ", ")
	utils.PauseScreen(msg)

	utils.ClearScreen()
}