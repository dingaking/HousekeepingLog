package handler

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

const STATIC_URL string = "/static/"
const ADMIN_URL string = "/admin/"
const STATIC_ROOT string = "templates/"

type Context struct {
	Title  string
	Static string
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(ADMIN_URL):]
	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	} else {
		f, err := http.Dir(STATIC_ROOT).Open("index.html")
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

func render(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = ADMIN_URL
	tmpl_list := []string{"templates/base.html",
		fmt.Sprintf("templates/%s.html", tmpl)}
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
