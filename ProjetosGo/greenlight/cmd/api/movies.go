package main

import (
	"net/http"
	"greenlight.liz.net/internal"
	"encoding/json"
)

var movies = []internal.Movie {
	{
	ID: 1,
	Title: "Toy Story",
	Author: "Liz",
	},
}

func (app *application) createMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Método não Permitido", http.StatusMethodNotAllowed)
		return
	}

	var newMovie internal.Movie

	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, "JSON INVÁLIDO", http.StatusBadRequest)
		return
	}

	newMovie.ID = int64(len(movies)) + 1

	movies = append(movies, newMovie)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

