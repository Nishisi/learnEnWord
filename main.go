package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	_ "github.com/lib/pq"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	if r.Method == "POST" {
		r.ParseForm()
		InsertData(r.Form["word"][0], r.Form["example"][0])
	}
	http.Redirect(w, r, "/", 301)
}

type Data struct {
	Words    []string
	Examples []string
}

func show(w http.ResponseWriter, r *http.Request) {

	words, examples := SelectData()

	p := Data{
		Words:    words,
		Examples: examples,
	}

	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
	tmpl.Execute(w, p)
}

func main() {
	http.Handle("/", &templateHandler{filename: "register.gtpl"})
	http.HandleFunc("/register", register)
	http.HandleFunc("/show", show)

	log.Println("Webサーバを開始します")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
