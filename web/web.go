package web

import (
	"net/http"
	"text/template"
)

// Values is a structure of variables to be replaced in the template.
type values struct {
	Title       string
	UsedSpace   float32
	TotalSpace  float32
	ActiveNodes int
}

// PageValues is a stucture of data to be pasted into web page template.
var PageValues values

func init() {
	PageValues.Title = "Arken Stats"

}

// Page is the default web handler to render the "stats" html page.
func Page(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	t = template.New("index.html")

	t, _ = t.ParseFiles("templates/index.html")

	t.Execute(w, PageValues)

}