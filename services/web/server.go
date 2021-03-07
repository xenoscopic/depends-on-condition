package main

import (
	"fmt"
	"io"
	"net/http"
)

// getDescription grabs a description from the describer service.
func getDescription() (string, error) {
	// Perform an HTTP request (and defer closure of the body stream.)
	response, err := http.Get("http://describer/")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Extract the description bytes.
	description, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Success.
	return string(description), nil
}

func main() {
	// Grab an initial description just to test that we can access the describer
	// service. If this fails, then exit immediately. This isn't necessary, but
	// it's useful to simulate a hard dependency on another service that needs
	// to be running and healthy before this service can succeed.
	if _, err := getDescription(); err != nil {
		panic("unable to grab initial description")
	}

	// Log when startup has completed.
	fmt.Println("Web is good to go!")

	// Create an HTTP endpoint that returns a random message.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		description, err := getDescription()
		if err != nil {
			http.Error(w, "unable to grab description", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf(
			"<!doctype html><html><body><h1>You are %s</h1></body></html>",
			description,
		)))
	})

	// Listen and serve requests.
	http.ListenAndServe(":http", nil)
}
