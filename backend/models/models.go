package models

import "gorm.io/gorm"

type User struct {
	ID        string          `gorm:"primaryKey" json:"id"` // can be username
	Name      string          `json:"name"`
	Password  string          `json:"-"` // store hashed password, omit in JSON
	Completed []CompletedTask `json:"completed"`
	Stats     UserStats       `json:"stats" gorm:"embedded"`
}

type UserStats struct {
	Fitness     int
	Learning    int
	Mindfulness int
}

type Task struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type CompletedTask struct {
	gorm.Model
	UserID string `json:"user_id"`
	TaskID string `json:"task_id"`
	Task   Task   `json:"task"`
}

type CompleteTaskRequest struct {
	UserID   string `json:"user_id"`
	TaskID   string `json:"task_id"`
	Category string `json:"category"` // add this tag!
}
