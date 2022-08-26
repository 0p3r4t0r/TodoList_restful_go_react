package models

import "gorm.io/gorm"

type CreateTaskDTO struct {
	Title  string  `json:"title"`
	Points float64 `json:"points"`
}

type Task struct {
	gorm.Model
	Title  string  `json:"title"`
	Points float64 `json:"points"`
}

// type Action struct {
// 	ID     string  `json:"id"`
// 	TaskID string  `json:"taskId"`
// 	Points float64 `json:"points"`
// }
