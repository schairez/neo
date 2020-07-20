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
	log.Println("we here!")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	if err := templates.ExecuteTemplate(w, "layout", data); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

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

var layoutDir string = "frontend/templates/layouts"
var staticDir string = "frontend/stylesheets"
var templates = template.Must(template.ParseGlob(layoutDir + "/*.gohtml"))

func main() {
	// port := ""
	port := os.Getenv("PORT")
	log.Printf("port from osgetenv %v", port)
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	fsCSS := http.FileServer(http.Dir(staticDir))
	mux.Handle("/stylesheets/", http.StripPrefix("/stylesheets", fsCSS))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/entry/create", createEntry)
	log.Printf("Starting server on :%v", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)

	}

}
