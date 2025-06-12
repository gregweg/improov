package models

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"` // emoji or icon code
	Rank        int    `json:"rank"` // order of display
}

var DefaultCategories = []Category{
	{Name: "Fitness", Description: "Stay fit and healthy"},
	{Name: "Learning", Description: "Expand your knowledge"},
	{Name: "Mindfulness", Description: "Practice mindfulness and meditation"},
}
