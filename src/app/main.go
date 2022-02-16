/**
 *=Simple web server that uses the text/template (and html/template) package's "block" feature
 * to implement a kind of template inheritance.
 *
 * It should be executed from the directory in which the source resides,
 * as it will look for its template files in the current directory.
 *
 * -for example
 *  > cd .\src\app
 *	> go build
//	> .\app.exe
//
*  Chrome: http://localhost:8077
*
*/
package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// запускаем сервер (ip-address:port) и логируем все фатальные ошибки в лог -- куда именно пока не важно
	log.Fatal(http.ListenAndServe("localhost:8077", nil))
}

// indexTemplate is the main site template.
// The default template includes two template blocks ("sidebar" and "content")
// that may be replaced in templates derived from this one.
var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))

// Index is a data structure used to populate an indexTemplate.
type Index struct {
	Title       string
	Description string
	Body        string
	Links       []Link
}

type Link struct {
	URL, Title string
}

// indexHandler is an HTTP handler that serves the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &Index{
		Title:       "HelloWorld Go App v0.0",
		Description: "This is a simple description about this App",
		Body:        "Hello from Go World :)",
	}

	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}
