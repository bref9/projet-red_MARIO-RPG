package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// ==================== COULEURS ====================

const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Purple  = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	BgRed   = "\033[41m"
	BgGreen = "\033[42m"
)

// Colorize retourne un texte coloré
func Colorize(text, color string) string {
	return color + text + Reset
}

// Clear efface l'écran du terminal
func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

// TypeWriter affiche un texte lettre par lettre (effet machine à écrire)
func TypeWriter(text string, delay time.Duration) {
	for _, ch := range text {
		fmt.Print(string(ch))
		time.Sleep(delay)
	}
	fmt.Println()
}

// ==================== ASCII ART ====================

const MarioLogo = `
 __  __    _    ____  ___ ___
|  \/  |  / \  |  _ \|_ _/ _ \
| |\/| | / _ \ | |_) || | | | |
| |  | |/ ___ \|  _ < | | |_| |
|_|  |_/_/   \_\_| \_\___\___/
        RPG ADVENTURE
`

const GoombaArt = `
     _____
    /     \
   | () () |
    \  ^  /
     |||||
`

const MarioArt = `
      ___
     /   \
    | o o |
    |  ^  |
    | --- |
     \___/
    /|   |\
`

const BowserArt = `
   /\___/\
  ( o   o )
  (  \_/  )
   \ === /
   /|   |\
`

const GameOverArt = `
  ██████╗  █████╗ ███╗   ███╗███████╗
 ██╔════╝ ██╔══██╗████╗ ████║██╔════╝
 ██║  ███╗███████║██╔████╔██║█████╗
 ██║   ██║██╔══██║██║╚██╔╝██║██╔══╝
 ╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗
  ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
            OVER
`

// ==================== MENUS GÉNÉRIQUES ====================

// PrintTitle affiche un titre encadré
func PrintTitle(title string) {
	line := strings.Repeat("═", len(title)+4)
	fmt.Println(Cyan + "╔" + line + "╗" + Reset)
	fmt.Println(Cyan + "║  " + Bold + Yellow + title + Reset + Cyan + "  ║" + Reset)
	fmt.Println(Cyan + "╚" + line + "╝" + Reset)
}

// PrintSeparator affiche une ligne de séparation
func PrintSeparator() {
	fmt.Println(Cyan + strings.Repeat("─", 50) + Reset)
}

// AskInt demande un entier à l'utilisateur
func AskInt(prompt string) int {
	fmt.Print(Yellow + prompt + Reset)
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		// vider le buffer en cas d'erreur
		var discard string
		fmt.Scanln(&discard)
		return -1
	}
	return choice
}

// Pause attend que l'utilisateur appuie sur Entrée
func Pause() {
	fmt.Print(Green + "\n[Appuyez sur Entrée pour continuer...]" + Reset)
	var discard string
	fmt.Scanln(&discard)
}
