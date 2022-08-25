package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/db/models"

	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetTask(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPostTask(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(models.Task{
		Title:  "Test",
		Points: 10,
	})
	req, _ := http.NewRequest("POST", "/api/v1/tasks", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
