package view

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

// var (
// 	tpIndex = parseTemplate("index.html")
// )

var m = minify.New()

func init() {
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
}

func joinTemplateDir(files ...string) []string {
	dirLevel1 := "views"
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(dirLevel1, f)
	}
	return r
}

func parseTemplate(file string) *template.Template {
	rootTemplate := []string{
		"template/root.html",
		"template/header.html",
		"template/menu.html",
	}
	filesOutput := joinTemplateDir(append(rootTemplate, file)...)
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(filesOutput...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	m.Minify("text/html", w, &buf)
}
