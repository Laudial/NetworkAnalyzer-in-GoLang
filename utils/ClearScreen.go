package utils

import (
	"fmt"
    "os"
	"os/exec"
	"runtime"
)

// ClearScreen nettoye l'écran du terminal.
func ClearScreen() {
    switch runtime.GOOS {
    case "windows":
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    case "linux", "darwin":
        fmt.Print("\033[H\033[2J")
    default:
        fmt.Println("Système d'exploitation non pris en charge pour le nettoyage du terminal.")
        return
    }
}