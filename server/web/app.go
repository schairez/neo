package web

// Create a new idea,todo,doing,done entry
// POST /categories

// Update an existing idea,todo,doing,done entry
// PUT /categories/12

// View the details of a idea,todo,doing,done entry
// GET /categories/12

// Delete an existing idea,todo,doing,done entry
// DELETE /categories/12
import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

// oauthConfig := &oauth2.Config{
// 	  ClientID:     "your own client id here",
//     ClientSecret: "your own client secret here",
//     Endpoint:     oauth2GitHub.Endpoint,
//     RedirectURL:  "localhost:8000/auth/google/callback",
//     Scopes:       []string{"user:email"},
// }

//http://localhost:8080/oauth/redirect?code=cfca85c8b054585670dd

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

const githubClientID = "db5cce4a188f854c1cc0"
const githubClientSecret = "d2f10479ea5f09fdada8f9840a0676a98a74a181"

func main() {
	port := os.Getenv("PORT")
	log.Printf("port from osgetenv %v\n", port)
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(neuteredFileSystem{http.Dir("./client")})
	mux.Handle("client", http.NotFoundHandler())
	mux.Handle("/client/", http.StripPrefix("/client", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/entry/create", createEntry)
	log.Printf("Starting server on :%v", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)

	}

	// Create a new redirect route route
	mux.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		// First, we need to get the value of the `code` query param
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		code := r.FormValue("code")

		// Next, lets for the HTTP request to call the github oauth enpoint
		// to get our access token
		reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", githubClientID, githubClientSecret, code)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		// We set this header since we want the response
		// as JSON
		req.Header.Set("accept", "application/json")

		// Send out the HTTP request
		res, err := httpClient.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		defer res.Body.Close()

		// Parse the request body into the `OAuthAccessResponse` struct
		var t OAuthAccessResponse
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		// Finally, send a response to redirect the user to the "welcome" page
		// with the access token
		w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
		w.WriteHeader(http.StatusFound)
	})

}

type OAuth2Mock struct{}

// AuthCodeURL redirects to our own server.
// A handler which is only available in development handles the request.
func (o *OAuth2Mock) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost",
		Path:   "login/oauth/authorize",
	}

	v := url.Values{}
	v.Set("state", state)

	u.RawQuery = v.Encode()
	return u.String()
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
