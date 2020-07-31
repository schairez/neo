package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//https://drstearns.github.io/tutorials/gomiddleware/
//https://gowebexamples.com/advanced-middleware/

//good implementation of chi
//https://getgophish.com/blog/post/2018-12-02-building-web-servers-in-go/#implementing-middleware
//https://getgophish.com/blog/post/database-migrations-in-go/

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

//IndexHandler handles the /about query
func IndexHandler(w http.ResponseWriter, r *http.Request) {
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

	LoadPage(w, "index", data)

}

//AboutHandler handles the /about query
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	LoadPage(w, "about", nil)

}

var templates *template.Template

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
	var tplDir string = dir + "/client/templates"
	// var tplDir string = "../../client/templates"

	templates = template.Must(template.ParseGlob(tplDir + "/*.gohtml"))

}

//LoadPage executes the template based on uri defined name
func LoadPage(wr http.ResponseWriter, name string, data interface{}) {
	if err := templates.ExecuteTemplate(wr, name, data); err != nil {
		log.Println(err.Error())
		http.Error(wr, err.Error(), http.StatusInternalServerError)

	}

}
