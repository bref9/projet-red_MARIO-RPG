package main

import "fmt"

type MerchantItem struct {
	Name  string
	Price int
}

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

func (c *Character) Merchant() {
	for {
		Clear()
		PrintTitle("🛒 MARCHAND")
		fmt.Printf("  %sVotre bourse : %d pièces%s\n\n", Yellow, c.Gold, Reset)

		for i, item := range merchantItems {
			color := White
			if c.Gold < item.Price {
				color = Red
			}
			fmt.Printf("  %s%d.%s %-30s %s%d pièces%s\n",
				Cyan, i+1, Reset, item.Name, color, item.Price, Reset)
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

func (c *Character) BuyItem(item MerchantItem) {
	if c.Gold < item.Price {
		fmt.Println(Red + "❌ Pas assez de pièces !" + Reset)
		return
	}

	if item.Name == "Augmentation d'inventaire" {
		if c.InventoryUpgrades >= 3 {
			fmt.Println(Red + "Vous avez déjà utilisé toutes vos augmentations !" + Reset)
			return
		}
		c.Gold -= item.Price
		c.UpgradeInventorySlot()
		return
	}

	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println(Red + "❌ Inventaire plein !" + Reset)
		return
	}

	c.Gold -= item.Price
	c.AddInventory(item.Name)
	fmt.Printf(Green+"✓ Vous achetez : %s (-%d pièces)\n"+Reset, item.Name, item.Price)
}
