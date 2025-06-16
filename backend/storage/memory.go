package storage

import "improov/models"

var Tasks = []models.Task{
	{ID: "1", Category: "Fitness", Description: "Do 10 push-ups"},
	{ID: "2", Category: "Fitness", Description: "Take a 5-minute walk"},
	{ID: "3", Category: "Learning", Description: "Read 1 page of a book"},
	{ID: "4", Category: "Mindfulness", Description: "Take 3 deep breaths"},
}

var Users = map[string]*models.User{
	"demo": {
		ID: "demo",
		Stats: models.UserStats{
			Fitness:     0,
			Learning:    0,
			Mindfulness: 0,
		},
	},
}
