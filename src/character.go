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

	// Champs bonus
	HasFreePotion      bool
	InventoryPurchases int
}

// ==================== INITIALISATION ====================

func InitCharacter(name, class string, maxHP, manaMax int) *Character {
	return &Character{
		Name:               name,
		Class:              class,
		Level:              1,
		MaxHP:              maxHP,
		CurrentHP:          maxHP / 2,
		Mana:               manaMax,
		ManaMax:            manaMax,
		Gold:               100,
		Inventory:          []string{},
		MaxInventory:       10,
		InventoryUpgrades:  0,
		Skills:             []string{"Saut"},
		Equipment:          Equipment{},
		Initiative:         10,
		Exp:                0,
		ExpMax:             100,
		HasFreePotion:      false,
		InventoryPurchases: 0,
	}
}

// ==================== CRÉATION ====================

// CharacterCreation gère tout le flux de création de personnage
func CharacterCreation() *Character {
	Clear()
	fmt.Println(Yellow + MarioLogo + Reset)

	// Étape 1 : Choix du personnage
	choice := askCharacterChoice()

	var name, class string
	var maxHP, manaMax int

	// Étape 2a : Personnages prédéfinis (Mario, Luigi, Peach)
	if choice >= 1 && choice <= 3 {
		switch choice {
		case 1:
			name, class, maxHP, manaMax = "Mario", "Mario", 100, 50
		case 2:
			name, class, maxHP, manaMax = "Luigi", "Luigi", 80, 80
		case 3:
			name, class, maxHP, manaMax = "Princesse Peach", "Princesse Peach", 90, 70
		}
	} else {
		// Étape 2b : Personnage personnalisé → demander nom puis classe
		Clear()
		PrintTitle("📝 CRÉATION DE PERSONNAGE")
		fmt.Println()
		fmt.Println(Cyan + "Crée ton propre héros ! Choisis ton nom." + Reset)
		fmt.Println()
		name = askName()

		Clear()
		PrintTitle("🎭 CHOISIS TA CLASSE")
		fmt.Println()
		fmt.Println("  1. " + Blue + "Humain" + Reset + "   (100 PV, 50 Mana) - Équilibré")
		fmt.Println("  2. " + Blue + "Toad" + Reset + "     (80 PV, 80 Mana)  - Mage agile")
		fmt.Println("  3. " + Blue + "Koopa" + Reset + "    (90 PV, 70 Mana)  - Soutien")
		fmt.Println()

		for {
			classChoice := AskInt("Choix : ")
			switch classChoice {
			case 1:
				class, maxHP, manaMax = "Humain", 100, 50
			case 2:
				class, maxHP, manaMax = "Toad", 80, 80
			case 3:
				class, maxHP, manaMax = "Koopa", 90, 70
			default:
				fmt.Println(Red + "Choix invalide." + Reset)
				continue
			}
			break
		}
	}

	// Étape 3 : Créer le personnage
	player := InitCharacter(name, class, maxHP, manaMax)

	// Étape 4 : Afficher l'ASCII art
	Clear()
	DisplayCharacterArt(class)

	fmt.Println(Green + "\n★ Personnage créé avec succès ! ★" + Reset)
	fmt.Printf("  Nom    : %s\n", player.Name)
	fmt.Printf("  Classe : %s\n", player.Class)
	fmt.Printf("  PV     : %d / %d\n", player.CurrentHP, player.MaxHP)
	fmt.Printf("  Mana   : %d / %d\n", player.Mana, player.ManaMax)
	fmt.Printf("  Pièces : %d\n", player.Gold)
	Pause()

	return player
}

// askCharacterChoice affiche le menu principal de choix de personnage
func askCharacterChoice() int {
	Clear()
	PrintTitle("🎮 CHOISIS TON PERSONNAGE")
	fmt.Println()
	fmt.Println("  1. " + Red + "Mario" + Reset + "           (100 PV, 50 Mana) - Équilibré")
	fmt.Println("  2. " + Green + "Luigi" + Reset + "           (80 PV, 80 Mana)  - Mage agile")
	fmt.Println("  3. " + Purple + "Princesse Peach" + Reset + " (90 PV, 70 Mana)  - Soutien")
	fmt.Println()
	fmt.Println("  4. " + Cyan + "Créer un nouveau personnage" + Reset)
	fmt.Println()

	for {
		choice := AskInt("Choix : ")
		if choice >= 1 && choice <= 4 {
			return choice
		}
		fmt.Println(Red + "Choix invalide." + Reset)
	}
}

// askName demande un nom (uniquement pour la création personnalisée)
func askName() string {
	for {
		fmt.Print(Green + "→ " + Reset)
		var name string
		fmt.Scanln(&name)

		if isValidName(name) {
			return formatName(name)
		}
		fmt.Println(Red + "❌ Nom invalide. Utilise uniquement des lettres." + Reset)
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

// DisplayCharacterArt affiche l'ASCII art du personnage selon sa classe
func DisplayCharacterArt(class string) {
	switch class {
	case "Mario":
		fmt.Println(Red + MarioArt + Reset)
		fmt.Println(Bold + Red + "                    MARIO" + Reset)

	case "Humain":
		fmt.Println(Red + MarioArt + Reset)
		fmt.Println(Bold + Red + "                    HUMAIN" + Reset)

	case "Luigi":
		fmt.Println(Green + LuigiArt + Reset)
		fmt.Println(Bold + Green + "                    LUIGI" + Reset)

	case "Toad":
		fmt.Println(White + ToadArt + Reset)
		fmt.Println(Bold + White + "                    TOAD" + Reset)

	case "Princesse Peach":
		fmt.Println(Purple + PeachArt + Reset)
		fmt.Println(Bold + Purple + "                    PRINCESSE PEACH" + Reset)

	case "Koopa":
		fmt.Println(Green + KoopaArt + Reset)
		fmt.Println(Bold + Green + "                    KOOPA" + Reset)

	default:
		fmt.Println(White + MarioArt + Reset)
		fmt.Println(Bold + White + "                    " + strings.ToUpper(class) + Reset)
	}
}

// ==================== AFFICHAGE ====================

func (c *Character) DisplayInfo() {
	// 🎨 Afficher l'ASCII art du personnage en premier
	DisplayCharacterArt(c.Class)
	fmt.Println()

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
	fmt.Printf(Cyan+"║"+Reset+"  Inventaire : %-27s"+Cyan+"║\n"+Reset,
		fmt.Sprintf("%d / %d", len(c.Inventory), c.MaxInventory))
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

		// 🎁 Bonus : agrandissement automatique de l'inventaire
		c.MaxInventory += 2

		fmt.Println(Purple + "\n★ ★ ★ NIVEAU SUPÉRIEUR ! ★ ★ ★" + Reset)
		fmt.Printf(Yellow+"Vous êtes maintenant niveau %d !\n"+Reset, c.Level)
		fmt.Printf("  +10 PV max (%d)\n", c.MaxHP)
		fmt.Printf("  +10 Mana max (%d)\n", c.ManaMax)
		fmt.Printf("  +1 Initiative (%d)\n", c.Initiative)
		fmt.Printf(Green+"  +2 Slots d'inventaire (%d)\n"+Reset, c.MaxInventory)

		// 🎁 Bonus : sort de guérison au niveau 3
		if c.Level == 3 && !hasSpell(c.Skills, "Étoile") {
			c.Skills = append(c.Skills, "Étoile")
			fmt.Println()
			fmt.Println(Yellow + "✨ ✨ ✨ SORT DÉBLOQUÉ ! ✨ ✨ ✨" + Reset)
			fmt.Println(Green + "Vous apprenez : Étoile (soin)" + Reset)
			fmt.Println(Cyan + "  Coût : 20 mana" + Reset)
			fmt.Println(Cyan + "  Effet : Restaure 40 PV" + Reset)
		}
	}
}

func hasSpell(skills []string, spell string) bool {
	for _, s := range skills {
		if s == spell {
			return true
		}
	}
	return false
}
