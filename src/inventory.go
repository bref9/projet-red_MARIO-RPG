package main

import (
	"fmt"
	"strings"
	"time"
)

// ==================== GESTION DE BASE ====================

func (c *Character) AddInventory(item string) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println(Red + "❌ Inventaire plein !" + Reset)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func (c *Character) RemoveInventory(item string) bool {
	for i, v := range c.Inventory {
		if v == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Character) CountItem(item string) int {
	count := 0
	for _, v := range c.Inventory {
		if v == item {
			count++
		}
	}
	return count
}

// ==================== AFFICHAGE ====================

func (c *Character) AccessInventory() {
	for {
		Clear()
		PrintTitle("INVENTAIRE")
		fmt.Printf("  Capacité : %s%d / %d%s\n\n", Yellow, len(c.Inventory), c.MaxInventory, Reset)

		if len(c.Inventory) == 0 {
			fmt.Println(Red + "  Inventaire vide." + Reset)
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("  %s%d.%s %s\n", Cyan, i+1, Reset, item)
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

// ==================== UTILITAIRE ====================

func hasPrefixAny(s string, prefixes []string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
