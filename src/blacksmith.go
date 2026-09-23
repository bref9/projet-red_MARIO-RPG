package main

import "fmt"

// ==================== STRUCTURES ====================

type Recipe struct {
	Name      string
	Cost      int
	Materials map[string]int
}

// ==================== RECETTES DU FORGERON ====================

var recipes = []Recipe{
	{"Casquette Mario", 5, map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1}},
	{"Salopette Mario", 5, map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1}},
	{"Bottes Kuribo", 5, map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1}},
}

// ==================== FORGERON ====================

func (c *Character) Blacksmith() {
	for {
		Clear()

		// 🎨 ASCII art du forgeron
		fmt.Println(Yellow + BlacksmithArt + Reset)

		PrintTitle("⚒  FORGERON")

		fmt.Printf("  %sVotre bourse : %d pièces%s\n", Yellow, c.Gold, Reset)
		fmt.Printf("  %sInventaire  : %d / %d%s\n\n",
			Cyan, len(c.Inventory), c.MaxInventory, Reset)

		// Afficher les recettes avec vérification des ressources
		for i, r := range recipes {
			// Vérifier si le joueur peut fabriquer
			canCraft := c.canCraft(r)

			// Couleur du titre
			titleColor := Green
			if !canCraft {
				titleColor = Red
			}

			fmt.Printf("  %s%d.%s %s%s%s (%d pièces)\n",
				Cyan, i+1, Reset, titleColor, r.Name, Reset, r.Cost)

			// Afficher les matériaux nécessaires
			for mat, qty := range r.Materials {
				have := c.CountItem(mat)
				matColor := Green
				if have < qty {
					matColor = Red
				}
				fmt.Printf("       %s- %s x%d (vous : %d)%s\n",
					matColor, mat, qty, have, Reset)
			}
			fmt.Println()
		}
		fmt.Printf("  %s0.%s Retour\n", Yellow, Reset)

		choice := AskInt("\nChoix : ")
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(recipes) {
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
			continue
		}

		c.CraftItem(recipes[choice-1])
		Pause()
	}
}

// ==================== FABRICATION D'ITEM ====================

func (c *Character) CraftItem(r Recipe) {
	// Vérifier l'or
	if c.Gold < r.Cost {
		fmt.Println(Red + "❌ Pas assez de pièces !" + Reset)
		return
	}

	// Vérifier les ressources
	for mat, qty := range r.Materials {
		if c.CountItem(mat) < qty {
			fmt.Printf(Red+"❌ Ressource manquante : %s x%d\n"+Reset, mat, qty)
			return
		}
	}

	// Vérifier la place dans l'inventaire
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println(Red + "❌ Inventaire plein !" + Reset)
		return
	}

	// Fabriquer
	c.Gold -= r.Cost
	for mat, qty := range r.Materials {
		for i := 0; i < qty; i++ {
			c.RemoveInventory(mat)
		}
	}
	c.AddInventory(r.Name)

	// 🎵 Son de marteau
	PlayHammerSound()

	fmt.Println()
	fmt.Println(Green + "╔════════════════════════════════════════╗" + Reset)
	fmt.Println(Green + "║       ⚒  FABRICATION RÉUSSIE ! ⚒        ║" + Reset)
	fmt.Println(Green + "╚════════════════════════════════════════╝" + Reset)
	fmt.Printf(Green+"✓ Vous fabriquez : %s\n"+Reset, r.Name)
	fmt.Printf(Yellow+"  -%d pièces (reste : %d)\n"+Reset, r.Cost, c.Gold)
}

// ==================== VÉRIFICATION ====================

// canCraft vérifie si le joueur peut fabriquer une recette
func (c *Character) canCraft(r Recipe) bool {
	// Vérifier l'or
	if c.Gold < r.Cost {
		return false
	}
	// Vérifier les ressources
	for mat, qty := range r.Materials {
		if c.CountItem(mat) < qty {
			return false
		}
	}
	// Vérifier la place
	if len(c.Inventory) >= c.MaxInventory {
		return false
	}
	return true
}
