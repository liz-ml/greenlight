package main

import (
	"net/http"
	"greenlight.liz.net/internal"
	"encoding/json"
	"time"
	
)

var movies = []internal.Movie {
	{
	ID: 1,
	Title: "Toy Story",
	Author: "Liz",
	Genres: []string{"action", "drama"}, 
	Version: 1,
	CreatedAt: time.Now(),
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

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	
	var movieFilter internal.Movie

	for _, movie := range movies {
		if movie.ID == id {
			movieFilter = movie
			break
		}
	}

	err = app.writeJSON(w, http.StatusOK, movieFilter, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}


