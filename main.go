package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	// This is inefficient - it reads the template from the
	// filesystem every time. This makes it much easier to
	// develop though, so I can edit my template and the
	// changes will be reflected without having to restart
	// the app.
	t, err := template.ParseGlob("static/template/*.gotmpl")
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

func logreq(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)

		f(w, r)
	})
}

type App struct {
	Port       string
	StaticBase string
}

func (a *App) Start() {
	port := getEnv("PORT", "3456")
	prefix := getEnv("PREFIX", "")
	logger := log.New(os.Stdout, "opg-sirius-header ", log.LstdFlags)

	if a.StaticBase == "/static" {
		log.Printf("serving static assets")
		http.Handle("/static/", logreq(staticHandler("static")))
	}
	http.Handle(prefix+"/", logreq(a.index))

	logger.Println("Sirius header running at port " + port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}

func (a App) index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "page.gotmpl", struct {
		AdminHref       string
		LpaHref         string
		LogOutHref      string
		SupervisionHref string
		WorkflowHref    string
		TestString      string
		StaticBase      string
	}{
		AdminHref:       "http://localhost:8080/admin",
		LpaHref:         "http://localhost:8080/lpa",
		LogOutHref:      "http://localhost:8080/auth/logout",
		SupervisionHref: "http://localhost:8080/supervision/supervision",
		WorkflowHref:    "http://localhost:8080/supervision/workflow",
		TestString:      "Kate",
		StaticBase:      a.StaticBase,
	})
}

func staticHandler(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir(dir))).ServeHTTP(w, r)
	}
}

func getEnv(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func main() {
	server := App{
		Port:       getEnv("PORT", "3456"),
		StaticBase: getEnv("STATIC_BASE", "/static"),
	}
	server.Start()
}
