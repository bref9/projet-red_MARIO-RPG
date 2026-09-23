package main

import (
	"fmt"
	"strings"
	"time"
)

// ==================== GESTION DE BASE ====================

// AddInventory ajoute un item à l'inventaire (retourne false si plein)
func (c *Character) AddInventory(item string) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println(Red + "❌ Inventaire plein !" + Reset)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

// RemoveInventory retire un item de l'inventaire
func (c *Character) RemoveInventory(item string) bool {
	for i, v := range c.Inventory {
		if v == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

// CountItem compte combien de fois un item est présent dans l'inventaire
func (c *Character) CountItem(item string) int {
	count := 0
	for _, v := range c.Inventory {
		if v == item {
			count++
		}
	}
	return count
}

// HasItem vérifie si un item est dans l'inventaire
func (c *Character) HasItem(item string) bool {
	return c.CountItem(item) > 0
}

// CanAddItem vérifie s'il reste de la place dans l'inventaire
func (c *Character) CanAddItem() bool {
	return len(c.Inventory) < c.MaxInventory
}

// InventoryFullPercent retourne le pourcentage de remplissage
func (c *Character) InventoryFullPercent() int {
	if c.MaxInventory == 0 {
		return 100
	}
	return len(c.Inventory) * 100 / c.MaxInventory
}

// ListItemsByPrefix retourne tous les items commençant par un préfixe
func (c *Character) ListItemsByPrefix(prefix string) []string {
	var result []string
	for _, item := range c.Inventory {
		if strings.HasPrefix(item, prefix) {
			result = append(result, item)
		}
	}
	return result
}

// ==================== AFFICHAGE ====================

// AccessInventory affiche l'inventaire et permet d'utiliser un item (hors combat)
func (c *Character) AccessInventory() {
	for {
		Clear()
		PrintTitle("INVENTAIRE")
		fmt.Printf("  Capacité : %s%d / %d%s\n\n",
			Yellow, len(c.Inventory), c.MaxInventory, Reset)

		if len(c.Inventory) == 0 {
			fmt.Println(Red + "  Inventaire vide." + Reset)
		} else {
			for i, item := range c.Inventory {
				icon := itemIcon(item)
				fmt.Printf("  %s%d.%s %s %s\n", Cyan, i+1, Reset, icon, item)
			}
		}
		fmt.Println()
		fmt.Printf("  %s0.%s Retour\n", Yellow, Reset)

		choice := AskInt("\nChoix : ")
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
			continue
		}

		item := c.Inventory[choice-1]
		c.UseItem(item)
		Pause()
	}
}

// AccessInventoryInCombat affiche l'inventaire pendant un combat.
// Retourne true si un item a été utilisé (tour consommé),
// false si le joueur a juste consulté puis est revenu.
func (c *Character) AccessInventoryInCombat() bool {
	for {
		Clear()
		PrintTitle("INVENTAIRE (combat)")
		fmt.Printf("  Capacité : %s%d / %d%s\n\n",
			Yellow, len(c.Inventory), c.MaxInventory, Reset)

		if len(c.Inventory) == 0 {
			fmt.Println(Red + "  Inventaire vide." + Reset)
		} else {
			for i, item := range c.Inventory {
				icon := itemIcon(item)
				fmt.Printf("  %s%d.%s %s %s\n", Cyan, i+1, Reset, icon, item)
			}
		}
		fmt.Println()
		fmt.Printf("  %s0.%s Retour (sans utiliser d'item)\n", Yellow, Reset)

		choice := AskInt("\nChoix : ")
		if choice == 0 {
			return false // ✅ Retour sans consommer le tour
		}
		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
			continue
		}

		item := c.Inventory[choice-1]
		c.UseItem(item)
		Pause()
		return true // ✅ Item utilisé → tour consommé
	}
}

// itemIcon retourne un emoji selon l'item
func itemIcon(item string) string {
	switch {
	case strings.Contains(item, "Champignon Super"):
		return "🍄"
	case strings.Contains(item, "Champignon Poison"):
		return "☠️"
	case strings.Contains(item, "Potion de Mana"):
		return "🔵"
	case strings.Contains(item, "Fleur de Feu"):
		return "🔥"
	case strings.Contains(item, "Casquette"):
		return "🧢"
	case strings.Contains(item, "Salopette"):
		return "👕"
	case strings.Contains(item, "Bottes"):
		return "👢"
	case strings.Contains(item, "Fourrure"):
		return "🐺"
	case strings.Contains(item, "Peau de Troll"):
		return "🧟"
	case strings.Contains(item, "Cuir"):
		return "🐗"
	case strings.Contains(item, "Plume"):
		return "🪶"
	default:
		return "📦"
	}
}

// ==================== UTILISATION D'ITEMS ====================

// UseItem utilise un item de l'inventaire
func (c *Character) UseItem(item string) {
	switch item {
	case "Champignon Super":
		c.TakePot()
	case "Champignon Poison":
		c.PoisonPot()
	case "Fleur de Feu":
		c.SpellBook()
	case "Potion de Mana":
		c.DrinkManaPot()
	case "Casquette Mario", "Salopette Mario", "Bottes Kuribo":
		c.EquipItem(item)
	default:
		fmt.Printf(Yellow+"%s ne peut pas être utilisé ici.\n"+Reset, item)
	}
}

// ==================== POTIONS ====================

// TakePot utilise un Champignon Super (soin +50 PV)
func (c *Character) TakePot() {
	if !c.RemoveInventory("Champignon Super") {
		fmt.Println(Red + "Pas de Champignon Super dans l'inventaire." + Reset)
		return
	}
	fmt.Println(Green + "🍄 Vous utilisez Champignon Super !" + Reset)
	c.CurrentHP += 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf(Green+"PV : %d / %d\n"+Reset, c.CurrentHP, c.MaxHP)
}

// PoisonPot utilise un Champignon Poison (-10 PV/s pendant 3s)
func (c *Character) PoisonPot() {
	if !c.RemoveInventory("Champignon Poison") {
		fmt.Println(Red + "Pas de Champignon Poison dans l'inventaire." + Reset)
		return
	}
	fmt.Println(Purple + "☠ Vous utilisez Champignon Poison !" + Reset)
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}
		fmt.Printf(Red+"Poison ! PV : %d / %d\n"+Reset, c.CurrentHP, c.MaxHP)
		if c.IsDead() {
			return
		}
	}
}

// DrinkManaPot utilise une Potion de Mana (+30 Mana)
func (c *Character) DrinkManaPot() {
	if !c.RemoveInventory("Potion de Mana") {
		fmt.Println(Red + "Pas de Potion de Mana dans l'inventaire." + Reset)
		return
	}
	fmt.Println(Blue + "🔵 Vous buvez une Potion de Mana !" + Reset)
	c.Mana += 30
	if c.Mana > c.ManaMax {
		c.Mana = c.ManaMax
	}
	fmt.Printf(Blue+"Mana : %d / %d\n"+Reset, c.Mana, c.ManaMax)
}

// ==================== SORTS ====================

// SpellBook apprend le sort "Boule de Feu" grâce à la Fleur de Feu
func (c *Character) SpellBook() {
	for _, s := range c.Skills {
		if s == "Boule de Feu" {
			fmt.Println(Yellow + "Vous connaissez déjà ce sort !" + Reset)
			return
		}
	}
	c.Skills = append(c.Skills, "Boule de Feu")
	c.RemoveInventory("Fleur de Feu")
	fmt.Println(Green + "🔥 Vous apprenez : Boule de Feu !" + Reset)
}

// ==================== ÉQUIPEMENT ====================

// EquipItem équipe un item et applique le bonus de PV
func (c *Character) EquipItem(item string) {
	var slot *string
	var bonus int

	switch item {
	case "Casquette Mario":
		slot = &c.Equipment.Head
		bonus = 10
	case "Salopette Mario":
		slot = &c.Equipment.Torso
		bonus = 25
	case "Bottes Kuribo":
		slot = &c.Equipment.Feet
		bonus = 15
	default:
		return
	}

	// Si un équipement est déjà équipé, on le renvoie dans l'inventaire
	if *slot != "" {
		oldItem := *slot
		oldBonus := equipmentBonus(oldItem)
		c.MaxHP -= oldBonus
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
		c.AddInventory(oldItem)
		fmt.Printf(Yellow+"↩ Vous rangez : %s\n"+Reset, oldItem)
	}

	*slot = item
	c.MaxHP += bonus
	c.RemoveInventory(item)
	fmt.Printf(Green+"✓ Vous équipez : %s (+%d PV max)\n"+Reset, item, bonus)
}

// equipmentBonus retourne le bonus de PV d'un équipement
func equipmentBonus(item string) int {
	switch item {
	case "Casquette Mario":
		return 10
	case "Salopette Mario":
		return 25
	case "Bottes Kuribo":
		return 15
	}
	return 0
}

// ==================== UPGRADE INVENTAIRE ====================

// UpgradeInventorySlot augmente la capacité de l'inventaire de +10
// Limité à 3 utilisations
func (c *Character) UpgradeInventorySlot() bool {
	if c.InventoryUpgrades >= 3 {
		fmt.Println(Red + "Vous avez déjà utilisé toutes vos augmentations !" + Reset)
		return false
	}
	c.MaxInventory += 10
	c.InventoryUpgrades++
	fmt.Printf(Green+"✓ Capacité d'inventaire augmentée à %d !\n"+Reset, c.MaxInventory)
	return true
}

// ==================== AFFICHAGE DES ÉQUIPEMENTS ====================

// DisplayEquipment affiche l'équipement actuel du personnage
func (c *Character) DisplayEquipment() {
	fmt.Println(Cyan + "\n=== Équipement actuel ===" + Reset)
	fmt.Printf("  🧢 Tête  : %s\n", emptyOr(c.Equipment.Head, "—"))
	fmt.Printf("  👕 Torse : %s\n", emptyOr(c.Equipment.Torso, "—"))
	fmt.Printf("  👢 Pieds : %s\n", emptyOr(c.Equipment.Feet, "—"))
}

// ==================== UTILITAIRES ====================

// hasPrefixAny vérifie si une string commence par l'un des préfixes
func hasPrefixAny(s string, prefixes []string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
