package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
	})
	http.ListenAndServe(":8080", r)

}
