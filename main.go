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
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"gopkg.in/yaml.v2"
	// "github.com/markbates/goth/providers/facebook"
	// "github.com/markbates/goth/providers/discord"
)

// Config struct
type Config struct {
	Security struct {
		Oauth2 struct {
			Github struct {
				ClientID     string `yaml:"clientId"`
				ClientSecret string `yaml:"clientSecret"`
			}
		}
	}
	Server struct {
		Host string `yaml:"host"`
		Port int16  `yaml:"port"`
	}
}

func main() {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := &Config{}

	if err := yaml.UnmarshalStrict(yamlFile, c); err != nil {
		log.Println("oOoopps")
		log.Fatalf("Unmarshal: %v", err)
		os.Exit(1)

	}
	fmt.Printf("Config Result: %v\n", c)
	githubProvider := github.New(c.Security.Oauth2.Github.ClientID, c.Security.Oauth2.Github.ClientSecret, "http://localhost:8000/auth/callback?provider=github")
	goth.UseProviders(githubProvider)

	port := fmt.Sprintf("%d", c.Server.Port)
	log.Printf("port from osgetenv %v\n", port)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(neuteredFileSystem{http.Dir("./client")})
	// mux.Handle("client", http.NotFoundHandler())
	r.Handle("/client/*", http.StripPrefix("/client/", fs))

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/entry/create", createEntry)

	r.Get("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Println("error ", err)
		}
		loadPage(w, "welcome", user)

	})

	r.Get("/auth", gothic.BeginAuthHandler)

	http.ListenAndServe(":"+port, r)

}

//StrictSlash(false)

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
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

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
