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
	BgBlue  = "\033[44m"
	BgCyan  = "\033[46m"
)

// Colorize retourne un texte coloré
func Colorize(text, color string) string {
	return color + text + Reset
}

// ==================== EFFACEMENT ÉCRAN ====================

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

// ==================== EFFETS TEXTE ====================

// TypeWriter affiche un texte lettre par lettre (effet machine à écrire)
func TypeWriter(text string, delay time.Duration) {
	for _, ch := range text {
		fmt.Print(string(ch))
		time.Sleep(delay)
	}
	fmt.Println()
}

// SlowPrint affiche un texte lettre par lettre (alias de TypeWriter)
func SlowPrint(text string, delay time.Duration) {
	for _, ch := range text {
		fmt.Print(string(ch))
		time.Sleep(delay)
	}
	fmt.Println()
}

// Dim retourne un texte en gris (moins visible)
func Dim(s string) string {
	return "\033[2m" + s + "\033[22m"
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
                    
     ==++#%%#    
     #==++%**++     
   .# =+:..:++++    
   *  =+  @@++++=   
  -*. =+. =+++++++- 
-+-====+++##.++++++-
:=+++++++++++++****+
 ***####**#######*+ 
      +++++++=      
    -=::::--==      
  =***###**+=+**+   
   ##%@    ##*###   
	 `

const MarioArt = `
                         
          +#*.      
        :###M*##    
       *%%*+::-     
       =+%-=++-:    
    .-**##+=-+= .-. 
  -:... #%##*#%+=*= 
   *-   ####-*      
    #%%*%%###       
  ##%%%%@%%##*- +   
            %###*#  
             %##    
`

const LuigiArt = `
                    
                    
          ..:-      
           .:-.     
      .=+***.:+     
     -#*=L=+***     
       =:=@@=**#     
         :+=**#     
      .-   +#*#     
       +. %###%%    
        #%%###%%    
      +##%%%%%%%    
     %%%#+  @#%     
       ##   %%      
           ##%      
          **#       		  `

const PeachArt = `
                    
      :*:           
    ==-P-==    
     -=:=++     
      :+:   +
     .-#- -=      
    .+*=+**       
     ++++*#..-      
    +::--+**#     
   -::---==+*       
   ------===+*      
  ---------==+++
  ---------===++++
 +-------=--===**   
   +===+++++++      
   `

const ToadArt = `
     ...+++++...    
   ...++++++++.:::  
  ....+++++++**:::= 
  ....+++++***:::-+ 
  +....::::::::---* 
  +:---=========-=  
    ==+:...:#:-+=   
        ::::::-     
       .-----=      
        -::::::     
   .:::-::.::##-:   
    - :.....:::= -  
      =:...:::-=    
     +-*#=-==#***   
      .:    -####   
	  `

const KoopaArt = `
       
     -. ..:     
     :: :.:      
   =-::K::::       
   =+--:--=-       
    ++==*+-     
       %*:#:%%    
   ++  =-::-::%%   
     --*=-::-#-%%   
        *+--+-%=%%    
       ##%= +-*-%%   
     %##%%%*++#%%     
      *+     %%# 
	`

const ThwompArt = `
      
     AAAAAAAAAAA
   <#############>  
   <##+#:==+=#-*=>    
   <=#  @*-  %=*->  
   <=# .===. ==*-> 
   <=##vvvvvvv####>  
   <##*      **=##>   
   <####AAAAAA####>  
     vvvvvvvvvvvv        
		  `

const BowserArt = `
   #     #####     # 
  ##  ###########  ##  
    ##  ######  ##   
    ##############  
   ###.########.### 
    ##. #    # .##: 
     *#        ##  
      # #   # #-   
      ########    	   
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

const VictoryArt = `
 ██╗   ██╗██╗ ██████╗████████╗ ██████╗ ██████╗ ██╗   ██╗
 ██║   ██║██║██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗╚██╗ ██╔╝
 ██║   ██║██║██║        ██║   ██║   ██║██████╔╝ ╚████╔╝
 ╚██╗ ██╔╝██║██║        ██║   ██║   ██║██╔══██╗  ╚██╔╝
  ╚████╔╝ ██║╚██████╗   ██║   ╚██████╔╝██║  ██║   ██║
   ╚═══╝  ╚═╝ ╚═════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝   ╚═╝
`

// ==================== ASCII PNJ (Marchand / Forgeron) ====================

const MerchantArt = `
                    
         ::-        
        :.:--       
     -:...::--:-    
.-::....+:::=::----.
    -:::%:::%::--   
     :::::::::-:    
     ::---------    
     :-:     :-:   
	 
	 `

const BlacksmithArt = `
    ------=...-=    
   =-----===.%-+    
   =======++--=*    
    ++++****...-    
 --=*  +@@+-:--     
 -#=#  -=+*++#      
   *+##...--===---- 
     .. .--=-+  ==-+
     ..---====++ ** 
  ----=-====+=+     
   ==-=+=****+      
     ==.=+***=-=    
           +===    
          +=-==     

	  `

// ==================== MENUS ET TITRES ====================

// PrintTitle affiche un titre encadré
func PrintTitle(title string) {
	// Calculer la largeur visible (sans les codes ANSI)
	visibleLen := visibleLength(title)
	line := strings.Repeat("═", visibleLen+4)

	fmt.Println(Cyan + "╔" + line + "╗" + Reset)
	fmt.Printf(Cyan+"║"+Reset+"  %s%s%s  "+Cyan+"║\n"+Reset,
		Bold+Yellow, title, Reset)
	fmt.Println(Cyan + "╚" + line + "╝" + Reset)
}

// visibleLength retourne la longueur visible d'un string (sans codes ANSI)
func visibleLength(s string) int {
	length := 0
	inEscape := false
	for _, r := range s {
		if r == '\033' {
			inEscape = true
			continue
		}
		if inEscape {
			if r == 'm' {
				inEscape = false
			}
			continue
		}
		length++
	}
	return length
}

// PrintSeparator affiche une ligne de séparation
func PrintSeparator() {
	fmt.Println(Cyan + strings.Repeat("─", 50) + Reset)
}

// PrintBox affiche un texte dans un cadre simple
func PrintBox(text string) {
	lines := strings.Split(text, "\n")
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	border := strings.Repeat("═", maxLen+4)

	fmt.Println(Cyan + "╔" + border + "╗" + Reset)
	for _, line := range lines {
		pad := maxLen - len(line)
		fmt.Printf(Cyan+"║"+Reset+"  %s%s  "+Cyan+"║\n"+Reset,
			line, strings.Repeat(" ", pad))
	}
	fmt.Println(Cyan + "╚" + border + "╝" + Reset)
}

// ==================== INTERACTION UTILISATEUR ====================

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

// AskString demande une chaîne de caractères
func AskString(prompt string) string {
	fmt.Print(Yellow + prompt + Reset)
	var s string
	fmt.Scanln(&s)
	return s
}

// Pause attend que l'utilisateur appuie sur Entrée
func Pause() {
	fmt.Print(Green + "\n[Appuyez sur Entrée pour continuer...]" + Reset)
	var discard string
	fmt.Scanln(&discard)
}

// Confirm demande une confirmation Oui/Non
func Confirm(prompt string) bool {
	fmt.Print(Yellow + prompt + " (o/n) : " + Reset)
	var response string
	fmt.Scanln(&response)
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "o" || response == "oui" || response == "y" || response == "yes"
}

// ==================== BANDEAUX D'INFORMATION ====================

// PrintHeader affiche un bandeau de section
func PrintHeader(title string) {
	fmt.Println()
	fmt.Println(Cyan + "┌" + strings.Repeat("─", 50) + "┐" + Reset)
	fmt.Printf(Cyan+"│"+Reset+"  %s%s%s%s"+Cyan+"│\n"+Reset,
		Bold+Yellow, title,
		Reset, strings.Repeat(" ", 46-visibleLength(title)))
	fmt.Println(Cyan + "└" + strings.Repeat("─", 50) + "┘" + Reset)
}

// PrintSuccess affiche un message de succès
func PrintSuccess(msg string) {
	fmt.Println(Green + "✓ " + msg + Reset)
}

// PrintError affiche un message d'erreur
func PrintError(msg string) {
	fmt.Println(Red + "❌ " + msg + Reset)
}

// PrintInfo affiche un message d'information
func PrintInfo(msg string) {
	fmt.Println(Cyan + "ℹ " + msg + Reset)
}

// PrintWarning affiche un avertissement
func PrintWarning(msg string) {
	fmt.Println(Yellow + "⚠ " + msg + Reset)
}

// ==================== BARRES VISUELLES ====================

// SimpleBar affiche une barre générique (sans couleur)
func SimpleBar(current, max int, width int) string {
	if max <= 0 {
		return strings.Repeat("░", width)
	}
	filled := current * width / max
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	return "[" + strings.Repeat("█", filled) +
		strings.Repeat("░", width-filled) + "]"
}

// ProgressBar affiche une barre de progression
func ProgressBar(current, max int, width int) string {
	return SimpleBar(current, max, width)
}

// ==================== ANIMATIONS ====================

// LoadingDots affiche une petite animation de chargement
func LoadingDots() {
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println()
}

// FlashText fait clignoter un texte
func FlashText(text string, color string, times int) {
	for i := 0; i < times; i++ {
		fmt.Print("\r" + color + text + Reset)
		time.Sleep(150 * time.Millisecond)
		fmt.Print("\r" + strings.Repeat(" ", len(text)))
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println()
}

// Spinner affiche un spinner pendant une durée donnée
func Spinner(duration time.Duration) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	end := time.Now().Add(duration)
	i := 0
	for time.Now().Before(end) {
		fmt.Printf("\r%s %s", Cyan+frames[i%len(frames)]+Reset, "Chargement...")
		time.Sleep(80 * time.Millisecond)
		i++
	}
	fmt.Print("\r                          \r")
}

// ==================== SONS DU TERMINAL ====================

// Beep émet un bip simple
func Beep() {
	fmt.Print("\a")
}

// BeepVictory émet 3 bips (victoire)
func BeepVictory() {
	for i := 0; i < 3; i++ {
		fmt.Print("\a")
		time.Sleep(150 * time.Millisecond)
	}
}

// BeepDamage émet un bip grave (dégâts)
func BeepDamage() {
	fmt.Print("\a")
	time.Sleep(100 * time.Millisecond)
	fmt.Print("\a")
}

// ==================== AFFICHAGE DIVERS ====================

// Center centre un texte dans une largeur donnée
func Center(text string, width int) string {
	visible := visibleLength(text)
	if visible >= width {
		return text
	}
	totalPad := width - visible
	leftPad := totalPad / 2
	rightPad := totalPad - leftPad
	return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
}

// Truncate tronque un texte à une longueur maximale
func Truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

// RepeatChar répète un caractère N fois
func RepeatChar(c string, n int) string {
	return strings.Repeat(c, n)
}

// ==================== CONSTANTES DE STYLE ====================

// IconSet contient les icônes utilisées dans l'interface
var IconSet = map[string]string{
	"player":   "🧙",
	"hp":       "❤",
	"mana":     "💧",
	"gold":     "💰",
	"exp":      "⭐",
	"attack":   "⚔",
	"defense":  "🛡",
	"potion":   "🧪",
	"book":     "📖",
	"bag":      "🎒",
	"sword":    "🗡",
	"shield":   "🛡",
	"star":     "★",
	"check":    "✓",
	"cross":    "❌",
	"arrow":    "→",
	"heart":    "♥",
	"skull":    "☠",
	"fire":     "🔥",
	"sparkles": "✨",
	"trophy":   "🏆",
}

// Icon retourne une icône par son nom
func Icon(name string) string {
	if icon, ok := IconSet[name]; ok {
		return icon
	}
	return "•"
}

// ==================== EFFETS DE COULEUR ====================

// Rainbow retourne un texte avec chaque lettre d'une couleur différente
func Rainbow(text string) string {
	colors := []string{Red, Yellow, Green, Cyan, Blue, Purple}
	var sb strings.Builder
	for i, ch := range text {
		sb.WriteString(colors[i%len(colors)])
		sb.WriteRune(ch)
	}
	sb.WriteString(Reset)
	return sb.String()
}

// HpColorClass retourne la classe de couleur selon un ratio (0-1)
func HpColorClass(ratio float64) string {
	switch {
	case ratio > 0.6:
		return Green
	case ratio > 0.3:
		return Yellow
	default:
		return Red
	}
}

// ==================== AIDE ET RACCOURCIS ====================

// PrintHelp affiche une aide générique
func PrintHelp(commands map[string]string) {
	fmt.Println()
	fmt.Println(Cyan + "Commandes disponibles :" + Reset)
	for cmd, desc := range commands {
		fmt.Printf("  %s%-12s%s %s\n", Yellow, cmd, Reset, desc)
	}
	fmt.Println()
}

// PrintDivider affiche un séparateur décoratif
func PrintDivider() {
	fmt.Println(Purple + "◆ " + strings.Repeat("─", 48) + " ◆" + Reset)
}

// PrintDoubleDivider affiche un double séparateur
func PrintDoubleDivider() {
	fmt.Println(Purple + "◆ " + strings.Repeat("═", 48) + " ◆" + Reset)
}

// ==================== ÉCRANS SPÉCIAUX ====================

// PrintVictory affiche l'écran de victoire
func PrintVictory() {
	Clear()
	fmt.Println(Green + VictoryArt + Reset)
	fmt.Println()
	fmt.Println(Green + "        ★ ★ ★ VICTOIRE ! ★ ★ ★" + Reset)
	fmt.Println()
}

// PrintGameOver affiche l'écran de game over
func PrintGameOver() {
	Clear()
	fmt.Println(Red + GameOverArt + Reset)
	fmt.Println()
	fmt.Println(Red + "        ☠ GAME OVER ☠" + Reset)
	fmt.Println()
}

// PrintLevelUp affiche l'écran de montée de niveau
func PrintLevelUp(level int) {
	fmt.Println()
	fmt.Println(Yellow + "╔════════════════════════════════════════╗" + Reset)
	fmt.Println(Yellow + "║       ★ ★ ★ NIVEAU SUPÉRIEUR ★ ★ ★      ║" + Reset)
	fmt.Printf(Yellow+"║            %sNiveau %d%s                    ║\n"+Reset,
		Bold+Green, level, Reset+Yellow)
	fmt.Println(Yellow + "╚════════════════════════════════════════╝" + Reset)
	fmt.Println()
}
