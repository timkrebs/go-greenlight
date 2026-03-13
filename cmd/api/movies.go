package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/timkrebs/greenlight/internal/data"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create new movie")
	w.Write([]byte("createMovieHandler"))
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// When httprouter is parsing a request, any interpolated URL parameters will be
	// stored in the request context. We can use the ParamsFromContext() function to
	// retrieve a slice containing these parameter names and values.
	//params := httprouter.ParamsFromContext(r.Context())
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Star Wars",
		Runtime:   180,
		Genres:    []string{"drama", "action", "sci-fi"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)

	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encounterd a problem and could not process your request", http.StatusInternalServerError)
	}

	// Interpolate the movie ID in a placeholder response
	fmt.Fprintf(w, "show the details of the movie: %d\n", id)
}
