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
	"html/template"
	"log"
	"net/http"
	"os"
)

const layoutDir = "templates/layouts"

//Todo struct
type Todo struct {
	Title string
	Done  bool
}

//TodoPageData struct
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layouts/layout.gohtml"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)

	// w.Write([]byte("Hello from Snippetbox"))

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
	// port := "8000"
	port := os.Getenv("PORT")
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/entry/create", createEntry)
	log.Println("Starting server on :8000")

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)

}
