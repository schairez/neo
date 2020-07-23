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
	"path/filepath"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.gohtml")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

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

func loadPage(wr http.ResponseWriter, name string, data interface{}) {
	if err := templates.ExecuteTemplate(wr, name, data); err != nil {
		log.Println(err.Error())
		http.Error(wr, err.Error(), http.StatusInternalServerError)

	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Path)

	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	loadPage(w, "index", data)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	loadPage(w, "about", nil)

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

var tplDir string = "client/templates"

var templates = template.Must(template.ParseGlob(tplDir + "/*.gohtml"))

func main() {
	port := os.Getenv("PORT")
	log.Printf("port from osgetenv %v\n", port)
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./client")})
	mux.Handle("client", http.NotFoundHandler())
	mux.Handle("/client/", http.StripPrefix("/client", fileServer))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/entry/create", createEntry)
	log.Printf("Starting server on :%v", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)

	}

}
