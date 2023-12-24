package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"text/template"

	"bipbop/config"
	"bipbop/models"

	"github.com/gorilla/mux"
)

func Serve() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./www/public/"))
	r.PathPrefix("/www").Handler(http.StripPrefix("/www", fs))

	r.HandleFunc("/", Hello)
	r.HandleFunc("/data", Read).Methods("GET")
	r.HandleFunc("/data", Create).Methods("POST")
	r.HandleFunc("/data/{id}", Update).Methods("PUT")
	r.HandleFunc("/data/{id}", Delete).Methods("DELETE")

	http.ListenAndServe(fmt.Sprintf(":%s", config.ApiConfig.Port), r)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World !!")
}

func Read(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tmpl := template.Must( template.ParseFiles("www/template/menu.html") )
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var dreq models.LivroRequest

	err := json.NewDecoder(r.Body).Decode(&dreq)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	err = models.Insert(dreq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "ok")
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi( vars["id"] )
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var dreq models.LivroRequest
	err = json.NewDecoder(r.Body).Decode(&dreq)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	err = models.Update( int64(id), dreq )
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "ok")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi( vars["id"] )
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = models.Delete( int64(id) )
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "ok")
}