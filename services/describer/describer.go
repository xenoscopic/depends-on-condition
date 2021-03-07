package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// descriptions are the descriptions that this service can return.
var descriptions = []string{
	"awesome 🤩",
	"goofy 🙃",
	"creative 🤔",
	"an awesome/goofy/creative developer 🐳",
}

func main() {
	// Sleep for 10 seconds to simulate a non-trivial service startup time.
	time.Sleep(10 * time.Second)

	// Log when startup has completed.
	fmt.Println("Describer is good to go!")

	// Create an HTTP endpoint that returns a random descriptions.
	http.ListenAndServe(":http", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(descriptions[rand.Intn(len(descriptions))]))
	}))
}
