package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a fixed-format JSON response from a string. Notice how we're using a raw
	// string literal (enclosed with backticks) so that we can include double quote
	// characters in the JSON without needing to escape them? We also use the %q verb to
	// wrap the interpolated values in double quotes.

	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
