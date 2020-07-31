package router

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"

	//"github.com/go-chi/cors"
	"github.com/schairez/neo/server/app"
	"github.com/schairez/neo/server/config"
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

//New returns a new router
func New(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.StripSlashes)
	r.Get("/", app.IndexHandler)
	r.Get("/about", app.AboutHandler)

	// fmt.Println("from the router yo")
	// dir, _ := os.Getwd()
	// fmt.Println(dir)
	fs := http.FileServer(neuteredFileSystem{http.Dir("client")})
	r.Handle("/client/*", http.StripPrefix("/client/", fs))

	githubProvider := github.New(
		cfg.Security.Oauth2.Github.ClientID,
		cfg.Security.Oauth2.Github.ClientSecret,
		cfg.Security.Oauth2.Github.RedirectURL)
	goth.UseProviders(githubProvider)
	r.Get("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Println("error ", err)
		}
		app.LoadPage(w, "welcome", user)

	})
	r.Get("/auth", gothic.BeginAuthHandler)

	return r

}
