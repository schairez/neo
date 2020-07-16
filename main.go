package main

// Create a new idea,todo,doing,done entry
// POST /categories

// Update an existing idea,todo,doing,done entry
// PUT /categories/12

// View the details of a idea,todo,doing,done entry
// GET /categories/12

// Delete an existing idea,todo,doing,done entry
// DELETE /categories/12
import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Welcome To my awesome site</h1>")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))

}
func createEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", 405)

		// w.Write([]byte("Method Not Allowed"))
		return

	}
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"name":"Alex"}`))
	w.Write([]byte("Create a new snippet..."))

}

func main() {
	port := "4000"

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/entry/create", createEntry)
	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)

}
