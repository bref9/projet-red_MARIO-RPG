package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const saveFile = "save.json"

func SaveGame(c *Character) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println(Red + "Erreur de sauvegarde." + Reset)
		return
	}
	if err := os.WriteFile(saveFile, data, 0644); err != nil {
		fmt.Println(Red + "Erreur d'écriture." + Reset)
		return
	}
	fmt.Println(Green + "✓ Partie sauvegardée !" + Reset)
}

func LoadGame() (*Character, bool) {
	data, err := os.ReadFile(saveFile)
	if err != nil {
		return nil, false
	}
	var c Character
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, false
	}
	return &c, true
}

func SaveExists() bool {
	_, err := os.Stat(saveFile)
	return err == nil
}
