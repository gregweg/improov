package storage

import "improov/models"

func GetOrCreateUser(userID string) (*models.User, error) {
	var user models.User
	if err := DB.Preload("Completed").FirstOrCreate(&user, models.User{ID: userID}).Error; err != nil {
		return nil, err
	}
	user.Stats = map[string]int{
		"fitness":     user.Fitness,
		"learning":    user.Learning,
		"mindfulness": user.Mindfulness,
	}

	return &user, nil
}
