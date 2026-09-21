package main

import (
	"fmt"
	"strings"
	"time"
)

// ==================== UTILITAIRES VISUELS ====================

// Barre de vie colorée : [████████░░░░░░░░]
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

// Barre de mana
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

// Effet "flash rouge" quand on prend des dégâts
func FlashDamage() {
	for i := 0; i < 3; i++ {
		fmt.Print("\r" + BgRed + "  💥 DÉGÂTS 💥  " + Reset)
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\r                  ")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

// Attaque animée : "Mario >>>>>>  💥  Goomba"
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

// Texte qui s'écrit lettre par lettre
func SlowPrint(text string, delay time.Duration) {
	for _, ch := range text {
		fmt.Print(string(ch))
		time.Sleep(delay)
	}
	fmt.Println()
}

// Petite animation de chargement
func LoadingDots() {
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println()
}

// ==================== AFFICHAGE DU COMBAT ====================

// Affiche le monstre avec son ASCII art
func DisplayMonster(m *Monster) {
	switch m.Name {
	case "Goomba d'entraînement":
		fmt.Println(Red + GoombaArt + Reset)
	case "Koopa Tropique":
		fmt.Println(Green + GoombaArt + Reset)
	case "Bowser":
		fmt.Println(Red + BowserArt + Reset)
	}
}

// Affiche le tableau de combat complet
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
	fmt.Println(MarioArt)

	// Zone du joueur
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

func TrainingFight(player *Character, monster Monster) {
	Clear()

	// Écran de rencontre
	DisplayMonster(&monster)
	SlowPrint(Yellow+"⚡ Un "+monster.Name+" apparaît !", 30*time.Millisecond)
	fmt.Print(White + "Préparation au combat")
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

	for player.CurrentHP > 0 && monster.CurrentHP > 0 {
		// Affichage du tour
		DisplayCombatUI(player, &monster, turn)
		time.Sleep(400 * time.Millisecond)

		if playerFirst {
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

	// Écran de fin
	Clear()
	if player.CurrentHP <= 0 {
		fmt.Println(Red + GameOverArt + Reset)
		time.Sleep(500 * time.Millisecond)
		SlowPrint(Red+"Vous vous effondrez...", 40*time.Millisecond)
		time.Sleep(500 * time.Millisecond)
		player.IsDead()
	} else {
		fmt.Println(Green + "╔══════════════════════════════════════╗" + Reset)
		fmt.Println(Green + "║          ★ ★ ★ VICTOIRE ! ★ ★ ★       ║" + Reset)
		fmt.Println(Green + "╚══════════════════════════════════════╝" + Reset)
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

func (c *Character) CharacterTurn(monster *Monster) {
	for {
		fmt.Printf("  %s1.%s %s Attaquer (Saut)%s\n", Cyan, Reset, Bold, Reset)
		fmt.Printf("  %s2.%s Inventaire\n", Cyan, Reset)
		fmt.Printf("  %s3.%s Sorts (%sMana: %d/%d%s)\n", Cyan, Reset, Blue, c.Mana, c.ManaMax, Reset)

		choice := AskInt("\nVotre choix : ")

		switch choice {
		case 1:
			c.PlayerAttack(monster, "Saut", 5)
			return
		case 2:
			c.AccessInventory()
			return
		case 3:
			if c.UseSpellInCombat(monster) {
				return
			}
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}

// Attaque du joueur avec animation
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

func (c *Character) UseSpellInCombat(monster *Monster) bool {
	Clear()
	fmt.Println(Cyan + "╔════════════════════════════════════════╗" + Reset)
	fmt.Println(Cyan + "║              📖 SORTS                  ║" + Reset)
	fmt.Println(Cyan + "╚════════════════════════════════════════╝" + Reset)
	fmt.Printf("  Mana : %s %s%d/%d%s\n\n", ManaBar(c.Mana, c.ManaMax, 15), Blue, c.Mana, c.ManaMax, Reset)

	for i, s := range c.Skills {
		cost := spellManaCost(s)
		dmg := spellDamage(s)
		color := Green
		if c.Mana < cost {
			color = Red
		}
		fmt.Printf("  %s%d.%s %-15s %s(%d dégâts, %d mana)%s\n",
			Cyan, i+1, Reset, s, color, dmg, cost, Reset)
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

func spellManaCost(spell string) int {
	switch spell {
	case "Saut":
		return 5
	case "Boule de Feu":
		return 15
	}
	return 0
}

func spellDamage(spell string) int {
	switch spell {
	case "Saut":
		return 8
	case "Boule de Feu":
		return 18
	}
	return 0
}
