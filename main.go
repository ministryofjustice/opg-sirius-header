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

func logreq(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)

		f(w, r)
	})
}

type App struct {
	Port       string
	StaticBase string
	SiriusUrl  string
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
		AdminHref        string
		ClientSearchHref string
		FinanceHref      string
		PoaHref          string
		SignOutHref      string
		SupervisionHref  string
		StaticBase       string
		WorkflowHref     string
		Url              string
	}{
		AdminHref:        a.SiriusUrl + "/admin",
		ClientSearchHref: a.SiriusUrl + "/supervision/#/clients/search-for-client",
		FinanceHref:      a.SiriusUrl + "/supervision/#/finance-hub/reporting",
		PoaHref:          a.SiriusUrl + "/lpa",
		SignOutHref:      a.SiriusUrl + "/auth/logout",
		SupervisionHref:  a.SiriusUrl + "/supervision",
		WorkflowHref:     a.SiriusUrl + "/supervision/workflow",
		StaticBase:       a.StaticBase,
		Url:              a.SiriusUrl,
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
		SiriusUrl:  getEnv("SIRIUS_URL", "http://localhost:8080"),
	}
	server.Start()
}
