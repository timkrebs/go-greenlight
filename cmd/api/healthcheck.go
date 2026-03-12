package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a fixed-format JSON response from a string. Notice how we're using a raw
	// string literal (enclosed with backticks) so that we can include double quote
	// characters in the JSON without needing to escape them? We also use the %q verb to
	// wrap the interpolated values in double quotes.

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountert a problem and could not process your request",
			http.StatusInternalServerError)
	}
}
