package main

import "fmt"

type Recipe struct {
	Name      string
	Cost      int
	Materials map[string]int
}

var recipes = []Recipe{
	{"Casquette Mario", 5, map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1}},
	{"Salopette Mario", 5, map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1}},
	{"Bottes Kuribo", 5, map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1}},
}

func (c *Character) Blacksmith() {
	for {
		Clear()
		PrintTitle("⚒  FORGERON")
		fmt.Printf("  %sVotre bourse : %d pièces%s\n\n", Yellow, c.Gold, Reset)

		for i, r := range recipes {
			fmt.Printf("  %s%d.%s %s (%d pièces)\n", Cyan, i+1, Reset, r.Name, r.Cost)
			for mat, qty := range r.Materials {
				have := c.CountItem(mat)
				color := Green
				if have < qty {
					color = Red
				}
				fmt.Printf("       %s- %s x%d (vous : %d)%s\n",
					color, mat, qty, have, Reset)
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

	// Vérifier la place
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
	fmt.Printf(Green+"✓ Vous fabriquez : %s !\n"+Reset, r.Name)
}
