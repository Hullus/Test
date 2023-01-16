package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"structs/internal/models"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) allMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "1993-03-31")

	highlander := models.Movie{
		ID:           1,
		Title:        "highlander",
		ReleaseDate:  rd,
		Runtime:      113,
		MPAARRating:  "sda",
		Description:  "Good",
		Image:        "asdsaas",
		CreatedAt:    time.Now(),
		UpdatedField: time.Now(),
	}

	movies = append(movies, highlander)
	movies = append(movies, highlander)
	movies = append(movies, highlander)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
