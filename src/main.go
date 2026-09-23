package main

import (
	"fmt"
	"time"
)

func main() {
	// 🎵 Démarrer la musique de fond
	PlayBackgroundMusic()

	// 🔇 Arrêter la musique quand on quitte
	defer StopBackgroundMusic()

	Clear()
	fmt.Println(Yellow + MarioLogo + Reset)
	fmt.Println(Cyan + "Bienvenue dans Mario RPG Adventure !" + Reset)
	Pause()

	// ==================== CRÉATION DU PERSONNAGE ====================
	player := CharacterCreation()

	// ==================== MENU PRINCIPAL ====================
	for {
		Clear()
		PrintTitle("🍄 MENU PRINCIPAL")
		fmt.Printf("  %s1.%s  Afficher les informations\n", Cyan, Reset)
		fmt.Printf("  %s2.%s  Accéder à l'inventaire\n", Cyan, Reset)
		fmt.Printf("  %s3.%s  Marchand\n", Cyan, Reset)
		fmt.Printf("  %s4.%s  Forgeron\n", Cyan, Reset)
		fmt.Println()
		fmt.Printf("  %s5.%s  Entraînement (Goomba)\n", Green, Reset)
		fmt.Printf("  %s6.%s  Combat contre Thwomp\n", Green, Reset)
		fmt.Printf("  %s7.%s  ★ BOSS : Bowser ★\n", Red, Reset)
		fmt.Println()
		fmt.Printf("  %s8.%s  Qui sont-ils ?\n", Cyan, Reset)
		fmt.Printf("  %s0.%s  Quitter\n", Yellow, Reset)

		choice := AskInt("\nChoix : ")

		switch choice {
		case 1:
			Clear()
			player.DisplayInfo()
			Pause()

		case 2:
			player.AccessInventory()

		case 3:
			player.Merchant()

		case 4:
			player.Blacksmith()

		case 5:
			TrainingFight(player, InitGoomba())

		case 6:
			TrainingFight(player, InitThwomp())

		case 7:
			TrainingFight(player, InitBowser())

		case 8:
			WhoAreThey()
			Pause()

		case 0:
			// 🎵 Son de sortie
			PlayQuitSound()

			// 🔇 Arrêter la musique de fond
			StopBackgroundMusic()

			Clear()
			fmt.Println()
			fmt.Println(Yellow + "╔════════════════════════════════════════╗" + Reset)
			fmt.Println(Yellow + "║     👋 MERCI D'AVOIR JOUÉ ! 👋          ║" + Reset)
			fmt.Println(Yellow + "║                                        ║" + Reset)
			fmt.Println(Yellow + "║          À BIENTÔT DANS MARIO !        ║" + Reset)
			fmt.Println(Yellow + "╚════════════════════════════════════════╝" + Reset)
			fmt.Println()

			// Laisser le son se terminer
			time.Sleep(2 * time.Second)
			return

		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
		}
	}
}

// ==================== QUI SONT-ILS ? (Mission 6) ====================

func WhoAreThey() {
	fmt.Println()
	fmt.Println(Yellow + "╔════════════════════════════════════════════════╗" + Reset)
	fmt.Println(Yellow + "║          🎭 LES ARTISTES CACHÉS 🎭             ║" + Reset)
	fmt.Println(Yellow + "╠════════════════════════════════════════════════╣" + Reset)
	fmt.Println(Yellow + "║                                                ║" + Reset)
	fmt.Println(Yellow + "║                     ABBA                       ║" + Reset)
	fmt.Println(Yellow + "║                                                ║" + Reset)
	fmt.Println(Yellow + "║                     QUEEN                      ║" + Reset)
	fmt.Println(Yellow + "║                                                ║" + Reset)
	fmt.Println(Yellow + "╚════════════════════════════════════════════════╝" + Reset)
}
