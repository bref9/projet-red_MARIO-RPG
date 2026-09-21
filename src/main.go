package main

import "fmt"

func main() {
	Clear()
	fmt.Println(Yellow + MarioLogo + Reset)
	fmt.Println(Cyan + "Bienvenue dans Mario RPG Adventure !" + Reset)
	Pause()

	var player *Character

	// Écran titre : nouvelle partie ou charger
	for {
		Clear()
		PrintTitle("MENU DE DÉMARRAGE")
		fmt.Printf("  %s1.%s Nouvelle partie\n", Cyan, Reset)
		if SaveExists() {
			fmt.Printf("  %s2.%s Charger la partie\n", Cyan, Reset)
		}
		fmt.Printf("  %s0.%s Quitter\n", Yellow, Reset)

		choice := AskInt("\nChoix : ")
		switch choice {
		case 1:
			player = CharacterCreation()
		case 2:
			if SaveExists() {
				loaded, ok := LoadGame()
				if ok {
					player = loaded
					fmt.Println(Green + "✓ Partie chargée !" + Reset)
					Pause()
				} else {
					fmt.Println(Red + "Erreur de chargement." + Reset)
					Pause()
					continue
				}
			} else {
				fmt.Println(Red + "Aucune sauvegarde." + Reset)
				Pause()
				continue
			}
		case 0:
			fmt.Println(Yellow + "À bientôt !" + Reset)
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
			continue
		}
		if player != nil {
			break
		}
	}

	// Menu principal
	for {
		Clear()
		PrintTitle("🍄 MENU PRINCIPAL")
		fmt.Printf("  %s1.%s Afficher les informations\n", Cyan, Reset)
		fmt.Printf("  %s2.%s Accéder à l'inventaire\n", Cyan, Reset)
		fmt.Printf("  %s3.%s Marchand\n", Cyan, Reset)
		fmt.Printf("  %s4.%s Forgeron\n", Cyan, Reset)
		fmt.Printf("  %s5.%s Entraînement (Goomba)\n", Cyan, Reset)
		fmt.Printf("  %s6.%s Combat contre Koopa\n", Cyan, Reset)
		fmt.Printf("  %s7.%s ★ BOSS : Bowser ★\n", Red, Reset)
		fmt.Printf("  %s8.%s Sauvegarder\n", Green, Reset)
		fmt.Printf("  %s9.%s Qui sont-ils ?\n", Cyan, Reset)
		fmt.Printf("  %s0.%s Quitter\n", Yellow, Reset)

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
			TrainingFight(player, InitKoopa())
		case 7:
			TrainingFight(player, InitBowser())
		case 8:
			SaveGame(player)
			Pause()
		case 9:
			WhoAreThey()
			Pause()
		case 0:
			fmt.Println(Yellow + "\nMerci d'avoir joué ! À bientôt !" + Reset)
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
		}
	}
}

func WhoAreThey() {
	fmt.Println(Yellow + "\n╔════════════════════════════════════════╗" + Reset)
	fmt.Println(Yellow + "║       LES ARTISTES CACHÉS              ║" + Reset)
	fmt.Println(Yellow + "╠════════════════════════════════════════╣" + Reset)
	fmt.Println(Yellow + "║  🎵 Koji Kondo                         ║" + Reset)
	fmt.Println(Yellow + "║     (compositeur de Mario)             ║" + Reset)
	fmt.Println(Yellow + "║  🎨 Shigeru Miyamoto                   ║" + Reset)
	fmt.Println(Yellow + "║     (créateur de Mario)                ║" + Reset)
	fmt.Println(Yellow + "╚════════════════════════════════════════╝" + Reset)
}
