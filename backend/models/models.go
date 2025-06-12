package models

import "gorm.io/gorm"

type User struct {
	ID          string          `gorm:"primaryKey" json:"id"`
	Stats       map[string]int  `gorm:"-" json:"stats,omitempty"`
	Fitness     int             `json:"fitness"`
	Learning    int             `json:"learning"`
	Mindfulness int             `json:"mindfulness"`
	Completed   []CompletedTask `json:"completed"`
	Username    string          `json:"username"`
	Password    string          `json:"-"` // store hashed password, omit in JSON
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
