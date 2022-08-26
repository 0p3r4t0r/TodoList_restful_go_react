package data

import "backend/db/models"

/**
 * This is mock data to use until I get the actual database queries running.
 */

var Tasks = []models.Task{
	{Title: "Wake up on time", Points: 10},
	{Title: "Study", Points: 25},
	{Title: "Exercise", Points: 25},
}
