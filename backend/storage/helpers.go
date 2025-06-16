package storage

import "improov/models"

func GetOrCreateUser(userID string) (*models.User, error) {
	var user models.User
	if err := DB.Preload("Completed").FirstOrCreate(&user, models.User{ID: userID}).Error; err != nil {
		return nil, err
	}
	user.Stats = models.UserStats{
		Fitness:     user.Stats.Fitness,
		Learning:    user.Stats.Learning,
		Mindfulness: user.Stats.Mindfulness,
	}

	return &user, nil
}
