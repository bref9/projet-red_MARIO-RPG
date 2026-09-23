package main

import "fmt"

// ==================== STRUCTURE ====================

type MerchantItem struct {
	Name  string
	Price int
}

// ==================== CATALOGUE DU MARCHAND ====================

var merchantItems = []MerchantItem{
	{"Champignon Super", 3},
	{"Champignon Poison", 6},
	{"Potion de Mana", 4},
	{"Fleur de Feu", 25},
	{"Fourrure de Loup", 4},
	{"Peau de Troll", 7},
	{"Cuir de Sanglier", 3},
	{"Plume de Corbeau", 1},
	{"Augmentation d'inventaire", 30},
}

// ==================== MARCHAND ====================

func (c *Character) Merchant() {
	for {
		Clear()

		// 🎨 ASCII art du marchand
		fmt.Println(Cyan + MerchantArt + Reset)

		PrintTitle("🛒 MARCHAND")

		// 🎁 BONUS : première visite → champignon gratuit
		if !c.HasFreePotion {
			fmt.Println()
			fmt.Println(Green + "🎁 Cadeau de bienvenue ! Le marchand vous offre un Champignon Super !" + Reset)
			c.AddInventory("Champignon Super")
			c.HasFreePotion = true
			fmt.Println(Green + "  → Champignon Super ajouté à votre inventaire." + Reset)

			// 🎵 Son de pièce (bonus gratuit)
			PlayCoinSound()

			Pause()
			continue
		}

		fmt.Printf("  %sVotre bourse : %d pièces%s\n", Yellow, c.Gold, Reset)
		fmt.Printf("  %sInventaire  : %d / %d%s\n\n",
			Cyan, len(c.Inventory), c.MaxInventory, Reset)

		for i, item := range merchantItems {
			color := White

			// Afficher en rouge si trop cher
			if c.Gold < item.Price {
				color = Red
			}

			// Afficher en violet si limite atteinte pour l'agrandissement
			if item.Name == "Augmentation d'inventaire" && c.InventoryPurchases >= 3 {
				color = Purple
			}

			suffix := ""
			if item.Name == "Augmentation d'inventaire" {
				suffix = fmt.Sprintf(" %s[%d/3 achetés]%s",
					Yellow, c.InventoryPurchases, Reset)
			}

			fmt.Printf("  %s%d.%s %-30s %s%d pièces%s%s\n",
				Cyan, i+1, Reset, item.Name, color, item.Price, Reset, suffix)
		}
		fmt.Printf("\n  %s0.%s Retour\n", Yellow, Reset)

		choice := AskInt("\nChoix : ")
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(merchantItems) {
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
			continue
		}

		c.BuyItem(merchantItems[choice-1])
		Pause()
	}
}

// ==================== ACHAT D'ITEM ====================

func (c *Character) BuyItem(item MerchantItem) {
	// 🎁 BONUS : limite de 3 achats d'agrandissement d'inventaire
	if item.Name == "Augmentation d'inventaire" {
		if c.InventoryPurchases >= 3 {
			fmt.Println(Red + "❌ Vous avez déjà acheté 3 agrandissements ! Limite atteinte." + Reset)
			return
		}
		if c.Gold < item.Price {
			fmt.Println(Red + "❌ Pas assez de pièces !" + Reset)
			return
		}
		c.Gold -= item.Price
		c.InventoryPurchases++
		c.InventoryUpgrades++
		c.MaxInventory += 10

		// 🎵 Son de pièce
		PlayCoinSound()

		fmt.Printf(Green+"✓ Capacité d'inventaire augmentée à %d ! (%d/3)\n"+Reset,
			c.MaxInventory, c.InventoryPurchases)
		return
	}

	// Vérifier l'or
	if c.Gold < item.Price {
		fmt.Println(Red + "❌ Pas assez de pièces !" + Reset)
		return
	}

	// Vérifier la place dans l'inventaire
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println(Red + "❌ Inventaire plein !" + Reset)
		return
	}

	// Acheter
	c.Gold -= item.Price
	c.AddInventory(item.Name)

	// 🎵 Son de pièce
	PlayCoinSound()

	fmt.Printf(Green+"✓ Vous achetez : %s (-%d pièces)\n"+Reset, item.Name, item.Price)
}
