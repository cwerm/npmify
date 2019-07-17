package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"npmify/util"
	"os"
)

type Dependencies struct {
	OutdatedCount   int   `json:"outdated_count"`
	TotalDependencies 	int 	`json:"total_dependencies"`
	Bower []struct {
		Name       string `json:"name"`
		Version    string `json:"version"`
		NpmVersion string `json:"npm_version"`
		Type       string `json:"type"`
		Outdated   bool   `json:"outdated"`
	} `json:"bower"`
}

func Init(dataFile string) {
	dependencies := Dependencies{}
	mux := http.NewServeMux()
	df, err := os.Open(dataFile)

	server := &http.Server{
		Addr:    ":1234",
		Handler: mux,
	}

	layout, err := template.ParseFiles("tpl/index.gohtml")
	util.CheckErr(err)

	err = json.NewDecoder(df).Decode(&dependencies)
	util.CheckErr(err)

	fileServer := http.FileServer(http.Dir("tpl/assets/"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
	mux.HandleFunc("/", home(layout, &dependencies))

	util.FancyPrint("Serving content on %s\n", "http://localhost:" + server.Addr)

	log.Fatal(server.ListenAndServe())
}

func home(layout *template.Template, deps *Dependencies) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		err := layout.Execute(w, deps)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panic(err)
		}

	}
}