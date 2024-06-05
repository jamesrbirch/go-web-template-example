package main

import (
	"bytes"
	"html/template"
	"net/http"
	"time"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type PhotoPageData struct {
	PageTitle string
	Photos    []Photo
}

type Photo struct {
	Id        int
	Filename  string
	Thumbnail string
	CreatedAt time.Time
}

type LayoutPageData struct {
	PageTitle string
	Content   template.HTML
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl := template.Must(template.ParseFiles("todo.html"))
	layout := template.Must(template.ParseFiles("layout.html"))

	var content bytes.Buffer

	_ = tmpl.Execute(&content, data)

	layoutContent := LayoutPageData{
		PageTitle: data.PageTitle,
		Content:   template.HTML(content.String()),
	}

	layout.Execute(w, layoutContent)
}

func photo(w http.ResponseWriter, r *http.Request) {

	data := PhotoPageData{
		PageTitle: "Photo Gallery",
		Photos: []Photo{
			{Filename: "image1.jpg", Thumbnail: "image1_thumb.jpg", CreatedAt: time.Now()},
		},
	}

	tmpl := template.Must(template.ParseFiles("photo.html"))
	tmpl.Execute(w, data)

}

func main() {

	http.HandleFunc("/", todo)
	http.HandleFunc("/photo", photo)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
