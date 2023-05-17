package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const (
	username = "admin"
	password = "password"
)

func executeCommand(w http.ResponseWriter, r *http.Request) {
	// Check if the request includes basic authentication credentials
	user, pass, ok := r.BasicAuth()
	if !ok || user != username || pass != password {
		// Authentication failed
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Command to list files in the current directory
	cmd := exec.Command("htop")

	// Output will be stored in this variable
	output, err := cmd.Output()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Print the output of the command
	fmt.Fprintf(w, "%s", output)
}

func main() {
	http.HandleFunc("/ls", executeCommand)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
