package utils

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func PauseScreen(msg string) {

	if err := termbox.Init(); err != nil {
		fmt.Println("Erreur lors de l'initialisation de termbox:", err)
		os.Exit(1)
	}
	defer termbox.Close()
	
	fmt.Println(msg, "\nAppuyez sur Entrée pour continuer...")

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEnter {
			break
		}
	}

}