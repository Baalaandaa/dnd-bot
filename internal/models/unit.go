package models

type Unit struct {
	Defence        int
	SaveDifficulty int
	Health         int
	MaxHealth      int

	Name    string
	Attacks []Attack
}
