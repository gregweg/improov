package storage

import "improov/models"

func SeedTasks() {
	tasks := []models.Task{
		// Fitness
		{ID: "fit-001", Category: "Fitness", Description: "Do 10 push-ups"},
		{ID: "fit-002", Category: "Fitness", Description: "Take a 5-minute walk"},
		{ID: "fit-003", Category: "Fitness", Description: "Stretch your hamstrings for 2 minutes"},
		{ID: "fit-004", Category: "Fitness", Description: "Do 20 jumping jacks"},

		// Learning
		{ID: "learn-001", Category: "Learning", Description: "Read 1 page from a book"},
		{ID: "learn-002", Category: "Learning", Description: "Watch a 2-minute educational video"},
		{ID: "learn-003", Category: "Learning", Description: "Google a topic you’re curious about"},
		{ID: "learn-004", Category: "Learning", Description: "Learn a new word and its meaning"},

		// Mindfulness
		{ID: "mind-001", Category: "Mindfulness", Description: "Take 3 deep breaths"},
		{ID: "mind-002", Category: "Mindfulness", Description: "Sit in silence for 1 minute"},
		{ID: "mind-003", Category: "Mindfulness", Description: "List 1 thing you're grateful for"},
		{ID: "mind-004", Category: "Mindfulness", Description: "Close your eyes and focus on your breath for 30 seconds"},
	}

	for _, task := range tasks {
		DB.FirstOrCreate(&task, models.Task{ID: task.ID})
	}
}

func SeedCategories() {
	categories := []models.Category{
		{ID: 1, Name: "Fitness", Icon: "🏃‍♂️", Rank: 1},
		{ID: 2, Name: "Learning", Icon: "📚", Rank: 2},
		{ID: 3, Name: "Mindfulness", Icon: "🧘‍♂️", Rank: 3},
	}

	for _, category := range categories {
		DB.FirstOrCreate(&category, models.Category{Name: category.Name})
	}
}
