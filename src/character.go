package main

import (
	"fmt"
	"strings"
	"unicode"
)

// ==================== STRUCTURES ====================

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name              string
	Class             string
	Level             int
	MaxHP             int
	CurrentHP         int
	Mana              int
	ManaMax           int
	Gold              int
	Inventory         []string
	MaxInventory      int
	InventoryUpgrades int
	Skills            []string
	Equipment         Equipment
	Initiative        int
	Exp               int
	ExpMax            int
}

// ==================== INITIALISATION ====================

func InitCharacter(name, class string, maxHP, manaMax int) *Character {
	return &Character{
		Name:              name,
		Class:             class,
		Level:             1,
		MaxHP:             maxHP,
		CurrentHP:         maxHP / 2,
		Mana:              manaMax,
		ManaMax:           manaMax,
		Gold:              100,
		Inventory:         []string{},
		MaxInventory:      10,
		InventoryUpgrades: 0,
		Skills:            []string{"Saut"},
		Equipment:         Equipment{},
		Initiative:        10,
		Exp:               0,
		ExpMax:            100,
	}
}

// ==================== CRÉATION ====================

func CharacterCreation() *Character {
	Clear()
	fmt.Println(Yellow + MarioLogo + Reset)

	name := askName()
	class, maxHP, manaMax := askClass()

	player := InitCharacter(name, class, maxHP, manaMax)

	fmt.Println(Green + "\n★ Personnage créé avec succès ! ★" + Reset)
	fmt.Printf("  Nom    : %s\n", player.Name)
	fmt.Printf("  Classe : %s\n", player.Class)
	fmt.Printf("  PV     : %d / %d\n", player.CurrentHP, player.MaxHP)
	fmt.Printf("  Mana   : %d / %d\n", player.Mana, player.ManaMax)
	fmt.Printf("  Pièces : %d\n", player.Gold)
	Pause()

	return player
}

func askName() string {
	for {
		fmt.Print(Cyan + "Entrez votre nom (lettres uniquement) : " + Reset)
		var name string
		fmt.Scanln(&name)

		if isValidName(name) {
			return formatName(name)
		}
		fmt.Println(Red + "Nom invalide. Réessayez." + Reset)
	}
}

func isValidName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formatName(name string) string {
	name = strings.ToLower(name)
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func askClass() (string, int, int) {
	fmt.Println(Cyan + "\n=== Choisissez votre héros ===" + Reset)
	fmt.Println("  1. " + Red + "Mario" + Reset + "   (100 PV, 50 Mana) - Équilibré")
	fmt.Println("  2. " + Green + "Luigi" + Reset + "   (80 PV, 80 Mana)  - Mage agile")
	fmt.Println("  3. " + Yellow + "Bowser" + Reset + "  (120 PV, 30 Mana) - Tank puissant")

	for {
		choice := AskInt("Choix : ")
		switch choice {
		case 1:
			return "Mario", 100, 50
		case 2:
			return "Luigi", 80, 80
		case 3:
			return "Bowser", 120, 30
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}

// ==================== AFFICHAGE ====================

func (c *Character) DisplayInfo() {
	fmt.Println(Cyan + "╔══════════════════════════════════════════╗" + Reset)
	fmt.Printf(Cyan+"║"+Reset+"  %s%-38s%s"+Cyan+"║\n"+Reset,
		Bold+Yellow, c.Name, Reset)
	fmt.Println(Cyan + "╠══════════════════════════════════════════╣" + Reset)
	fmt.Printf(Cyan+"║"+Reset+"  Classe     : %-27s"+Cyan+"║\n"+Reset, c.Class)
	fmt.Printf(Cyan+"║"+Reset+"  Niveau     : %-27d"+Cyan+"║\n"+Reset, c.Level)
	fmt.Printf(Cyan+"║"+Reset+"  PV         : %s%-27s%s"+Cyan+"║\n"+Reset,
		hpColor(c), fmt.Sprintf("%d / %d", c.CurrentHP, c.MaxHP), Reset)
	fmt.Printf(Cyan+"║"+Reset+"  Mana       : %s%-27s%s"+Cyan+"║\n"+Reset,
		Blue, fmt.Sprintf("%d / %d", c.Mana, c.ManaMax), Reset)
	fmt.Printf(Cyan+"║"+Reset+"  Pièces     : %s%-27s%s"+Cyan+"║\n"+Reset,
		Yellow, fmt.Sprintf("%d", c.Gold), Reset)
	fmt.Printf(Cyan+"║"+Reset+"  Expérience : %-27s"+Cyan+"║\n"+Reset,
		fmt.Sprintf("%d / %d", c.Exp, c.ExpMax))
	fmt.Printf(Cyan+"║"+Reset+"  Initiative : %-27d"+Cyan+"║\n"+Reset, c.Initiative)
	fmt.Println(Cyan + "╠══════════════════════════════════════════╣" + Reset)
	fmt.Printf(Cyan+"║"+Reset+"  Équipement :%-28s"+Cyan+"║\n"+Reset, "")
	fmt.Printf(Cyan+"║"+Reset+"    Tête  : %-31s"+Cyan+"║\n"+Reset, emptyOr(c.Equipment.Head, "—"))
	fmt.Printf(Cyan+"║"+Reset+"    Torse : %-31s"+Cyan+"║\n"+Reset, emptyOr(c.Equipment.Torso, "—"))
	fmt.Printf(Cyan+"║"+Reset+"    Pieds : %-31s"+Cyan+"║\n"+Reset, emptyOr(c.Equipment.Feet, "—"))
	fmt.Printf(Cyan+"║"+Reset+"  Sorts      : %-27s"+Cyan+"║\n"+Reset, strings.Join(c.Skills, ", "))
	fmt.Println(Cyan + "╚══════════════════════════════════════════╝" + Reset)
}

func hpColor(c *Character) string {
	if c.MaxHP == 0 {
		return Red
	}
	ratio := float64(c.CurrentHP) / float64(c.MaxHP)
	if ratio > 0.6 {
		return Green
	}
	if ratio > 0.3 {
		return Yellow
	}
	return Red
}

func emptyOr(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}

// ==================== MORT / RÉSURRECTION ====================

func (c *Character) IsDead() bool {
	if c.CurrentHP <= 0 {
		c.CurrentHP = c.MaxHP / 2
		fmt.Println(Red + "\n☠  Vous êtes mort ! ☠" + Reset)
		fmt.Printf(Green+"Vous êtes ressuscité avec %d/%d PV.\n"+Reset, c.CurrentHP, c.MaxHP)
		return true
	}
	return false
}

// ==================== PROGRESSION ====================

func (c *Character) GainExp(amount int) {
	c.Exp += amount
	for c.Exp >= c.ExpMax {
		c.Exp -= c.ExpMax
		c.Level++
		c.ExpMax = c.ExpMax * 2
		c.MaxHP += 10
		c.CurrentHP += 10
		c.ManaMax += 10
		c.Mana = c.ManaMax
		c.Initiative += 1

		fmt.Println(Purple + "\n★ ★ ★ NIVEAU SUPÉRIEUR ! ★ ★ ★" + Reset)
		fmt.Printf(Yellow+"Vous êtes maintenant niveau %d !\n"+Reset, c.Level)
		fmt.Printf("  +10 PV max (%d)\n", c.MaxHP)
		fmt.Printf("  +10 Mana max (%d)\n", c.ManaMax)
		fmt.Printf("  +1 Initiative (%d)\n", c.Initiative)
	}
}
