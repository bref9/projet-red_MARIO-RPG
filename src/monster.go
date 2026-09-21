package main

import (
	"fmt"
	"time"
)

// ==================== STRUCTURE ====================

type Monster struct {
	Name       string
	MaxHP      int
	CurrentHP  int
	Attack     int
	Initiative int
	ExpReward  int
	GoldReward int
}

// ==================== INITIALISATIONS ====================

func InitGoomba() Monster {
	return Monster{
		Name:       "Goomba d'entraînement",
		MaxHP:      40,
		CurrentHP:  40,
		Attack:     5,
		Initiative: 8,
		ExpReward:  50,
		GoldReward: 10,
	}
}

func InitKoopa() Monster {
	return Monster{
		Name:       "Koopa Tropique",
		MaxHP:      60,
		CurrentHP:  60,
		Attack:     8,
		Initiative: 6,
		ExpReward:  80,
		GoldReward: 20,
	}
}

func InitBowser() Monster {
	return Monster{
		Name:       "Bowser",
		MaxHP:      150,
		CurrentHP:  150,
		Attack:     15,
		Initiative: 12,
		ExpReward:  300,
		GoldReward: 100,
	}
}

// ==================== PATTERN D'ATTAQUE ====================

// Pattern du monstre : attaque normale, sauf tous les 3 tours (x2)
func (m *Monster) Pattern(turn int, player *Character) {
	damage := m.Attack
	special := false

	if turn%3 == 0 {
		damage = m.Attack * 2
		special = true
	}

	// Nettoyage et séparateur
	Clear()
	fmt.Println(Red + "═══════════════════════════════════════════════════════════" + Reset)
	fmt.Println()

	// Annonce de l'attaque spéciale
	if special {
		SlowPrint(Purple+"⚡ Le monstre prépare une attaque spéciale !", 20*time.Millisecond)
		time.Sleep(600 * time.Millisecond)
		fmt.Println()
	}

	// Animation de l'attaque
	AnimateAttack(m.Name, player.Name)

	// Effet de flash sur les dégâts
	FlashDamage()

	// Application des dégâts
	player.CurrentHP -= damage
	if player.CurrentHP < 0 {
		player.CurrentHP = 0
	}

	// Message de dégâts
	SlowPrint(Red+fmt.Sprintf("%s inflige %d dégâts à %s !",
		m.Name, damage, player.Name), 15*time.Millisecond)

	time.Sleep(300 * time.Millisecond)
	fmt.Println()

	// Affichage de la barre de vie mise à jour
	fmt.Printf("  %s%s%s  %s  %s%d/%d%s\n",
		Bold, player.Name, Reset,
		HPBar(player.CurrentHP, player.MaxHP, 25),
		Green, player.CurrentHP, player.MaxHP, Reset)

	// Vérification de mort
	if player.CurrentHP <= 0 {
		time.Sleep(300 * time.Millisecond)
		fmt.Println()
		SlowPrint(Red+"☠ Le joueur est à terre...", 30*time.Millisecond)
	}
}

// ==================== AFFICHAGE DE LA VIE ====================

// DisplayHP affiche la vie du monstre avec une barre visuelle
func (m *Monster) DisplayHP() {
	fmt.Printf("  %s%s%s  %s  %s%d/%d PV%s\n",
		Bold, m.Name, Reset,
		HPBar(m.CurrentHP, m.MaxHP, 25),
		Red, m.CurrentHP, m.MaxHP, Reset)
}

// Version simple (ancien style) pour compatibilité
func (m *Monster) DisplayHPSimple() {
	fmt.Printf("%s : %s%d / %d PV%s\n",
		m.Name, hpColorMonster(m), m.CurrentHP, m.MaxHP, Reset)
}

// ==================== COULEUR SELON HP ====================

func hpColorMonster(m *Monster) string {
	if m.MaxHP <= 0 {
		return Red
	}
	ratio := float64(m.CurrentHP) / float64(m.MaxHP)
	if ratio > 0.6 {
		return Green
	}
	if ratio > 0.3 {
		return Yellow
	}
	return Red
}

// ==================== ANIMATIONS SPÉCIALES MONSTRE ====================

// Animation d'apparition du monstre
func (m *Monster) AppearAnimation() {
	fmt.Println()
	SlowPrint(Yellow+"⚡ Un "+m.Name+" apparaît !", 30*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	// Petit tremblement
	for i := 0; i < 3; i++ {
		fmt.Print("  ")
		fmt.Print(Red + "▓▓▓" + Reset)
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\r     \r")
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Print(White + "Préparation au combat" + Reset)
	LoadingDots()
	time.Sleep(400 * time.Millisecond)
}

// Animation de mort du monstre
func (m *Monster) DeathAnimation() {
	SlowPrint(Red+"Le "+m.Name+" s'effondre...", 25*time.Millisecond)
	time.Sleep(200 * time.Millisecond)

	// Effet de disparition
	art := monsterArt(m.Name)
	lines := len(art)

	for i := lines; i > 0; i-- {
		Clear()
		fmt.Println()
		// Afficher seulement les premières lignes
		displayPartialArt(art, i)
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(300 * time.Millisecond)
	fmt.Println()
	SlowPrint(Green+"✨ Le monstre est vaincu !", 20*time.Millisecond)
}

// ==================== ART DU MONSTRE ====================

func monsterArt(name string) string {
	switch name {
	case "Goomba d'entraînement":
		return GoombaArt
	case "Koopa Tropique":
		return GoombaArt
	case "Bowser":
		return BowserArt
	}
	return ""
}

func displayPartialArt(art string, keepLines int) {
	lines := splitLines(art)
	max := keepLines
	if max > len(lines) {
		max = len(lines)
	}
	for i := 0; i < max; i++ {
		fmt.Println(lines[i])
	}
}

func splitLines(s string) []string {
	var lines []string
	current := ""
	for _, ch := range s {
		if ch == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

// ==================== ATTAQUE SPÉCIALE DU MONSTRE ====================

// AttackPlayer : variante avec animation d'attaque personnalisée
func (m *Monster) AttackPlayer(player *Character, damage int) {
	// Charge
	fmt.Printf("  %s%s%s charge", Bold, m.Name, Reset)
	LoadingDots()
	time.Sleep(300 * time.Millisecond)

	// Frappe
	AnimateAttack(m.Name, player.Name)
	FlashDamage()

	player.CurrentHP -= damage
	if player.CurrentHP < 0 {
		player.CurrentHP = 0
	}

	SlowPrint(Red+fmt.Sprintf("%s inflige %d dégâts !", m.Name, damage), 15*time.Millisecond)
}

// ==================== VÉRIFICATIONS ====================

// IsDeadMonster vérifie si le monstre est mort
func (m *Monster) IsDeadMonster() bool {
	return m.CurrentHP <= 0
}

// HealthRatio retourne le ratio de vie (0.0 à 1.0)
func (m *Monster) HealthRatio() float64 {
	if m.MaxHP <= 0 {
		return 0
	}
	return float64(m.CurrentHP) / float64(m.MaxHP)
}

// StatusEmoji retourne un emoji selon l'état du monstre
func (m *Monster) StatusEmoji() string {
	ratio := m.HealthRatio()
	switch {
	case ratio > 0.7:
		return "💪"
	case ratio > 0.4:
		return "😐"
	case ratio > 0.15:
		return "😰"
	case ratio > 0:
		return "🩸"
	default:
		return "☠"
	}
}
