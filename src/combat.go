package main

import (
	"fmt"
	"strings"
	"time"
)

// ==================== UTILITAIRES VISUELS ====================

// HPBar affiche une barre de vie colorée : [████████░░░░░░░░]
func HPBar(current, max int, width int) string {
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

	var color string
	ratio := float64(current) / float64(max)
	switch {
	case ratio > 0.6:
		color = Green
	case ratio > 0.3:
		color = Yellow
	default:
		color = Red
	}

	bar := color + strings.Repeat("█", filled) + Reset +
		White + strings.Repeat("░", width-filled) + Reset

	return "[" + bar + "]"
}

// ManaBar affiche une barre de mana bleue
func ManaBar(current, max int, width int) string {
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
	bar := Blue + strings.Repeat("█", filled) + Reset +
		White + strings.Repeat("░", width-filled) + Reset
	return "[" + bar + "]"
}

// FlashDamage fait clignoter un message de dégâts
func FlashDamage() {
	for i := 0; i < 3; i++ {
		fmt.Print("\r" + BgRed + "  💥 DÉGÂTS 💥  " + Reset)
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\r                  ")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

// AnimateAttack anime une attaque : "Mario »»»»» 💥 Goomba"
func AnimateAttack(attacker, target string) {
	fmt.Printf("\n%s%s%s ", Bold, attacker, Reset)
	for i := 0; i < 5; i++ {
		fmt.Print(Yellow + "»" + Reset)
		time.Sleep(80 * time.Millisecond)
	}
	fmt.Print(" 💥 ")
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("%s%s%s\n", Bold, target, Reset)
	time.Sleep(300 * time.Millisecond)
}

// ⚠️ SlowPrint et LoadingDots sont définis dans ui.go — NE PAS les redéfinir ici

// ==================== AFFICHAGE DU COMBAT ====================

// DisplayMonster affiche l'ASCII art du monstre
func DisplayMonster(m *Monster) {
	switch m.Name {
	case "Goomba d'entraînement":
		fmt.Println(Red + GoombaArt + Reset)
	case "Thwomp":
		fmt.Println(Blue + ThwompArt + Reset)
	case "Bowser":
		fmt.Println(Red + BowserArt + Reset)
	}
}

// DisplayCombatUI affiche l'interface complète du combat
func DisplayCombatUI(player *Character, monster *Monster, turn int) {
	Clear()

	// Bandeau du tour
	fmt.Println(Cyan + "╔══════════════════════════════════════════════════════════╗" + Reset)
	fmt.Printf(Cyan+"║"+Reset+"                  %s⚔  TOUR %d ⚔%s                     "+Cyan+"║\n"+Reset,
		Bold+Yellow, turn, Reset)
	fmt.Println(Cyan + "╚══════════════════════════════════════════════════════════╝" + Reset)
	fmt.Println()

	// Zone du monstre
	fmt.Println(Red + "                    ╔════════════════╗" + Reset)
	fmt.Printf(Red+"                    ║"+Reset+"  %-12s  "+Red+"║\n"+Reset, monster.Name)
	fmt.Println(Red + "                    ╚════════════════╝" + Reset)

	// Art du monstre
	DisplayMonster(monster)

	// Barre de vie monstre
	monsterHP := HPBar(monster.CurrentHP, monster.MaxHP, 20)
	fmt.Printf("       %s%s%s  %s  %s%d/%d%s\n\n",
		Bold, monster.Name, Reset,
		monsterHP,
		Red, monster.CurrentHP, monster.MaxHP, Reset)

	// Séparateur
	fmt.Println(White + "──────────────────────────────────────────────────────────" + Reset)
	fmt.Println()

	// Art du joueur
	DisplayCharacterArt(player.Class)

	// Barres du joueur
	playerHP := HPBar(player.CurrentHP, player.MaxHP, 20)
	playerMana := ManaBar(player.Mana, player.ManaMax, 20)

	fmt.Printf("       %s%s%s  %s  %s%d/%d PV%s\n",
		Bold, player.Name, Reset,
		playerHP,
		Green, player.CurrentHP, player.MaxHP, Reset)
	fmt.Printf("       %sMana%s   %s  %s%d/%d%s\n",
		Blue, Reset,
		playerMana,
		Blue, player.Mana, player.ManaMax, Reset)
	fmt.Printf("       %sPièces : %d%s   %sExp : %d/%d%s\n\n",
		Yellow, player.Gold, Reset,
		Purple, player.Exp, player.ExpMax, Reset)

	fmt.Println(White + "──────────────────────────────────────────────────────────" + Reset)
}

// ==================== COMBAT PRINCIPAL ====================

// TrainingFight lance un combat tour par tour
func TrainingFight(player *Character, monster Monster) {
	// 🎵 Démarrer la musique de combat (arrête la musique de fond)
	PlayBattleMusic()

	// À la fin du combat, on remet la musique de fond
	defer func() {
		StopBattleMusic()
		PlayBackgroundMusic()
	}()

	Clear()

	// Écran de rencontre
	DisplayMonster(&monster)
	SlowPrint(Yellow+"⚡ Un "+monster.Name+" apparaît !", 30*time.Millisecond)
	fmt.Print(White + "Préparation au combat" + Reset)
	LoadingDots()
	time.Sleep(500 * time.Millisecond)
	Pause()

	turn := 1

	// Initiative (Mission 1)
	playerFirst := player.Initiative >= monster.Initiative

	Clear()
	if playerFirst {
		SlowPrint(Green+"⚡ Vous êtes plus rapide que le monstre !", 25*time.Millisecond)
	} else {
		SlowPrint(Purple+"⚡ Le monstre est plus rapide que vous !", 25*time.Millisecond)
	}
	time.Sleep(1 * time.Second)

	// Boucle de combat
	for player.CurrentHP > 0 && monster.CurrentHP > 0 {
		// Affichage du tour
		DisplayCombatUI(player, &monster, turn)
		time.Sleep(400 * time.Millisecond)

		if playerFirst {
			// Le joueur commence
			player.CharacterTurn(&monster)
			if monster.CurrentHP <= 0 {
				break
			}
			fmt.Println()
			fmt.Print(Yellow + "Le monstre riposte" + Reset)
			LoadingDots()
			time.Sleep(400 * time.Millisecond)
			DisplayCombatUI(player, &monster, turn)
			time.Sleep(300 * time.Millisecond)
			monster.Pattern(turn, player)
		} else {
			// Le monstre commence
			monster.Pattern(turn, player)
			if player.CurrentHP <= 0 {
				break
			}
			fmt.Println()
			fmt.Print(Green + "À vous de jouer" + Reset)
			LoadingDots()
			time.Sleep(400 * time.Millisecond)
			DisplayCombatUI(player, &monster, turn)
			time.Sleep(300 * time.Millisecond)
			player.CharacterTurn(&monster)
		}

		turn++
		Pause()
	}

	// ==================== FIN DU COMBAT ====================
	Clear()

	if player.CurrentHP <= 0 {
		// ☠ DÉFAITE — son identique pour TOUS les combats
		fmt.Println(Red + GameOverArt + Reset)
		time.Sleep(500 * time.Millisecond)

		SlowPrint(Red+"Vous vous effondrez...", 40*time.Millisecond)

		// 🎵 Son de mort de Mario
		PlayDeathSound()
		time.Sleep(2 * time.Second)

		player.IsDead()
	} else {
		// ★ VICTOIRE — son identique pour TOUS les combats
		fmt.Println(Green + "╔══════════════════════════════════════╗" + Reset)
		fmt.Println(Green + "║          ★ ★ ★ VICTOIRE ! ★ ★ ★       ║" + Reset)
		fmt.Println(Green + "╚══════════════════════════════════════╝" + Reset)

		// 🎵 Son de victoire
		PlayVictorySound()

		time.Sleep(500 * time.Millisecond)
		fmt.Println()
		SlowPrint(Yellow+fmt.Sprintf("Vous gagnez %d EXP et %d pièces !",
			monster.ExpReward, monster.GoldReward), 25*time.Millisecond)
		player.GainExp(monster.ExpReward)
		player.Gold += monster.GoldReward
	}
	Pause()
}

// ==================== TOUR DU JOUEUR ====================

// CharacterTurn gère le tour du joueur (menu d'action)
func (c *Character) CharacterTurn(monster *Monster) {
	for {
		fmt.Printf("  %s1.%s %s Attaquer (Saut)%s\n", Cyan, Reset, Bold, Reset)
		fmt.Printf("  %s2.%s Inventaire\n", Cyan, Reset)
		fmt.Printf("  %s3.%s Sorts (%sMana: %d/%d%s)\n", Cyan, Reset, Blue, c.Mana, c.ManaMax, Reset)

		choice := AskInt("\nVotre choix : ")

		switch choice {
		case 1:
			c.PlayerAttack(monster, "Saut", 5)
			return // ✅ tour consommé

		case 2:
			// ✅ Utilise la version combat qui dit si un item a été utilisé
			if c.AccessInventoryInCombat() {
				return // ✅ tour consommé si un item a été utilisé
			}
			// Sinon on reste dans le menu du tour
			Clear()
			DisplayCombatUI(c, monster, 0)
			continue

		case 3:
			if c.UseSpellInCombat(monster) {
				return // ✅ tour consommé si un sort a été lancé
			}

		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}

// PlayerAttack exécute une attaque physique du joueur
func (c *Character) PlayerAttack(monster *Monster, attackName string, damage int) {
	Clear()
	fmt.Println(Cyan + "═══════════════════════════════════════════════════════════" + Reset)
	fmt.Println()

	AnimateAttack(c.Name, monster.Name)
	FlashDamage()

	monster.CurrentHP -= damage
	if monster.CurrentHP < 0 {
		monster.CurrentHP = 0
	}

	SlowPrint(Green+fmt.Sprintf("%s utilise %s et inflige %d dégâts !",
		c.Name, attackName, damage), 15*time.Millisecond)
	time.Sleep(400 * time.Millisecond)

	// Afficher la barre de vie mise à jour
	fmt.Println()
	fmt.Printf("  %s%s%s  %s  %s%d/%d%s\n",
		Bold, monster.Name, Reset,
		HPBar(monster.CurrentHP, monster.MaxHP, 25),
		Red, monster.CurrentHP, monster.MaxHP, Reset)
}

// ==================== SORTS EN COMBAT ====================

// UseSpellInCombat affiche le menu des sorts et lance le sort choisi
func (c *Character) UseSpellInCombat(monster *Monster) bool {
	Clear()
	fmt.Println(Cyan + "╔════════════════════════════════════════╗" + Reset)
	fmt.Println(Cyan + "║              📖 SORTS                  ║" + Reset)
	fmt.Println(Cyan + "╚════════════════════════════════════════╝" + Reset)
	fmt.Printf("  Mana : %s %s%d/%d%s\n\n",
		ManaBar(c.Mana, c.ManaMax, 15), Blue, c.Mana, c.ManaMax, Reset)

	for i, s := range c.Skills {
		cost := spellManaCost(s)
		dmg := spellDamage(s)
		color := Green
		if c.Mana < cost {
			color = Red
		}

		// Affichage différent pour le sort de soin
		if s == "Étoile" {
			fmt.Printf("  %s%d.%s %-15s %s(SOIN +40 PV, %d mana)%s\n",
				Cyan, i+1, Reset, s, color, cost, Reset)
		} else {
			fmt.Printf("  %s%d.%s %-15s %s(%d dégâts, %d mana)%s\n",
				Cyan, i+1, Reset, s, color, dmg, cost, Reset)
		}
	}
	fmt.Printf("  %s0.%s Retour\n", Yellow, Reset)

	choice := AskInt("\nChoix : ")
	if choice == 0 {
		return false
	}
	if choice < 1 || choice > len(c.Skills) {
		fmt.Println(Red + "Choix invalide." + Reset)
		time.Sleep(500 * time.Millisecond)
		return false
	}

	spell := c.Skills[choice-1]
	cost := spellManaCost(spell)

	if c.Mana < cost {
		fmt.Println(Red + "❌ Mana insuffisant !" + Reset)
		time.Sleep(800 * time.Millisecond)
		return false
	}

	c.Mana -= cost

	// 🎁 Bonus : sort de soin Étoile
	if spell == "Étoile" {
		return c.CastHealSpell()
	}

	damage := spellDamage(spell)

	// Animation du sort
	Clear()
	fmt.Println(Cyan + "═══════════════════════════════════════════════════════════" + Reset)
	fmt.Println()

	if spell == "Boule de Feu" {
		// Animation feu
		fmt.Printf("%s%s%s lance une ", Bold, c.Name, Reset)
		time.Sleep(200 * time.Millisecond)
		fmt.Print(Red + "🔥 " + Yellow + "BOULE " + Red + "DE " + Yellow + "FEU " + Red + "🔥" + Reset)
		time.Sleep(400 * time.Millisecond)
		fmt.Println()
		for i := 0; i < 5; i++ {
			fmt.Print(Yellow + "»" + Red + "»" + Yellow + "»" + Reset + " ")
			time.Sleep(80 * time.Millisecond)
		}
		fmt.Print(" 💥💥💥")
		fmt.Println()
	} else {
		AnimateAttack(c.Name, monster.Name)
	}

	FlashDamage()

	monster.CurrentHP -= damage
	if monster.CurrentHP < 0 {
		monster.CurrentHP = 0
	}

	time.Sleep(300 * time.Millisecond)
	SlowPrint(Green+fmt.Sprintf("%s lance %s et inflige %d dégâts !",
		c.Name, spell, damage), 15*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	fmt.Println()
	fmt.Printf("  %s%s%s  %s  %s%d/%d%s\n",
		Bold, monster.Name, Reset,
		HPBar(monster.CurrentHP, monster.MaxHP, 25),
		Red, monster.CurrentHP, monster.MaxHP, Reset)
	fmt.Printf("  %sMana restant : %s %s%d/%d%s\n",
		Blue, ManaBar(c.Mana, c.ManaMax, 15), Blue, c.Mana, c.ManaMax, Reset)
	return true
}

// ==================== SORT DE SOIN ÉTOILE ====================

// CastHealSpell lance le sort de soin Étoile (+40 PV)
func (c *Character) CastHealSpell() bool {
	Clear()
	fmt.Println(Cyan + "═══════════════════════════════════════════════════════════" + Reset)
	fmt.Println()

	fmt.Printf("%s%s%s invoque ", Bold, c.Name, Reset)
	time.Sleep(200 * time.Millisecond)
	fmt.Print(Yellow + "✨ ÉTOILE ✨" + Reset)
	time.Sleep(400 * time.Millisecond)
	fmt.Println()

	// Animation d'étoiles
	for i := 0; i < 5; i++ {
		fmt.Print(Yellow + "★ " + Reset)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()

	// Soin
	before := c.CurrentHP
	c.CurrentHP += 40
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	healed := c.CurrentHP - before

	time.Sleep(300 * time.Millisecond)
	SlowPrint(Green+fmt.Sprintf("Vous récupérez %d PV !", healed), 15*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	fmt.Println()
	fmt.Printf("  %s%s%s  %s  %s%d/%d%s\n",
		Bold, c.Name, Reset,
		HPBar(c.CurrentHP, c.MaxHP, 25),
		Green, c.CurrentHP, c.MaxHP, Reset)
	fmt.Printf("  %sMana restant : %s %s%d/%d%s\n",
		Blue, ManaBar(c.Mana, c.ManaMax, 15), Blue, c.Mana, c.ManaMax, Reset)
	return true
}

// ==================== COÛTS ET DÉGÂTS DES SORTS ====================

// spellManaCost retourne le coût en mana d'un sort
func spellManaCost(spell string) int {
	switch spell {
	case "Saut":
		return 5
	case "Boule de Feu":
		return 15
	case "Étoile":
		return 20
	}
	return 0
}

// spellDamage retourne les dégâts d'un sort (0 pour les soins)
func spellDamage(spell string) int {
	switch spell {
	case "Saut":
		return 8
	case "Boule de Feu":
		return 18
	case "Étoile":
		return 0 // soin, pas de dégâts
	}
	return 0
}
