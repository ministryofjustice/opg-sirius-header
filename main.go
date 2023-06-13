package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type Template interface {
	ExecuteTemplate(io.Writer, string, interface{}) error
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	// This is inefficient - it reads the template from the
	// filesystem every time. This makes it much easier to
	// develop though, so I can edit my template and the
	// changes will be reflected without having to restart
	// the app.
	t, err := template.ParseGlob("web/template/*.gotmpl")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err.Error()), 500)
		return
	}

	err = t.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err.Error()), 500)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "page.gotmpl", struct {
		prefix string
		sirius string
	}{
		prefix: "web/assets/static",
		sirius: "localhost:8080/",
	})
}

func logreq(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)

		f(w, r)
	})
}

type App struct {
	Port string
}

func (a *App) Start() {
	port := getEnv("PORT", "3456")
	prefix := getEnv("PREFIX", "")
	logger := log.New(os.Stdout, "opg-sirius-header ", log.LstdFlags)

	http.Handle(prefix+"/", logreq(index))

	logger.Println("Sirius header running at port " + port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	server := App{
		Port: getEnv("PORT", "3456"),
	}
	server.Start()
}

func getEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}
