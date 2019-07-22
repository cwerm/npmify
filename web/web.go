package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"npmify/msg"
	"npmify/state"
	"os"
)

func Init(dataFile string) {
	dependencies := state.Dependencies{}
	mux := http.NewServeMux()
	df, err := os.Open(dataFile)

	server := &http.Server{
		Addr:    ":1234",
		Handler: mux,
	}

	layout, err := template.ParseFiles("tpl/index.gohtml")
	msg.CheckErr(err)

	js, err := template.ParseFiles("tpl/assets/npmify.js")

	err = json.NewDecoder(df).Decode(&dependencies)
	msg.CheckErr(err)

	fileServer := http.FileServer(http.Dir("tpl/assets/"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
	mux.HandleFunc("/", home(layout, &dependencies))
	mux.HandleFunc("/assets/npmify.js", jsTpl(js, &dependencies))

	msg.FancyPrint("Serving content on %s\n", "http://localhost:" + server.Addr)

	log.Fatal(server.ListenAndServe())
}

func home(layout *template.Template, deps *state.Dependencies) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		err := layout.Execute(w, deps)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panic(err)
		}

	}
}

func jsTpl(js *template.Template, deps *state.Dependencies) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		err := js.Execute(w, deps)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panic(err)
		}

	}
}